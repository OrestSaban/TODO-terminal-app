package main

import (
    "time"
    "errors" 
    "fmt"    
    "github.com/olekukonko/tablewriter"
    "os"
    "encoding/json"
    "io/ioutil"
)

type Todo struct {
	Title       string     `json:"title"`
	CreatedAt   time.Time  `json:"created_at"`
	Completed   bool       `json:"completed"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

type Todos []Todo

func (todos *Todos) add(task string) {
	todo := Todo{
		Title: task,
		CreatedAt: time.Now(),
		Completed: false,
		CompletedAt: nil,
	}
	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if(index < 0 || index >= len(*todos)) {
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) delete(index int){
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return
	}

	*todos = append(t[:index], t[index+1:]...)
}

func (todos *Todos) toggle(index int) {
	t := *todos

	if err := t.validateIndex(index); err != nil{
		return
	}

	isCompleted := t[index].Completed

	if !isCompleted{
		now := time.Now()
		t[index].Completed = !isCompleted
		t[index].CompletedAt = &now
	} else {
		t[index].Completed = !isCompleted
		t[index].CompletedAt = nil
	}
}

func (todos *Todos) update(index int, task string) {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return
	}
	t[index].Title = task
	t[index].Completed = false
	t[index].CompletedAt = nil
}

func (todos *Todos) print() {
    table := tablewriter.NewWriter(os.Stdout)
    table.SetHeader([]string{"Index", "Task", "Status", "Created At", "Completed At"})

    for i, todo := range *todos {
        status := "❌"
        if todo.Completed {
            status = "✅"
        }

        completedAt := "N/A"
        if todo.CompletedAt != nil {
            completedAt = todo.CompletedAt.Format(time.RFC1123)
        }

        table.Append([]string{
            fmt.Sprintf("%d", i),
            todo.Title,
            status,
            todo.CreatedAt.Format(time.RFC1123),
            completedAt,
        })
    }

    table.Render() // Print the table to stdout
}

func (todos *Todos) saveToFile(filename string) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

func (todos *Todos) loadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // No file to load, start with an empty list
		}
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, todos)
}

func (todos *Todos) clear() {
    *todos = Todos{} // Reset the slice to an empty slice
}





