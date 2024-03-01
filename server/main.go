package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/EliasPereyra/youtube-monitoring/server/websocket"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
} 

func showStats(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)

	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	go websocket.Writer(ws)
}

func Routes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/stats", showStats)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("Youtube Monitoring")
	Routes()
}