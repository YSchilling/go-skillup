package db

import (
	"01-todo-list/config"
	"encoding/csv"
	"errors"
	"os"
	"strconv"
)

type Task struct {
	Id    int
	Title string
	Done  bool
}

func AddTask(title string) error {
	// get id
	tasks, err := getAllTasksFromDB()
	if err != nil {
		return err
	}

	var id int
	if len(tasks) == 0 {
		id = 0
	} else {
		lastTask := tasks[len(tasks)-1]
		id = lastTask.Id + 1
	}

	// add task
	err = writeTaskToDB(Task{id, title, false})
	if err != nil {
		return err
	}

	return nil
}

func GetTasks() ([]Task, error) {
	tasks, err := getAllTasksFromDB()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func CompleteTask(id int) error {
	// get tasks
	tasks, err := getAllTasksFromDB()
	if err != nil {
		return err
	}

	// modify tasks
	foundTask := false
	for i := range len(tasks) {
		if tasks[i].Id == id {
			tasks[i].Done = true
			foundTask = true
			break
		}
	}
	if !foundTask {
		return errors.New("Couldn't find task with the specified id: " + strconv.Itoa(id))
	}

	// save tasks
	err = writeAllTasksToDB(tasks)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTask(id int) error {
	// get tasks
	tasks, err := getAllTasksFromDB()
	if err != nil {
		return err
	}

	// modify tasks
	foundTask := false
	for i := range len(tasks) {
		if tasks[i].Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			foundTask = true
			break
		}
	}
	if !foundTask {
		return errors.New("Couldn't find task with the specified id: " + strconv.Itoa(id))
	}

	// save tasks
	err = writeAllTasksToDB(tasks)
	if err != nil {
		return err
	}

	return nil
}

func writeTaskToDB(newTask Task) error {
	// open file
	file, err := os.OpenFile(config.DBPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// create writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// write task
	if err := writer.Write([]string{
		strconv.Itoa(newTask.Id),
		newTask.Title,
		strconv.FormatBool(newTask.Done),
	}); err != nil {
		return err
	}

	return nil
}

func writeAllTasksToDB(tasks []Task) error {
	// open file
	file, err := os.OpenFile(config.DBPath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// create writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// writer header
	headers := []string{"id", "title", "done"}
	if err := writer.Write(headers); err != nil {
		return err
	}

	// write tasks
	for i := range len(tasks) {
		if err := writer.Write([]string{
			strconv.Itoa(tasks[i].Id),
			tasks[i].Title,
			strconv.FormatBool(tasks[i].Done),
		}); err != nil {
			return err
		}
	}

	return nil
}

func getAllTasksFromDB() ([]Task, error) {
	// open file
	file, err := os.Open(config.DBPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// read lines
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// convert lines to tasks
	var tasks []Task
	for i := 1; i < len(records); i++ {

		id, err := strconv.Atoi(records[i][0])
		if err != nil {
			return nil, err
		}

		done, err := strconv.ParseBool(records[i][2])
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, Task{
			id,
			records[i][1],
			done,
		})
	}

	return tasks, nil
}
