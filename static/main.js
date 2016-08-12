var socket = new WebSocket("ws://localhost:8000/socket");
console.log(socket.readyState);

socket.onopen = function () {
	console.log('Web socket has opened');
};

socket.onclosed = function () {
	console.log('Web socket has closed');
};

socket.onmessage = function (event) {
	var p = document.createElement("p"); 
	var text = document.createTextNode(event.data);
	p.appendChild(text);
	document.getElementById("body").appendChild(p);
}

document.getElementById("message")
    .addEventListener("keyup", function(event) {
    event.preventDefault();
    if (event.keyCode == 13) {
    	var message = document.getElementById("message").value;
    	document.getElementById("message").value = "";
    	processInput(message)
    }
});

function processInput(message) {
	socket.send(message);
}