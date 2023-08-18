package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
  ReadBufferSize: 1024,
  WriteBufferSize: 1024,
}

func main() {
  http.HandleFunc("/echo", echo_handler)
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "index.html")
}

func echo_handler(w http.ResponseWriter, r *http.Request) {
  conn, _ := upgrader.Upgrade(w, r, nil)

  for {
    msgType, msg, err := conn.ReadMessage()
    if err != nil {
      return
    }

    fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

    if err = conn.WriteMessage(msgType, msg); err != nil {
      return
    }
  }
}
