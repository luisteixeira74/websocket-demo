package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Permitir qualquer origem para testes
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Erro ao fazer upgrade:", err)
		return
	}
	defer conn.Close()

	log.Println("Cliente conectado!")

	// Enviar a hora a cada 5 segundos
	for {
		currentTime := time.Now().Format("15:04:05")
		err := conn.WriteMessage(websocket.TextMessage, []byte("Hora atual: "+currentTime))
		if err != nil {
			log.Println("Erro ao enviar mensagem:", err)
			break
		}
		time.Sleep(5 * time.Second)
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/ws", wsHandler)

	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
