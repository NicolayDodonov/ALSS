// создаём соединение по web-socket при заходе на страницу
const socket = new WebSocket('ws://'+ window.location.host+'/ws')
socket.onopen = function() {
    console.log('WebSocket is connected.')
}

form = document.getElementById("form")
form.addEventListener('submit', SendForm)

//обработка отправки сообщения
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

// получаем объект canvas для рисования кадра
let canvas = document.getElementById("gameCanvas")
let ctx = canvas.getContext('2d')

// при входящем сообщении по web-socket выполняем команду
socket.onmessage = function (event){
    let message = JSON.parse(event.data)
    console.log(message)

    map = message.map.cells
    sea = message.map.sea_level
    year = message.stat.world_year

    cellSizeX = 600/message.map.x_size
    cellSizeY = 600/message.map.y_size
    for(y = 0; y<map.length; y++){
        for(x=0;x<map[y].length; x++){

            if (map[y][x].agent !== null){
                ctx.fillStyle = agent(map[y][x].agent,"age")
                ctx.fillRect(x*cellSizeX, y*cellSizeY, cellSizeX, cellSizeY)
            } else {
                ctx.fillStyle = ground(map[y][x], sea, "mineral")
                ctx.fillRect(x*cellSizeX, y*cellSizeY, cellSizeX, cellSizeY)
            }

        }
    }
}


function agent(agent, type){
    switch (type){
        case "ration":
            switch (agent.Ration){
                case "1":   // охота
                    return "#dd0020"
                case "2":   // фотосинтез
                    return "#00dd20"
                case "3":   // хемосинтез
                    return "#20dddd"
                case "4":   // очистка атмосферы
                    return "#dd20dd"
                default:    // ничего не делал
                    return "#777777"
            }
        case "energy":
            num = agent.Energy
            if (num <100){
                return "#707000"
            } else if (num <250){
                return "#a3a300"
            } else if (num < 500){
                return "#ffff00"
            } else {
                return "#FFFF99"
            }
        case "age":
            num = agent.Age
            if (num <100){
                return "#FFc0c0"
            } else if (num <250){
                return "#ff5a5a"
            } else if (num < 500){
                return "#a30000"
            } else {
                return "#700000"
            }
    }
}

function ground(cell, sea, type){
    switch (type){
        case "height":
            num = cell.height
            if (num > sea){
                return "rgb(" + [num*10, num*10, num*10].join(",") + ")";
            } else {
                return "rgb(" + [0, 0, num*10].join(",") + ")";
            }
        case "mineral":
            num = cell.mineral
            if (num < 10){
                return "#cceeff"
            } else if (num <64){
                return "#99ddff"
            } else if (num <128){
                return "#55ccff"
            } else if (num < 192) {
                return "#10aaff"
            } else if (num < 200){
                return "#0080ff"
            } else {
                return "#5f00ff"
            }
    }
}

