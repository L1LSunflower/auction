const socket = new WebSocket('ws://localhost:4000/auctions/action/18');

// Обработка открытия соединения
socket.addEventListener('open', event => {
    console.log('WebSocket connection opened');
});

// Обработка закрытия соединения
socket.addEventListener('close', event => {
    console.log('WebSocket connection closed');
});

// Обработка получения сообщения от сервера
socket.addEventListener('message', event => {
    console.log('Received message from server:', event.data);
});