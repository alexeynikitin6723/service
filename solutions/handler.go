package solutions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

type handler struct {
	task Task
}

func (s *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()
	// ctx = common.WithUserToken(ctx, r.Header.Get(common.TokenHeaderName))

	fmt.Println("request: ", r.Method, r.URL)
	if r.URL.Path == "/ping" {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("pong"))
		panicIfError(err)
		return
	}
	if r.URL.Path == "/tasks/check" {
		w.WriteHeader(http.StatusOK)
		absPath, _ := filepath.Abs("./tasks/check/data.json")
		plan, _ := ioutil.ReadFile(absPath)
		var data interface{}
		err := json.Unmarshal(plan, &data)
		panicIfError(err)
		fmt.Println(data)
		_, err = w.Write([]byte(namesTask[0]))
		panicIfError(err)
		return
	}
	if r.URL.Path == "/tasks/cyclic" {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(namesTask[1]))
		panicIfError(err)
		return
	}
	if r.URL.Path == "/tasks/search" {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(namesTask[2]))
		panicIfError(err)
		return
	}
	if r.URL.Path == "/tasks/wonderful" {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(namesTask[3]))
		panicIfError(err)
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
