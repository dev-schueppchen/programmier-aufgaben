/**
 * Shortcut function for getting a DOM element by query.
 * @param {string} query 
 * @returns {Element} the DOM element
 */
var $ = (query) => document.querySelector(query);

/**
 * Class like object contructor for a web socket
 * client with registered event handlers.
 * @class
 * @param {stirng} url WS server URL
 */
function WsClient(url) {
    // Create web socket client
    this.ws = new WebSocket(url);
    this.events = {};

    /**
     * Register an event handler by name of the event.
     * @param {stirng} name Name of the event
     * @param {function} handler Event handler which is getting
     *                           passed the event data.
     */
    this.on = (name, handler) => this.events[name] = handler;

    /**
     * Send an event with the name of the event and
     * the data the event contains.
     * @param {stirng} name Name of the event
     * @param {object} data Data of the event
     */
    this.send = (name, data) => {
        let raw = JSON.stringify({name, data});
        this.ws.send(raw);
    };

    // If a message was received by the web socket client,
    // try to parse the message content with the JSON parser
    // to an event object.
    // If this was successful, search for the name in the
    // registered event handlers and execute the corresponding
    // handler, if defined.
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

// Open wbe socket connection to server.
// The URL is used from the windows location URL.
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