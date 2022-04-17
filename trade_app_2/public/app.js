window.addEventListener("DOMContentLoaded", (_) => {
  let websocket = new WebSocket("ws://" + window.location.host + "/websocket");
  let room = document.getElementById("chat-text");

  websocket.addEventListener("message", function (e) {
    let data = JSON.parse(e.data);
    // creating html element
    let p = document.createElement("p");
    p.innerHTML = `<strong>${data.msg_sender}</strong>: ${data.msg_text}      <p>${data.timestamp1}</p>`;

    //append para at top of chat-text
    room.prepend(p);
  });

  let form = document.getElementById("input-form");

  //event listner for submit button
  form.addEventListener("submit", function (event) {
    event.preventDefault();
    let date1 = new Date().toString();
    console.log(date1);
    let username = document.getElementById("input-username");
    let text = document.getElementById("input-text");
    websocket.send(
      JSON.stringify({
        msg_sender: username.value,
        msg_text: text.value,
        timestamp1: date1,
      })
    );
    text.value = "";
  });
});
