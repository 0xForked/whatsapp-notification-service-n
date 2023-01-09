const sse = new EventSource("http://localhost:8000/api/v1/stream")
const loaderSection = document.getElementById("loader-section")
const mainSection = document.getElementById("main-section")

setTimeout(() => {
    loaderSection.classList.add('hidden')
    mainSection.classList.remove('hidden')
}, 500)


const qr = document.getElementById("qrcode")
const status = document.getElementById("status")
const logs = document.getElementById("logs")
const qrcode = new QRCode(qr, {
    width: 350,
    height: 350,
    colorDark: "#000000",
    colorLight: "#ffffff",
    correctLevel: QRCode.CorrectLevel.H
});

sse.onmessage = (event) => {
    if (!event.data.includes("logged", "in", "new", "message")) {
        qr.classList.remove("hidden")
        qrcode.clear()
        qrcode.makeCode(event.data.padEnd(220))
    }

    if (event.data.includes("logged", "in")) {
        qr.classList.add("hidden")
        status.innerHTML = `<div class="mt-5">${event.data}</div>`
    }

    if (event.data.includes("new", "message")) {
        qr.classList.add("hidden")
        console.log(event.data.includes("new", "message"))
        logs.innerHTML = `${event.data} <br> ${logs.innerHTML}`
    }
}