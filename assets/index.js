window.onload = () => {
    let menu = document.getElementById("menu");
    let messagebox = document.getElementById("messages");
    let inputMessage = document.getElementById("input");
    // let inputbox = document.getElementById('inputbox');
    let inputBox = document.getElementById("submit");
    let updatelist = document.getElementById("update");

    // inputbox.addEventListener("submit", () => { send(inputMessage.value) })
    inputBox.addEventListener("click", () => { send(inputMessage.value) });
    updatelist.addEventListener("click", () => { updateChat() });

    let login = document.getElementById("login");
    let chat = document.getElementById("chat");
    let user1 = document.getElementById("user1");
    let user2 = document.getElementById("user2");

    class User {
        constructor(uid) {
            let colors = {
                1: "aqua",
                2: "blueviolet"
            }
            this.userID = uid;
            this.color = colors[uid];
            this.tag = document.getElementById("userID");
            this.tag.style.backgroundColor = this.color;
            this.tag.innerHTML = (uid + "-User");
        }
    }

    let user = {};

    user1.onclick = () => { user = new User(1); hideLoginScreen(); }
    user2.onclick = () => { user = new User(2); hideLoginScreen(); }

    function hideLoginScreen() {
        login.style.display = "none";
        chat.style.display = "flex";
        chat.style.flexDirection = "column";
    }

    async function updateChat() {
        let data = await listMessages();
        data.forEach(element => {
            makeMessage(element.MessageText);
        });
    }

    function listMessages() {
        return new Promise((resolve, reject) => {
            fetch("/api/message", { method: "GET" })
                .then(res => res.json())
                .then((list) => { resolve(list); });
        })
    }

    function makeMessage(msg) {
        let text = document.createElement("p");
        text.style.backgroundColor = user.color;
        text.className = "message";
        text.innerHTML = msg;
        messagebox.appendChild(text);
    }

    function send(msg) {
        if (msg != "") {
            // let text = document.createElement("p");
            // text.style.backgroundColor = user.color;
            // text.className = "message";
            // text.innerHTML = msg;
            // messagebox.appendChild(text);
            makeMessage(msg);
            inputMessage.value = "";

            fetch("/api/message", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({ UserID: user.userID, MessageText: msg, MessageType: "txt" }),
            })
                .then(function (response) {
                    return response.json();
                })
                .catch(function (error) {
                    console.log(error);
                });
        }
    }
}