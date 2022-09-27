package main

import (
	"encoding/json"
	"fmt"
	"log"
	rand2 "math/rand"
	"net/http"
	"os"
)

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

	commands := []string{"F", "T", "T", "T", "T", "T", "T", "T", "R", "L"}
	command_Uturn := []string{"R", "F"}
	var stack_Uturn int := 0
	rand := rand2.Intn(10)//rand는 랜덤 계산 결과값
	
	// TODO add your implementation here to replace the random response
	//코너 빠져나가기
	if stack_Uturn >= 2 {
		stack_Uturn = 0	
	}
	if input.Arena.State["X"] == 0 || input.Arena.State["X"] == input.Arena.Dimensions[len(input.Arena.Dimensions) - 1] || input.Arena.State["Y"] == 0 || input.Arena.State["Y"] == input.Arena.Dimensions[len(input.Arena.Dimensions) - 1] ) {
		return command_Uturn[stack_Utrun++]
	}
	
	return commands[rand]
}
