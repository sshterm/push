package main

import (
	_ "embed"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/token"
)

//go:embed config/kid
var kid []byte

//go:embed config/iss
var iss []byte

//go:embed config/key.p8
var key []byte

func main() {
	authKey, err := token.AuthKeyFromBytes(key)
	if err != nil {
		log.Fatal("token error:", err)
	}
	token := &token.Token{
		AuthKey: authKey,
		KeyID:   string(kid),
		TeamID:  string(iss),
	}

	topic := []string{"cn.sshterm.pro", "cn.sshterm.free", "cn.sshterm.dev"}

	client := apns2.NewTokenClient(token).Development() //.Production()
	http.HandleFunc("/apn_push", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Please send a request body", http.StatusBadRequest)
			return
		}
		var data Body
		err = json.Unmarshal(body, &data)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if data.Token == "" || data.Topic == "" || data.Alert == "" {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		var valid bool
		for _, v := range topic {
			if strings.Compare(data.Topic, v) == 0 {
				valid = true
				break
			}
		}
		if !valid {
			http.Error(w, "Invalid topic", http.StatusBadRequest)
			return
		}

		notification := &apns2.Notification{
			DeviceToken: data.Token,
			Topic:       data.Topic,
			Payload:     []byte(`{"aps":{"alert":"` + data.Alert + `"}}`),
		}
		res, err := client.Push(notification)
		if err != nil {
			http.Error(w, "Error sending push notification", 500)
		} else {
			json.NewEncoder(w).Encode(res)
		}
	})
	http.ListenAndServe(":8080", nil)
}

type Body struct {
	Token string `json:"token"`
	Topic string `json:"topic"`
	Alert string `json:"alert"`
}