package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode("Home Page!")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode("is alive!")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type RequestProcessBody struct {
	Process int `json:"process"`
	Worker int `json:"worker"`
}

func Process(w http.ResponseWriter, r *http.Request) {
	var request RequestProcessBody
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}

	data := make(map[string]bool)

	for i := 0; i < request.Process; i++ {
		key := fmt.Sprintf("process_%d", i)
		data[key] = ProcessId(i)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println("Done")
}

func ProcessId(id int) bool {
	fmt.Println("ProcessId", id)
	time.Sleep(2 * time.Second)

	rest := id % 2;
	if (rest == 0)	{
		return true
	} 
	return false
}

func GoProcess(w http.ResponseWriter, r *http.Request) {
	var request RequestProcessBody
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	var wg sync.WaitGroup
	data := make(map[string]bool)
	processChan := make(chan bool)

	for i := 0; i < request.Process; i++ {
		wg.Add(1)
		go GoProcessId(i, processChan, &wg)
	}

	go func() {
		wg.Wait()
		close(processChan)
	}()

	for result := range processChan {
		data[fmt.Sprintf("process_%d", len(data))] = result
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println("Done")
}

func GoProcessId(id int, ch chan<- bool, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("ProcessId", id)
	time.Sleep(2 * time.Second)
	
	rest := id % 2;
	if (rest == 0)	{
		ch <- true
	} else {
		ch <- false
	}
}