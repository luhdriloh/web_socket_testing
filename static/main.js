var socket = new WebSocket("ws://localhost:8000/socket");
console.log(socket.readyState);

socket.onopen = function () {
	console.log('Web socket has opened');
};

socket.onmessage = function (event) {
  console.log(event.data);
}

