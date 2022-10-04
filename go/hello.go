package hello

import (
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"

	"golang.org/x/net/websocket"
)

type EventBus interface {
	SendEvent(channel, message string)
}

func StartListening() {
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(time.Second * 2)
			fmt.Fprintln(w, "Hello Http")

		})
		http.Handle("/websocket", websocket.Handler(onWebSocket))

		http.ListenAndServe(":8080", nil)

	}()
}
func onWebSocket(ws *websocket.Conn) {
	message := "Hello Web Socket"
	websocket.Message.Send(ws, message)

}
func WriteFile(baseDir string) {
	d1 := []byte("hello go\n")
	file := filepath.Join(baseDir, "test.txt")
	ioutil.WriteFile(file, d1, 0644)
}
func ReadFile(baseDir string) string {
	file := filepath.Join(baseDir, "test.txt")
	bytes, _ := ioutil.ReadFile(file)
	return string(bytes)
}

func HelloWorld(bus EventBus) string {
	go func() {
		time.Sleep(time.Second * 3)
		bus.SendEvent("general", "After 3 seconds!")
		time.Sleep(time.Second * 3)
		bus.SendEvent("general", "After 6 seconds!")
	}()
	return "Hello Native Bridge"
}
