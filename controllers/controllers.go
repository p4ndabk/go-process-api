package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"golang.org/x/crypto/bcrypt"
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

	log.Println("is alive!")

	err := json.NewEncoder(w).Encode("is alive!")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type RequestProcessBody struct {
	Process int `json:"process"`
	Worker  int `json:"worker"`
}

func Process(w http.ResponseWriter, r *http.Request) {
	var request RequestProcessBody
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}

	data := make(map[string]string)

	for i := 0; i < request.Process; i++ {
		key := fmt.Sprintf("process_%d", i)
		data[key] = ProcessHashById(i)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println("Done")
}

func ProcessHashById(id int) string {
	key := fmt.Sprintf("Mudar@123-%s", strconv.Itoa(id))

	hash, err := bcrypt.GenerateFromPassword([]byte(key), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Erro ao gerar o hash bcrypt:", err)
		panic(err)
	}

	return string(hash)
}

func GoProcess(w http.ResponseWriter, r *http.Request) {
	var request RequestProcessBody
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	var wg sync.WaitGroup
	data := make(map[string]string)
	processChan := make(chan string)

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

func GoProcessId(id int, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- ProcessHashById(id)
}
