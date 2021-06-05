const express = require('express');
const WebSocket = require('ws');
const app = express();
const wss = new WebSocket.Server({ port: 1337 });




wss.on('connection', ws => {
    ws.on('message', (data) => {
        try {
            var json = JSON.parse(data)
            if (json.username == "" || json.content == "") return ws.send("")
            wss.clients.forEach(_ws => {
                var time = new Date()
                _ws.send(`${json.username} @ ${time}: ${json.content}`)
            }) 
        } catch (e) {
            ws.send("wrong json/no json")
        }
    })
})

app.get('/', (req, res) => {
    res.sendFile(__dirname+'\\chat.html')
})

app.listen(8080)
