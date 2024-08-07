package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

type Task struct {
	Task string `json:"task"`
	Done bool   `json:"done"`
}

func (a *App) PostTask(task Task) {
	url := "http://localhost:4040/api/v1/todo"

	data, err := json.Marshal(task)
	if err != nil {
		log.Println("failed to marshal data: ", err)
		return
	}

	log.Println("Request Body:", string(data))

	client := http.Client{}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		log.Println("failed to create request: ", err)
		return
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		log.Println("failed to send request: ", err)
		return
	}
	defer response.Body.Close()

	log.Println("Response Status:", response.Status)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("failed to read response body: ", err)
		return
	}
	log.Println("Response Body:", string(body))
}

type GetTasks struct {
	ID        int        `json:"id"`
	Task      string     `json:"task"`
	Done      bool       `json:"done"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (a *App) GetAllTasks() ([]GetTasks, error) {
	url := "http://localhost:4040/api/v1/todo"
	client := http.Client{}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println("failed to create request:", err)
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		log.Println("failed to get response:", err)
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Printf("unexpected response status: %d %s", response.StatusCode, response.Status)
		return nil, fmt.Errorf("unexpected response status: %d %s", response.StatusCode, response.Status)
	}

	var tasks []GetTasks
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&tasks)
	if err != nil {
		log.Println("failed to unmarshal JSON:", err)
		return nil, err
	}

	log.Println("Decoded Tasks:", tasks)
	return tasks, nil
}

func (a *App) UpdateTask(id int, updatedTask Task) {
	url := fmt.Sprintf("http://localhost:4040/api/v1/todo/%d", id)
	client := http.Client{}

	data, err := json.Marshal(updatedTask)
	if err != nil {
		log.Println("failed to marshal data:", err)
		return
	}

	request, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(data))
	if err != nil {
		log.Println("failed to create request:", err)
		return
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		log.Println("failed to send request:", err)
		return
	}
	defer response.Body.Close()

	log.Println("Response Status:", response.Status)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("failed to read response body:", err)
		return
	}
	log.Println("Response Body:", string(body))
}

func (a *App) DeleteTask(id int) {
	url := fmt.Sprintf("http://localhost:4040/api/v1/todo/%d", id)
	client := http.Client{}

	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		log.Println("failed to create request:", err)
		return
	}

	response, err := client.Do(request)
	if err != nil {
		log.Println("failed to send request:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Printf("unexpected response status: %d %s", response.StatusCode, response.Status)
		return
	}

	log.Println("Task successfully deleted")
}
