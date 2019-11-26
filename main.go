package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Starting pod tester")
	mux1 := http.NewServeMux()
	mux1.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logAndJSONResponse(8080, r, w)
	})
	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logAndJSONResponse(5200, r, w)
	})

	go http.ListenAndServe(":8080", mux1)
	http.ListenAndServe(":5200", mux2)

}

func logAndJSONResponse(port int, r *http.Request, w http.ResponseWriter) {
	data := map[string]interface{}{
		"time": time.Now().String(),
		"port": port,
		"host": r.Host,
		"from": r.RemoteAddr,
	}
	b, _ := json.Marshal(data)
	fmt.Printf("%s: S: %s F:%s\n", data["time"], r.Host, r.RemoteAddr)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
