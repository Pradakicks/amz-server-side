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
    console.log(req.body)
    console.log(Date.now() - req.body.time)
    em.emit('restock', JSON.stringify(req.body))
    res.send(JSON.stringify(req.body))
})

io.on("connect", (socket) => {

    console.log(`connect ${socket.id}`);

    em.on("restock", data => {
        console.log(`Local Event Emitter Data : ${data}`)
        console.log(Date.now() - JSON.parse(data).time)
        socket.emit("restock", data)

    })

    socket.on("disconnect", () => {
        console.log(`disconnect ${socket.id}`);
    });
    
});

server.listen(port, () => {
    console.log(`application is running at: http://localhost:${port}`);
  });