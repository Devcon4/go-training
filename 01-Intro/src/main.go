package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/chat/{id}", getChatRequest).Methods("GET")

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}

// chat: Entity for storing a chat
type chat struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
	IsRead  bool
}

var data = []chat{
	chat{
		ID:      0,
		Message: "Chat-1!",
		IsRead:  false,
	},
	chat{
		ID:      1,
		Message: "Chat-2",
		IsRead:  true,
	},
}

// getChatRequest: Get handler for a chat.
func getChatRequest(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])
	var res *chat

	for i := range data {
		if data[i].ID == id {
			res = &data[i]
			break
		}
	}

	if res == nil {
		http.Error(w, "No item found.", http.StatusNotFound)
		return
	}

	resData, err := json.Marshal(res)

	if err != nil {
		http.Error(w, "Json parse error.", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resData)
}
