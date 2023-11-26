package repository

import (
	"ama-back/internal/models"
	"database/sql"
)

type DatabaseRepo interface {
	Connection() *sql.DB

	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)

	AllQuestions() ([]models.Question, error)
	GetQuestion(id int) (*models.Question, error)
	InsertQuestion(question models.Question) (int, error)
	UpdateQuestion(question models.Question) error
	DeleteQuestion(id int) error
}
