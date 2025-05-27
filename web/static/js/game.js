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

let canvas = document.getElementById("gameCanvas")
let ctx = canvas.getContext('2d')

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
                //ctx.fillStyle = selectColor(map[y][x].agent.energy, "en")
                ctx.fillStyle = 'green'
                ctx.fillRect(x*cellSizeX, y*cellSizeY, cellSizeX, cellSizeY)
            } else {
                if (map[y][x].height<=sea){
                    ctx.fillStyle = selectColor(map[y][x].height, "sea")
                } else {
                    ctx.fillStyle = selectColor(map[y][x].height, "bw")
                }
                ctx.fillRect(x*cellSizeX, y*cellSizeY, cellSizeX, cellSizeY)
            }

        }
    }
}


function selectColor(num, type){
    switch (type){
        case "bw":
            return "rgb(" + [num*10, num*10, num*10].join(",") + ")";
        case "en":
            if (num <100){
                return "#707000"
            } else if (num <250){
                return "#a3a300"
            } else if (num < 500){
                return "#ffff0a"
            } else {
                return "#FFFFc0"
            }
        case "sea":
            return "rgb(" + [0, 0, num*10].join(",") + ")";
    }

}


