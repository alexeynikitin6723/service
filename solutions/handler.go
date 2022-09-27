package solutions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"service/tasks/check"
	"service/tasks/cyclic"
	"service/tasks/search"
	"service/tasks/wonderful"
)

type handler struct {
	task Task
}

type dataTask struct {
	A   [][]int `json:"data"`
	Res [][]int `json:"res"`
	K   []int   `json:"k"`
}

func (s *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()
	// ctx = common.WithUserToken(ctx, r.Header.Get(common.TokenHeaderName))

	fmt.Println("request: ", r.Method, r.URL)

	if r.URL.Path == "/tasks/solution" {
		res := DataRequest(CheckTask)
		_, err := w.Write([]byte(res))
		panicIfError(err)
		w.WriteHeader(http.StatusAccepted)
		return
	}

	if r.URL.Path == "/ping" {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("pong"))
		panicIfError(err)
		return
	}

	if r.URL.Path == "/tasks" {
		w.WriteHeader(http.StatusOK)
		res := DataResponse(CheckTask)
		_, err := w.Write([]byte(res))
		panicIfError(err)
		printToConsole(res)
		res = DataResponse(CyclicTask)
		_, err = w.Write([]byte(res))
		panicIfError(err)
		printToConsole(res)
		res = DataResponse(SearchTask)
		_, err = w.Write([]byte(res))
		panicIfError(err)
		printToConsole(res)
		res = DataResponse(WonderfulTask)
		_, err = w.Write([]byte(res))
		panicIfError(err)
		printToConsole(res)
		return
	}
	if r.URL.Path == "/tasks/check" {
		w.WriteHeader(http.StatusOK)
		res := DataResponse(CheckTask)
		_, err := w.Write([]byte(res))
		panicIfError(err)
		printToConsole(res)
		return
	}
	if r.URL.Path == "/tasks/cyclic" {
		w.WriteHeader(http.StatusOK)
		res := DataResponse(CyclicTask)
		_, err := w.Write([]byte(res))
		panicIfError(err)
		printToConsole(res)
		return
	}
	if r.URL.Path == "/tasks/search" {
		w.WriteHeader(http.StatusOK)
		res := DataResponse(SearchTask)
		_, err := w.Write([]byte(res))
		panicIfError(err)
		printToConsole(res)
		return
	}
	if r.URL.Path == "/tasks/wonderful" {
		w.WriteHeader(http.StatusOK)
		res := DataResponse(WonderfulTask)
		_, err := w.Write([]byte(res))
		panicIfError(err)
		printToConsole(res)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	_, err := w.Write([]byte("request url not found"))
	panicIfError(err)
}

func responseIfError(w http.ResponseWriter, err error) bool {
	errorExists := err != nil

	if errorExists {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	return errorExists
}

func panicIfError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func MakeHandler() error {
	return http.ListenAndServe("localhost:3000", &handler{})
}

func GetData(numberTask Task) dataTask {
	absPath, _ := filepath.Abs("./tasks/" + numberTask.GetPath() + "/data.json")
	plan, _ := ioutil.ReadFile(absPath)
	var data dataTask
	err := json.Unmarshal(plan, &data)
	panicIfError(err)
	return data
}

func GetResult(numberTask Task) ([]int, [][]int) {
	data := GetData(numberTask)
	result := make([]int, 0)
	for i := 0; i < len(data.A); i++ {
		switch numberTask.GetPath() {
		case "check":
			result = append(result, check.Solution(data.A[i]))
		case "cyclic":
			result = append(result, cyclic.Solution(data.A[i], data.K[i])...)
		case "search":
			result = append(result, search.Solution(data.A[i]))
		case "wonderful":
			result = append(result, wonderful.Solution(data.A[i]))
		}
	}
	return result, data.A
}

func DataRequest(numberTask Task) []byte {
	result, A := GetResult(numberTask)
	res := Results{
		Payload: A,
		Results: result,
	}
	Req := Request{
		UserName: "alexey4",
		Task:     namesTask[numberTask],
		Results:  res,
	}
	request, err := json.Marshal(Req)
	panicIfError(err)
	return request
}

func DataResponse(numberTask Task) []byte {
	data := GetData(numberTask)
	result, _ := GetResult(numberTask)
	count := 0
	var failArray []Fail
	var dataRes []int
	for _, value := range data.Res {
		for _, val := range value {
			dataRes = append(dataRes, val)
		}
	}
	fmt.Println(dataRes)
	for i, value := range result {
		fail := Fail{
			OriginalResult: dataRes[i],
			ExternalResult: value,
		}
		if value == dataRes[i] {
			count++
		}
		failArray = append(failArray, fail)
	}
	if len(result) == 0 {
		panic("Нет данных")
	}

	percent := float64(count) / float64(len(result)) * 100
	Resp := Response{
		Percent: percent,
		Fails:   failArray,
	}
	request, err := json.Marshal(Resp)
	panicIfError(err)
	return request
}

func printToConsole(res []byte) {
	var data map[string]interface{}
	err := json.Unmarshal(res, &data)
	panicIfError(err)
	fmt.Println(data)
}
