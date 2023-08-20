package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var msgChan chan string

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/event", sseHandle)
	r.HandleFunc("/time", timeHandle)

	log.Fatal(http.ListenAndServe(":3500", r))
}
func sseHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	msgChan = make(chan string)
	defer func() {
		close(msgChan)
		msgChan = nil
		fmt.Println(" Client Close connection")
	}()

	flusher, ok := w.(http.Flusher)
	if !ok {
		fmt.Println("cloud not init http.Flusher")
	}
	for {
		select {
		case message := <-msgChan:
			fmt.Printf("data: %s\n", message)
			fmt.Fprintf(w, "data: %s\n\n", message)
			flusher.Flush()
		case <-r.Context().Done():
			fmt.Println(" Client Close connection")
			return
		}
	}
}
func timeHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if msgChan != nil {
		msg := time.Now().Format("15:04:05")
		msgChan <- msg
	}
}
