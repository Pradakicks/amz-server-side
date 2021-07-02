const { json } = require("express");
var bodyParser = require('body-parser');
const app = require("express")(json());
const server = require("http").createServer(app);
const port = process.env.PORT || 3030;

const io = require('socket.io')(server);

// const { Server } = require("socket.io")
// const io = new Server(8080);

// Local Event Emitter

var events = require('events');

var em = new events.EventEmitter()



app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

app.post('/', (req, res) => {
    res.setHeader('Content-Type', 'application/json');
    em.emit('restock', JSON.stringify(req.body))
    res.send("Good")
})

io.on("connect", (socket) => {

    console.log(`connect ${socket.id}`);

    socket.on("ping", (cb) => {
        console.log("ping");
        cb();
    });

    em.on("restock", data => {
        console.log(`Local Event Emitter Data : ${data}`)
        socket.emit("restock", data)
    })
    // setInterval(() => {
    //     socket.emit("restock", {
    //         pid : Date.now(),
    //         offerId : "1234"
    //     });
    // }, 1000)
    

    socket.on("disconnect", () => {
        console.log(`disconnect ${socket.id}`);
    });
    
});

server.listen(port, () => {
    console.log(`application is running at: http://localhost:${port}`);
  });