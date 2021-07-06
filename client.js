const { io } = require("socket.io-client") 

// const socket = io("ws://localhost:3030/", {});
const socket = io("ws://159.203.179.167:3030/", {});
https:///
socket.on("connect", () => {
    console.log(`connect ${socket.id}`);
});

socket.on("disconnect", () => {
    console.log(`disconnect`);
});

// setInterval(() => {
//     const start = Date.now();
//     socket.emit("ping", () => {
//         console.log(`pong (latency: ${Date.now() - start} ms)`);
//     });
// }, 1000);

socket.on("restock", (data) => {  
    console.log(JSON.parse(data))
})