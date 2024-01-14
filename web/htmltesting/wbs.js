var webSocket = new WebSocket("ws://localhost:8080/checkcode");

console.log("is sending connected")
let isbut = document.getElementById("tap")
let text = document.getElementById("senderm")
let istext = document.getElementById("messageText")

webSocket.onopen = function(){
    console.log("Connected to " + webSocket.OPEN)}

webSocket.onmessage = function(event){
    console.log("is got messsage" + event.data)
    istext.innerHTML = event.data    
}

function sendm(){
    console.log("send message: " + text.value)
    webSocket.send(text.value)
    text.value = ""
}


// isbut.addEventListener("click", function(){
//     console.log("is sending")
//     webSocket.send(istext.innerHTML)


    
// })
// }


// webSocket.close();