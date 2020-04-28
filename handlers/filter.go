package handlers

import (
	"encoding/json"
)

//PBResponse is pushbullet generic response
type PBResponse struct {
	Type    string   `json:"type"`
	Targets []string `json:"targets"`
	Push    struct {
		Type             string `json:"type"`
		SourceDeviceIden string `json:"source_device_iden"`
		SourceUserIden   string `json:"source_user_iden"`
		ClientVersion    int    `json:"client_version"`
		Dismissible      bool   `json:"dismissible"`
		Icon             string `json:"icon"`
		Title            string `json:"title"`
		Body             string `json:"body"`
		ApplicationName  string `json:"application_name"`
		PackageName      string `json:"package_name"`
		NotificationID   string `json:"notification_id"`
		Notifications    []struct {
			ThreadID  string `json:"thread_id"`
			Title     string `json:"title"`
			Body      string `json:"body"`
			Timestamp int    `json:"timestamp"`
		} `json:"notifications"`
		Actions []struct {
			Label      string `json:"label"`
			TriggerKey string `json:"trigger_key"`
		} `json:"actions"`
	} `json:"push"`
}

//Filter (s) out unuseful messages
func Filter(m Message) (bool, Message) {
	var data PBResponse
	json.Unmarshal(m.Data, &data)
	if data.Type == "push" || data.Type == "mirror" {
		for _, el := range data.Push.Notifications {
			if len(data.Push.Title) > 0 {
				data.Push.Title += "\n" + el.Title
			} else {
				data.Push.Title = el.Title
			}

			if len(data.Push.Body) > 0 {
				data.Push.Body += "\n" + el.Body
			} else {
				data.Push.Body = el.Body
			}
		}
		return true, Message{
			Data:    []byte("Message Type: " + data.Type + "\nMessage Title: " + data.Push.Title + "\nMessage Body: " + data.Push.Body),
			Address: m.Address,
		}
	}
	return false, m
}
