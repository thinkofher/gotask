package db

import (
	"encoding/json"
	"time"
)

// Represents single taks
type Task struct {
	Id   int       `json:"id"`
	Body string    `json:"body"`
	Tags []string  `json:"tags"`
	Date time.Time `json:"date"`
}

// Returns the JSON encoding of task struct
func (t Task) ToJson() ([]byte, error) {
	return json.Marshal(t)
}

func (t *Task) SetCurrDate() {
	t.Date = time.Now()
}

func TaskFromJson(jsonBytes []byte, t *Task) error {
	return json.Unmarshal(jsonBytes, &t)
}

func TasksFromJson(jsonBytes []byte, t *[]Task) error {
	return json.Unmarshal(jsonBytes, &t)
}
