const socket = new WebSocket('ws://'+ window.location.host+'/ws')
socket.onopen = function() {
    console.log('WebSocket is connected.')
}

form = document.getElementById("form")
form.addEventListener('submit', SendForm)

function SendForm(event){
    event.preventDefault()
    console.log('Нажата клавиша запуск')

    let data = new FormData(form)
    data.get("season")

    let jsonStruct = {
        count: data.get("count"),
        sun: data.get("sun"),
        sea: data.get("sea"),
        age: data.get("age"),
        energy: data.get("energy")
    }
    let jsonString = JSON.stringify(jsonStruct)
    console.log(jsonString)

    socket.send(jsonString)
}

socket.onmessage = function (event){
    console.log(event.data);
}

