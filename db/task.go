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

// Chceks for specific informations about single task
type Checker func(Task) bool

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
		"  Content: %s\nGlobal Id: %d\n     Tags: %s\n    Added: %s",
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

// Apply given Checker funcs to slice of task and returns
// Task slice, which fullfils Checkers requirements
func TaskSelection(tasks []Task, funcs ...Checker) []Task {
	var ans []Task
	for _, task := range tasks {
		for _, f := range funcs {
			if f(task) {
				ans = append(ans, task)
			}
		}
	}
	return ans
}

// Return Checker which checks if Task
// contains given tag
func TagChecker(tag string) Checker {
	return func(t Task) bool { return stringInSlice(tag, t.Tags) }
}

// Checks if given []string slice contains
// given string
func stringInSlice(s string, list []string) bool {
	for _, val := range list {
		if val == s {
			return true
		}
	}
	return false
}
