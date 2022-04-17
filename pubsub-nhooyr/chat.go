package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"golang.org/x/time/rate"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

// chatServer enables broadcasting to a set of subscribers.
type chatServer struct {
	// subscriberMessageBuffer controls the max number
	// of messages that can be queued for a subscriber
	// before it is kicked.
	//
	// Defaults to 16.
	subscriberMessageBuffer int

	// publishLimiter controls the rate limit applied to the publish endpoint.
	//
	// Defaults to one publish every 100ms with a burst of 8.
	publishLimiter *rate.Limiter

	// logf controls where logs are sent.
	// Defaults to log.Printf.
	logf func(f string, v ...interface{})

	// serveMux routes the various endpoints to the appropriate handler.
	serveMux http.ServeMux

	subscribersMu sync.Mutex
	subscribers   map[*subscriber]struct{}
}

// newChatServer constructs a chatServer with the defaults.
func newChatServer() *chatServer {
	cs := &chatServer{
		subscriberMessageBuffer: 16,
		logf:                    log.Printf,
		subscribers:             make(map[*subscriber]struct{}),
		publishLimiter:          rate.NewLimiter(rate.Every(time.Millisecond*100), 8),
	}
	cs.serveMux.Handle("/", http.FileServer(http.Dir(".")))
	cs.serveMux.HandleFunc("/events", cs.subscribeHandler)

	return cs
}

type WSMessage struct {
	Action  string `json:"action"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

// subscriber represents a subscriber.
// Messages are sent on the msgs channel and if the client
// cannot keep up with the messages, closeSlow is called.
type subscriber struct {
	msgs               chan *WSMessage
	closeSlow          func()
	subscribedChannels map[string]struct{}
}

func (cs *chatServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cs.serveMux.ServeHTTP(w, r)
}

// subscribeHandler accepts the WebSocket connection and then subscribes
// it to all future messages.
func (cs *chatServer) subscribeHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("---->Subscribers", cs.subscribers)
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		cs.logf("%v", err)
		return
	}
	defer c.Close(websocket.StatusInternalError, "")

	err = cs.subscribe(r.Context(), c)
	if errors.Is(err, context.Canceled) {
		return
	}
	if websocket.CloseStatus(err) == websocket.StatusNormalClosure ||
		websocket.CloseStatus(err) == websocket.StatusGoingAway {
		return
	}
	if err != nil {
		cs.logf("%v", err)
		return
	}
}

var t1 int64
var t2 int64

// subscribe subscribes the given WebSocket to all broadcast messages.
// It creates a subscriber with a buffered msgs chan to give some room to slower
// connections and then registers the subscriber. It then listens for all messages
// and writes them to the WebSocket. If the context is cancelled or
// an error occurs, it returns and deletes the subscription.
//
// It uses CloseRead to keep reading from the connection to process control
// messages and cancel the context if the connection drops.
func (cs *chatServer) subscribe(ctx context.Context, c *websocket.Conn) error {
	ctx, cancel := context.WithCancel(ctx)

	s := &subscriber{
		msgs: make(chan *WSMessage, cs.subscriberMessageBuffer),
		closeSlow: func() {
			c.Close(websocket.StatusPolicyViolation, "connection too slow to keep up with messages")
		},
		subscribedChannels: map[string]struct{}{"general": {}},
	}
	cs.addSubscriber(s)
	defer cs.deleteSubscriber(s)

	go func() {
		defer cancel()
		for {
			var wsMessage WSMessage
			err := wsjson.Read(ctx, c, &wsMessage)
			if err != nil {
				break
			}
			fmt.Println("")
			switch wsMessage.Action {
			case "TEXT":
				wsMessage.Text = strings.TrimSpace(strings.Replace(wsMessage.Text, "\n", " ", -1))
				cs.publish(&wsMessage)
				fmt.Println("Subscribers", cs.subscribers)
				fmt.Println("Channels", s.subscribedChannels)

			case "SUBSCRIBE":
				t1 = time.Now().UnixNano()
				s.subscribedChannels[wsMessage.Channel] = struct{}{}
				wsMessage.Text = "subscribe OK!"
				wsjson.Write(ctx, c, wsMessage)
				t2 = time.Now().UnixNano()
				fmt.Println("Time diff:", (t2 - t1))
			case "UNSUBSCRIBE":
				if _, ok := s.subscribedChannels[wsMessage.Channel]; ok {
					delete(s.subscribedChannels, wsMessage.Channel)
					wsMessage.Text = "unsubscribe OK!"
				} else {
					wsMessage.Text = fmt.Sprintf("Hey, you were not subscribed to channel `%s`...", wsMessage.Channel)
				}
				wsjson.Write(ctx, c, wsMessage)
			}
		}
	}()

	for {
		select {
		case msg := <-s.msgs:
			if _, ok := s.subscribedChannels[msg.Channel]; ok {
				err := writeTimeout(ctx, time.Second*5, c, msg)
				if err != nil {
					return err
				}
			}
		case <-ctx.Done():
			fmt.Println("Deleting done")
			return ctx.Err()
		}
	}
}

// publish publishes the msg to all subscribers.
// It never blocks and so messages to slow subscribers
// are dropped.
func (cs *chatServer) publish(msg *WSMessage) {
	cs.subscribersMu.Lock()
	defer cs.subscribersMu.Unlock()

	cs.publishLimiter.Wait(context.Background())

	for s := range cs.subscribers {
		select {
		case s.msgs <- msg:
		default:
			go s.closeSlow()
		}
	}
}

// addSubscriber registers a subscriber.
func (cs *chatServer) addSubscriber(s *subscriber) {
	cs.subscribersMu.Lock()
	cs.subscribers[s] = struct{}{}
	cs.subscribersMu.Unlock()
}

// deleteSubscriber deletes the given subscriber.
func (cs *chatServer) deleteSubscriber(s *subscriber) {
	cs.subscribersMu.Lock()
	fmt.Println("Deleted in delete sub")
	delete(cs.subscribers, s)
	cs.subscribersMu.Unlock()
}

func writeTimeout(ctx context.Context, timeout time.Duration, c *websocket.Conn, msg *WSMessage) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return wsjson.Write(ctx, c, msg)
}
