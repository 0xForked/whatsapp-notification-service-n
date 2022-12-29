let stream = new EventSource("/api/v1/stream")

// ID REGISTRAR
// const container = document.getElementById("container")
const events = document.getElementById("events")

stream.addEventListener("message", function (event) {
    events.innerHTML = `${event.data}</br>` + events.innerHTML
})