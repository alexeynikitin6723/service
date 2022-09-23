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

func (t Task) GetName() string {
	return namesTask[t]
}
