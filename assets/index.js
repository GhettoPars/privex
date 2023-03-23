window.onload = () => {
    let menu = document.getElementById("menu");
    let messagebox = document.getElementById("messages");
    let inputMessage = document.getElementById("input");
    let inputBox = document.getElementById("submit");

    inputBox.addEventListener("click", () => { send(inputMessage.value) });

    let chat = document.getElementById("chat");

    class User {
        constructor(uid) {
            let colors = {
                1: "aqua",
                2: "blueviolet"
            }
            this.userID = uid;
            this.color = colors[uid];
            this.tag = document.getElementById("userID");
            //this.tag.style.backgroundColor = this.color;
            //this.tag.innerHTML = (uid + "-User");
        }
    }

    let user = {};
    user = new User(1);

    async function updateChat() {
        let data = await listMessages();
        messagebox.innerHTML = "";
        data.forEach((element, index) => {
            makeMessage(data[(data.length - 1) - index].MessageText, data[(data.length - 1) - index].MessageID);
        });
    }

    function listMessages() {
        return new Promise((resolve, reject) => {
            fetch("https://privex.onrender.com/api/message", { method: "GET" })
                .then(res => res.json())
                .then((list) => { resolve(list); });
        })
    }

    function makeMessage(msg, id) {
        let message = document.createElement("div");
        let del = document.createElement("p");
        let text = document.createElement("p");
        message.id = id;
        //text.style.backgroundColor = user.color;
        message.style.display = "flex";
        message.style.flexDirection = "row";
        message.className = "message";
        text.className = "msg";
        text.innerHTML = msg;
        del.className = "delete";
        del.innerHTML = "❌";
        del.onclick = () => {
            fetch("https://privex.onrender.com/api/message/" + id, { method: "DELETE" })
                .then(function (response) {
                    console.log(response);
                    updateChat();
                })
                .catch(function (error) {
                    console.log(error);
                });
        }
        message.appendChild(del);
        message.appendChild(text);
        messagebox.appendChild(message);
    }

    function send(msg) {
        if (msg != "") {
            // let text = document.createElement("p");
            // text.style.backgroundColor = user.color;
            // text.className = "message";
            // text.innerHTML = msg;
            // messagebox.appendChild(text);
            //makeMessage(msg);
            inputMessage.value = "";

            fetch("https://privex.onrender.com/api/message", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({ UserID: user.userID, MessageText: msg, MessageType: "txt" }),
            })
                .then(function (response) {
                    updateChat();
                    return response.json();
                })
                .catch(function (error) {
                    console.log(error);
                });
        }
    }
    updateChat();
}