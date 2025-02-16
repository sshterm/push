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
		if data.Token == "" || data.Topic == "" || data.Notification.APS.Alert.Title == "" || data.Notification.APS.Alert.Subtitle == "" {
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
		if data.Priority > 10 || data.Priority < 5 {
			http.Error(w, "Invalid priority", http.StatusBadRequest)
			return
		}

		notification := &apns2.Notification{
			DeviceToken: data.Token,
			Topic:       data.Topic,
			Payload:     data.Notification,
			PushType:    apns2.PushTypeAlert,
			Priority:    data.Priority,
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
	Token        string       `json:"token"`
	Topic        string       `json:"topic"`
	Notification Notification `json:"notification"`
	Priority     int          `json:"priority"`
}
type APS struct {
	Alert Alert `json:"alert"`
}
type Alert struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Body     string `json:"body"`
}
type Notification struct {
	APS APS `json:"aps"`
}
