var $ = (query) => document.querySelector(query);

function WsClient(url) {
    this.ws = new WebSocket(url);
    this.events = {};

    this.on = (name, handler) => this.events[name] = handler;

    this.send = (name, data) => {
        let raw = JSON.stringify({name, data});
        this.ws.send(raw);
    };

    this.ws.onmessage = (res) => {
        console.log(res);
        try {
            let data = JSON.parse(res.data);
            if (data) {
                let cb = this.events[data.name];
                if (cb)
                    cb(data.data);
            }
        } catch (err) {
            console.log(err);
        }
    };
}

var ws = new WsClient(
    window.location.href.replace(/((http)|(https)):\/\//gm, 'ws://') + 'ws');

$('#btSend').onclick = () => {
    let val = $('#tbSend').value;
    let nick = $('#tbName').value;
    if (!val || val.length < 1 || !nick || nick.length < 1)
        return;

    ws.send("message", {
        sender: nick,
        content: val,
        timestamp: Date.now(),
    });
};

ws.on("message", (data) => {
    let p = document.createElement('p');
    p.innerText = `${data.sender} @ ${new Date(data.timestamp)}: ${data.content}`
    $('#msgs').appendChild(p);
});