const ws = new WebSocket("ws://localhost:8000/api/v1/ws")
const loaderSection = document.getElementById("loader-section")
const mainSection = document.getElementById("main-section")
const qr = document.getElementById("qrcode")
const status = document.getElementById("status")
const account = document.getElementById("account")
const logs = document.getElementById("logs")
const message = document.getElementById("message")
const qrcode = new QRCode(qr, {
    width: 300,
    height: 300,
    colorDark: "#000000",
    colorLight: "#ffffff",
    correctLevel: QRCode.CorrectLevel.H
});

setTimeout(() => {
    loaderSection.classList.add('hidden')
    mainSection.classList.remove('hidden')
}, 500)

ws.onopen = () => sendEvent("status")
ws.onclose = () => closeHandler()
ws.onmessage = (event) => proceedEvent(event)

function sendEvent(action) {
    if (!ws) { closeHandler() }
    ws.send(action)
}

function closeHandler() {
    if (confirm("Websocket connection closed, OK to reload!") === true) {
        window.location.reload()
    }
}

function proceedEvent(event) {
    if (!isJsonString(event.data)) {
        alert("something went wrong when proceed event data!")
        return
    }

    const callback = JSON.parse(event.data)
    if (callback.status === "error") {
        alert(callback.data)
        return
    }

    switch (callback.action) {
        case "qrcode":
            qrcode.clear()
            qr.classList.remove("hidden")
            status.classList.add("hidden")
            logs.classList.add("hidden")
            qrcode.makeCode(callback.data.padEnd(220))
            break
        case "loggedIn":
            qr.classList.add("hidden")
            status.classList.remove("hidden")
            logs.classList.remove("hidden")
            account.innerHTML = callback.data
            break
        case "loggedOut":
            if (confirm("Logged out!") === true) {
                qr.classList.remove("hidden")
                status.classList.add("hidden")
                logs.classList.add("hidden")
                sendEvent("status")
            }
            break
        case "log":
            message.innerHTML = `${callback.data} <br> ${message.innerHTML}`
            break
    }
}

function isJsonString(str) {
    try {
        JSON.parse(str)
        return true
    } catch (e) {
        return false
    }
}