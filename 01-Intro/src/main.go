package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	server := http.Server{
		Addr: ":3000",
	}

	http.HandleFunc("/chat/", getChatRequest)
	log.Fatal(server.ListenAndServe())
}

// chat: Entity for storing a chat
type chat struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

var data = []chat{
	chat{
		ID:      0,
		Message: "Chat-1!",
	},
	chat{
		ID:      1,
		Message: "Chat-2",
	},
}

// getChatRequest: Get handler for a chat.
func getChatRequest(w http.ResponseWriter, r *http.Request) {

	reqParsed, err := strconv.ParseInt(strings.TrimPrefix(r.URL.Path, "/chat/"), 10, 32)

	if err != nil {
		http.Error(w, "Unable to parse request id.", http.StatusInternalServerError)
		return
	}

	reqID := int(reqParsed)
	var res *chat

	for i := range data {
		if data[i].ID == reqID {
			res = &data[i]
			break
		}
	}

	if res == nil {
		http.Error(w, "No item found.", http.StatusNotFound)
		return
	}

	json, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
