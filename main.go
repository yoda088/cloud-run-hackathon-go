package main

import (
	"encoding/json"
	"fmt"
	"log"
	rand2 "math/rand"
	"net/http"
	"os"
)

var i int = -1
var rand int = 0

func main() {
	port := "8080"
	if v := os.Getenv("PORT"); v != "" {
		port = v
	}
	http.HandleFunc("/", handler)

	log.Printf("starting server on port :%s", port)
	err := http.ListenAndServe(":"+port, nil)
	log.Fatalf("http listen error: %v", err)
}

func handler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		fmt.Fprint(w, "Let the battle begin!")
		return
	}

	var v ArenaUpdate
	defer req.Body.Close()
	d := json.NewDecoder(req.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&v); err != nil {
		log.Printf("WARN: failed to decode ArenaUpdate in response body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := play(v)
	fmt.Fprint(w, resp)
}

func play(input ArenaUpdate) (response string) {
	log.Printf("IN: %#v", input)

	//command := []string{"F", "T", "T", "T", "T", "T", "R", "L"}
	commands := [][]string{{"R", "F", "T", "L", "F", "T", "L", "F", "T", "R"}, {"F", "T", "F", "T"},
		{"T", "R", "T", "R", "T", "R", "T"}, {"T", "L", "T", "L", "T", "L", "T"}, {"R", "R", "F", "F"},
		{"F", "F", "F"}, {"R", "F", "T"}, {"L", "F", "T"}, {"T", "T", "T", "T", "T"}, {"T", "R", "T", "R", "T", "R", "T", "R"}}
	if i >= (len(commands[rand]) - 1) {
		i = -1
	}
	if i == -1 {
		rand = rand2.Intn(10) //rand는 랜덤 계산 결과값
	}
	i++
	return commands[rand][i]
}
