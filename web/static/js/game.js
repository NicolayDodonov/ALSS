// создаём соединение по web-socket при заходе на страницу
const socket = new WebSocket('ws://'+ window.location.host+'/ws')
socket.onopen = function() {
    console.log('WebSocket is connected.')
}

//добавляем обработчик игрового меню
form = document.getElementById("form-game-menu")
let button = document.getElementById("start-button");
button.disabled = false
form.addEventListener('submit', SendForm)

//обработка отправки сообщения
function SendForm(event){
    event.preventDefault()
    button.disabled = true
    button.textContent = "Модель запущена"
    console.log('Нажата клавиша запуск')

    let jsonStruct = {
        count:  parseInt(document.getElementById("countAgent").value),
        sun: parseInt(document.getElementById("startSun").value),
        sea: parseInt(document.getElementById("seaLevel").value),
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
                if (renderAgent == 4){
                    ctx.fillStyle = ground(map[y][x], sea)
                    button2.textContent="Режим Агентов - не рисовать"
                } else {
                    ctx.fillStyle = agent(map[y][x].agent)

                }
                ctx.fillRect(x*cellSizeX, y*cellSizeY, cellSizeX, cellSizeY)
            } else {
                ctx.fillStyle = ground(map[y][x], sea)
                ctx.fillRect(x*cellSizeX, y*cellSizeY, cellSizeX, cellSizeY)
            }

        }
    }
}

let renderAgent = 1
function RenderType(){
    renderAgent++
    if (renderAgent > 4){
        renderAgent = 1
    }
}
let renderFon = 1
function RenderFon(){
    renderFon++
    if (renderFon > 2){
        renderFon = 1
    }
}

let button2 = document.getElementById("bt-agent");
button2.disabled = true
let button3 = document.getElementById("bt-fon");
button3.disabled = true
function agent(agent){
    button2.disabled = false
    switch (renderAgent){
        case 1:
            button2.textContent="Режим Агентов - питание"
            switch (agent.Ration){
                case 1:   // охота
                    return "#dd0020"
                case 2:   // фотосинтез
                    return "#00dd20"
                case 3:   // минерализация
                    return "#10aaff"
                case 4:   // хемосинтез
                    return "#dd20dd"
                default:    // ничего не делал
                    return "#777777"
            }
        case 2:
            button2.textContent="Режим Агентов - энергия"
            num = agent.Energy
            if (num <100){
                return "#707000"
            } else if (num <450){
                return "#a3a300"
            } else if (num < 850){
                return "#ffff00"
            } else {
                return "#FFFF99"
            }
        case 3:
            button2.textContent="Режим Агентов - возраст"
            num = agent.Age
            if (num <100){
                return "#FFc0c0"
            } else if (num <450){
                return "#ff5a5a"
            } else if (num < 850){
                return "#a30000"
            } else {
                return "#700000"
            }

    }
}

function ground(cell, sea){
    button3.disabled = false
    switch (renderFon){
        case 1:
            button3.textContent = "Режим фона - высота"
            num = cell.height
            if (cell.mineral >= 225) {
                return "#5f00ff"
            }
            if (num > sea){
                return "rgb(" + [num*10, num*10, num*10].join(",") + ")";
            } else {
                return "rgb(" + [0, 0, num*10].join(",") + ")";
            }
        case 2:
            button3.textContent = "Режим фона - минералы"
            num = cell.mineral
            if (num < 10){
                return "#cceeff"
            } else if (num <64){
                return "#99ddff"
            } else if (num <128){
                return "#55ccff"
            } else if (num < 192) {
                return "#10aaff"
            } else if (num < 225){
                return "#0080ff"
            } else {
                return "#5f00ff"
            }
    }
}

