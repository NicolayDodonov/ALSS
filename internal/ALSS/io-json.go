package ALSS

// FrameJSON это кадр состояния модели, который ALSS.IO отправляет на клиент для обработки в изображение и сопутствующую
// информацию. TODO: Чисто технически можно использовать для сохранения и загрузки модели.
type FrameJSON struct {
	World  WorldJSON
	Map    MapJson
	Agents AgentsJSON
}

type WorldJSON struct {
	Temp int
	//todo: other stat info
}

type MapJson struct {
	X   int
	Y   int
	Map []CellJson
}

type CellJson struct {
	//todo: cell info
}

type AgentsJSON struct {
	Count  int
	agents []AgentJson
}

type AgentJson struct {
	ID    int
	HP    int
	Age   int
	X     int
	Y     int
	Color Color //todo: IDN need it or now...
}

type Color struct {
	R, G, B int
}
