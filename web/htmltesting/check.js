(function(){function r(e,n,t){function o(i,f){if(!n[i]){if(!e[i]){var c="function"==typeof require&&require;if(!f&&c)return c(i,!0);if(u)return u(i,!0);var a=new Error("Cannot find module '"+i+"'");throw a.code="MODULE_NOT_FOUND",a}var p=n[i]={exports:{}};e[i][0].call(p.exports,function(r){var n=e[i][1][r];return o(n||r)},p,p.exports,r,e,n,t)}return n[i].exports}for(var u="function"==typeof require&&require,i=0;i<t.length;i++)o(t[i]);return o}return r})()({1:[function(require,module,exports){

},{}],2:[function(require,module,exports){
var net = require('net');

const client = new net.Socket()

let isbut = document.getElementById("tap")
let istext = document.getElementById("messageText").innerHTML

client.connect(1200,  "127.0.0.1", () => {
    isbut.addEventListener("click", function(){
        alert("checking")
        client.write(`${istext}\n`)

        client.on("data", (data) => {
            console.log(data.toString("utf-8"))
            istext.innerHTML = data
        })
        
        reader.on("close", () => {
            client.end()})
    })
})
},{"net":1}]},{},[2]);
