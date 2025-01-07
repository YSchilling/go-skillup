package db

import (
	"01-todo-list/config"
	"encoding/csv"
	"os"
	"strconv"
)

type Task struct {
	id    int
	title string
	done  bool
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
		id = lastTask.id + 1
	}

	// add task
	err = writeTaskToDB(Task{id, title, false})
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
		strconv.Itoa(newTask.id),
		newTask.title,
		strconv.FormatBool(newTask.done),
	}); err != nil {
		return err
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
