package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
}

func CreateTask(task *Task) error {
	db.AutoMigrate(&task)
	return db.Create(task).Error
}

func GetTasks() ([]Task, error) {
	var tasks []Task
	err := db.Find(&tasks).Error
	return tasks, err
}

func DeleteTask(id uint) error {
	return db.Delete(&Task{}, id).Error
}
