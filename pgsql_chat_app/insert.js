function test1(params) {
  const { Client } = require("pg");
  const client = new Client({
    user: "postgres",
    password: "psql",
    host: "localhost",
    database: "chat_app",
  });

  client
    .connect()
    .then(() => console.log("Connected successfuly"))
    .catch((e) => console.log(e))
    .finally(() => client.end());
}

test1();
// const { Client } = require("pg");

// const client = new Client({
//   user: "postgres",
//   host: "localhost",
//   database: "testdb",
//   password: "1234abcd",
//   port: 5432,
// });

// client.connect();

// var sender = document.getElementById("sender").innerHTML;
// var msg = document.getElementById("text").innerHTML;
// var date1 = new Date();

// // var x = `INSERT INTO chat_logs (msg_sender, msg_text, timestamp1) VALUES (${
// //   five + ten
// // } and not ${2 * five + ten}.`;

// const query = `INSERT INTO chat_logs (msg_sender, msg_text, timestamp1) VALUES (${sender},${msg},${date1})`;

// client.query(query, (err, res) => {
//   if (err) {
//     console.error(err);
//     return;
//   }
//   console.log("Data insert successful");
//   client.end();
// });
