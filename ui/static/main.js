const username = prompt("Please enter your name") || "Don Pendejo"
const url = new URL(`./ws/?username=${username}`, location.href)
url.protocol = url.protocol.replace("http", "ws")

const socket = new WebSocket(url)

socket.onmessage = (e) => {
    try {
        const data = JSON.parse(e.data)

        switch(data.event) {
            case 0:
                listUsers(data.users)
                break
            case 1:
                addMessage(data)
                break
            default:
                console.log("Unknown event")
                break
        }
    } catch (err) {
        console.error("Failed to parse message:", err)
    }
}

function listUsers(users) {
    const userList = document.getElementById("users")
    userList.replaceChildren()

    for (const u of users) {
        const listItem = document.createElement("li")
        listItem.textContent = u
        userList.appendChild(listItem)
    }
}

function addMessage(message) {
    const template = document.getElementById("message");
    const clone = template.content.cloneNode(true);
    clone.querySelector("span").textContent = message.username;
    clone.querySelector("p").textContent = message.content;
    document.getElementById("chat").append(clone);
}


const inputElement = document.getElementById("msg")
const form = document.getElementById("form")
form.onsubmit = (e) => {
    e.preventDefault()
    if (!socket) return
    if (!inputElement.value) return
    
    const msg = {event: 1, username: username, content: inputElement.value}
    socket.send(JSON.stringify(msg))

    inputElement.value = ""
}
