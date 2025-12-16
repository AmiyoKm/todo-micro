package main

import (
	"bytes"
	"encoding/json"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Todo struct {
	ID          string `json:"id"`
	UserId      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type Message struct {
	Topic string `json:"topic"`
	Todo  Todo   `json:"todo"`
}

func (a *app) handleMessage(msg amqp.Delivery) error {

	var message Message
	time.Sleep(time.Second * 2)
	err := json.NewDecoder(bytes.NewReader(msg.Body)).Decode(&message)
	if err != nil {
		return err
	}

	if message.Topic == "todo.update" {
		err = a.updateTodo(&message.Todo)
		if err != nil {
			return err
		}
	}

	if message.Topic == "todo.delete" {
		err = a.deleteTodo(message.Todo.ID, message.Todo.UserId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *app) updateTodo(todo *Todo) error {
	query := `UPDATE todos SET title = $1, description = $2, done = $3 WHERE id = $4 AND user_id = $5`
	_, err := a.db.Exec(query, todo.Title, todo.Description, todo.Done, todo.ID, todo.UserId)
	if err != nil {
		return err
	}

	return nil
}

func (a *app) deleteTodo(id string, userId string) error {

	query := `DELETE FROM todos WHERE id = $1 AND user_id = $2`

	_, err := a.db.Exec(query, id, userId)
	if err != nil {
		return err
	}

	return nil
}
