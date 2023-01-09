// const sse = new EventSource("http://localhost:8000/api/v1/sse")
// const ws = new WebSocket("ws://localhost:8000/api/v1/ws")

// SERVER SENT EVENT
// const sseContainer = document.getElementById("sse-container")
// sse.onmessage = (event) => {
//     sseContainer.innerHTML = `${event.data}</br>${sseContainer.innerHTML}`
// }

// WEBSOCKET
// const wsContainer = document.getElementById("ws-container")
// ws.onopen = () => { ws.send(JSON.stringify({'action': "stream"})) }
// ws.onclose = () => console.log("ws closed")
// ws.onmessage = (event) => {
//     const message = JSON.parse(event.data)
//     wsContainer.innerHTML = `${message.data}</br>${wsContainer.innerHTML}`
// }

const loaderSection = document.getElementById("loader-section")
const mainSection = document.getElementById("main-section")

setTimeout(() => {
    loaderSection.classList.add('hidden')
    mainSection.classList.remove('hidden')
}, 1000)