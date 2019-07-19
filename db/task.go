package db

import (
	"encoding/json"
	"strings"
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

// Creates a list of strings (tags) from tags in
// string separated with given character
func (t *Task) ParseTags(tagStr string, sep string) {
	t.Tags = strings.Split(tagStr, sep)
}

func TaskFromJson(jsonBytes []byte, t *Task) error {
	return json.Unmarshal(jsonBytes, &t)
}

func TasksFromJson(jsonBytes []byte, t *[]Task) error {
	return json.Unmarshal(jsonBytes, &t)
}
