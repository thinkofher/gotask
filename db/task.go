package db

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Represents single taks.
type Task struct {
	Id   int       `json:"id"`
	Body string    `json:"body"`
	Tags []string  `json:"tags"`
	Date time.Time `json:"date"`
}

// Returns the JSON encoding of task struct.
func (t Task) ToJson() ([]byte, error) {
	return json.Marshal(t)
}

// Updates fields of Task with data from given
// []byte slice.
func (t *Task) ReadFromJson(jsonBytes []byte) error {
	return json.Unmarshal(jsonBytes, &t)
}

func (t *Task) SetCurrDate() {
	t.Date = time.Now()
}

// Creates a list of strings (tags) from tags in
// string separated with given character.
func (t *Task) ParseTags(tagStr string, sep string) {
	t.Tags = strings.Split(tagStr, sep)
}

func (t Task) String() string {
	tags := strings.Join(t.Tags, ", ")
	return fmt.Sprintf(
		"Task: %s\nId: %d\nTags: %s\nAdded: %s",
		t.Body, t.Id, tags, t.Date.Format(time.ANSIC))
}

func TaskFromJson(jsonBytes []byte) (Task, error) {
	var t Task
	err := json.Unmarshal(jsonBytes, &t)
	if err != nil {
		return t, err
	}
	return t, nil
}

func TasksFromJson(jsonBytes []byte) ([]Task, error) {
	var t []Task
	err := json.Unmarshal(jsonBytes, &t)
	if err != nil {
		return t, err
	}
	return t, nil
}
