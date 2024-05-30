package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type List struct {
	Items []Item
}

func (l *List) Add(task string) {
	id := len(l.Items) + 1
	item := Item{ID: id, Task: task}

	l.Items = append(l.Items, item)
}

func (l *List) Remove(id int) {
	for i, item := range l.Items {
		if item.ID == id {
			l.Items = append(l.Items[:i], l.Items[i+1:]...)
			return
		}
	}
}

func (l *List) Completed(id int) {
	for i, item := range l.Items {
		if item.ID == id {
			l.Items[i].Completed = true
			return
		}
	}
}

func (l *List) List() {
	for _, item := range l.Items {
		status := " "
		if item.Completed {
			status = "X"
		}
		fmt.Printf("[%s] %d: %s\n", status, item.ID, item.Task)
	}
}

func (l *List) Save(fileName string) error {
	data, err := json.Marshal(l.Items)

	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, data, 0644)
}

func (l *List) Load(fileName string) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(data, &l.Items)
}
