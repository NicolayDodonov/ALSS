const socket = new WebSocket('ws://'+ window.location.host+'/ws')
socket.onopen = function() {
    console.log('WebSocket is connected.')
}

form = document.getElementById("form")
form.addEventListener('submit', SendForm)

function SendForm(event){
    event.preventDefault()
    console.log('Нажата клавиша запуск')



    let jsonStruct = {
        count:  parseInt(document.getElementById("countAgent").value),
        sun: parseInt(document.getElementById("startSun").value),
        sea: parseInt(document.getElementById("seaLevel").value),
        age: parseInt(document.getElementById("maxAge").value),
        energy: parseInt(document.getElementById("maxEnergy").value)
    }
    let jsonString = JSON.stringify(jsonStruct)
    console.log(jsonString)

    socket.send(jsonString)
}

socket.onmessage = function (event){
    let message = JSON.parse(event.data)
    console.log('We get frame!')

    map = message.map.cells
    console.log(map)

    for(y = 0; y<map.length; y++){
        for(x=0;x<map[y].length; x++){
            drawCell(x, y, map[y][x])
        }
    }
}

function drawCell(x,y, cell){

}
