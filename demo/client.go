package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/load"
	"github.com/shirou/gopsutil/v4/mem"
	"gopkg.in/yaml.v3"
)

func main() {
	var config Config
	data, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := load.Avg()
		if err != nil {
			log.Fatal(err)
		}
		if c.Load1 > config.Load {
			title := fmt.Sprintf("Server load is %.2f", c.Load1)
			subtitle := config.Subtitle

			body := ""
			if v, err := cpu.Percent(time.Second, false); err == nil && len(v) > 0 {
				body += fmt.Sprintf("CPU usage is %.2f %%", v[0])
			}
			if v, err := mem.VirtualMemory(); err == nil {
				body += fmt.Sprintf(",Memory usage is %.2f %%", v.UsedPercent)
			}
			log.Println(title, subtitle, body)
			push(title, subtitle, body, config.Token, config.Topic, config.Priority, config.Node, config.Dev)
		}
		time.Sleep(time.Minute)
	}
}

type Config struct {
	Topic    string  `yaml:"topic"`
	Token    string  `yaml:"token"`
	Node     string  `yaml:"node"`
	Load     float64 `yaml:"load"`
	Priority int     `yaml:"priority"`
	Dev      bool    `yaml:"dev"`
	Subtitle string  `yaml:"subtitle"`
}

func push(title, subtitle, body, token, topic string, priority int, url string, dev bool) (resBody *Response, err error) {
	var data []byte
	data, err = json.Marshal(Body{
		Dev:   dev,
		Token: token,
		Topic: topic,
		Notification: Notification{
			APS: APS{
				Alert: Alert{
					Title:    title,
					Subtitle: subtitle,
					Body:     body,
				},
			},
		},
		Priority: priority,
	})
	if err != nil {
		return
	}

	var res *http.Response

	res, err = http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&resBody)

	return
}

type Body struct {
	Dev          bool         `json:"dev"`
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

type Response struct {

	// The HTTP status code returned by APNs.
	// A 200 value indicates that the notification was successfully sent.
	// For a list of other possible status codes, see table 6-4 in the Apple Local
	// and Remote Notification Programming Guide.
	StatusCode int

	// The APNs error string indicating the reason for the notification failure (if
	// any). The error code is specified as a string. For a list of possible
	// values, see the Reason constants above.
	// If the notification was accepted, this value will be "".
	Reason string

	// The APNs ApnsID value from the Notification. If you didn't set an ApnsID on the
	// Notification, this will be a new unique UUID which has been created by APNs.
	ApnsID string

	// If the value of StatusCode is 410, this is the last time at which APNs
	// confirmed that the device token was no longer valid for the topic.
	Timestamp time.Time

	// An identifier that is only available in the Developement enviroment. Use
	// this to query Delivery Log information for the corresponding notification
	// in Push Notifications Console.
	ApnsUniqueID string
}
