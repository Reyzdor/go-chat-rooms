let ws = null;

function createRoom() {
    fetch('/create')
    .then(res => res.json())
    .then(data => {
        document.getElementById('token').value = data.token;
    });
}

function joinRoom() {
    const nick = document.getElementById('nick').value;
    const token = document.getElementById('token').value;

    ws = new WebSocket(`ws://localhost:8080/ws?nick=${nick}&token=${token}`);

    ws.onmessage = (event) => {
        addMessage(event.data, 'user-message');
    };

    ws.onopen = () => {
        document.getElementById('form').style.display = 'none';
        document.getElementById('chat').classList.add('active');
        addMessage('Вы подключились к комнате', 'system-message');
    };

    ws.onclose = () => {
        addMessage('Вы отключились от комнаты', 'system-message');
    };
}

function sendMessage() {
    const input = document.getElementById('message');
    if (input.value.trim() !== '') {
        ws.send(input.value);
        input.value = '';
    }
}

function addMessage(text, className = 'user-message') {
    const div = document.createElement('div');
    div.textContent = text;
    div.className = className;
    document.getElementById('messages').appendChild(div);
    
    const messages = document.getElementById('messages');
    messages.scrollTop = messages.scrollHeight;
}
