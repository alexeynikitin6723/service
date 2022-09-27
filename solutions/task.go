package solutions

type Task int

const CheckTask Task = 0
const CyclicTask Task = 1
const SearchTask Task = 2
const WonderfulTask Task = 3

var namesTask = map[Task]string{
	CheckTask:     "Проверка последовательности",
	CyclicTask:    "Циклическая ротация",
	SearchTask:    "Поиск отсутствующего элемента",
	WonderfulTask: "Чудные вхождения в массив",
}
var pathsTask = map[Task]string{
	CheckTask:     "check",
	CyclicTask:    "cyclic",
	SearchTask:    "search",
	WonderfulTask: "wonderful",
}

func (t Task) GetName() string {
	return namesTask[t]
}
func (t Task) GetPath() string {
	return pathsTask[t]
}

type Request struct {
	UserName string  `json:"user_name"`
	Task     string  `json:"task"`
	Results  Results `json:"results"`
}

type Results struct {
	Payload [][]int `json:"payload"`
	Results []int   `json:"results"`
}

type Response struct {
	Percent float64 `json:"percent"`
	Fails   []Fail  `json:"fails"`
}

type Fail struct {
	OriginalResult int
	ExternalResult int
}
