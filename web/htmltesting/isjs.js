var net = require('net');

const readline = require('readline').createInterface({
    input: process.stdin,
    output: process.stdout
  });

const client = new net.Socket()

// let isbut = document.getElementById("tap")
// let istext = document.getElementById("messageText").innerHTML

client.connect(1341,  "127.0.0.1", () => {
   // while (true){

    readline.question('write', name => {
        console.log("checking")
        client.write(`${name}\n`)

        client.on("data", (data) => {
            console.log("Got text")
            console.log(data.toString("utf-8"))
        })
        
    readline.close()
    })
//}
})

client.on("close", () => {
    client.end()})



  
  