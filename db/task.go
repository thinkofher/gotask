package db

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Task represents single taks.
type Task struct {
	ID   int       `json:"id"`
	Body string    `json:"body"`
	Tags []string  `json:"tags"`
	Date time.Time `json:"date"`
}

// Checker chceks for specific informations about single task.
type Checker func(Task) bool

// ToJSON returns the JSON encoding of task struct.
func (t Task) ToJSON() ([]byte, error) {
	return json.Marshal(t)
}

// ReadFromJSON updates fields of Task with
// data from given []byte slice.
func (t *Task) ReadFromJSON(JSONBytes []byte) error {
	return json.Unmarshal(JSONBytes, &t)
}

// SetCurrDate updates Task time to current one.
func (t *Task) SetCurrDate() {
	t.Date = time.Now()
}

// ParseTags creates a list of strings (tags) from tags in
// string separated with given character.
func (t *Task) ParseTags(tagStr string, sep string) {
	t.Tags = strings.Split(tagStr, sep)
}

func (t Task) String() string {
	tags := strings.Join(t.Tags, ", ")
	return fmt.Sprintf(
		"  Content: %s\nGlobal ID: %d\n     Tags: %s\n    Added: %s",
		t.Body, t.ID, tags, t.Date.Format(time.ANSIC))
}

// TaskFromJSON returns Task parsed from given
// bytes slice containg JSON.
func TaskFromJSON(JSONBytes []byte) (Task, error) {
	var t Task
	err := json.Unmarshal(JSONBytes, &t)
	if err != nil {
		return t, err
	}
	return t, nil
}

// TasksFromJSON returns Task slice parsed from given
// bytes slice containg json.
func TasksFromJSON(JSONBytes []byte) ([]Task, error) {
	var t []Task
	err := json.Unmarshal(JSONBytes, &t)
	if err != nil {
		return t, err
	}
	return t, nil
}

// TaskSelection applies given Checker funcs to slice of tasks
// and returns Task slice, which fullfils Checkers requirements.
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

// TagChecker returns Checker which checks if Task
// contains given tag.
func TagChecker(tag string) Checker {
	return func(t Task) bool { return stringInSlice(tag, t.Tags) }
}

// IdChecker returns Checker which checks if id of Task
// is equal to given id.
func IDChecker(id int) Checker {
	return func(t Task) bool { return t.ID == id }
}

// Checks if given []string slice contains
// given string.
func stringInSlice(s string, list []string) bool {
	for _, val := range list {
		if val == s {
			return true
		}
	}
	return false
}
