package dbrepo

import (
	"ama-back/internal/models"
	"context"
	"database/sql"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

func (m *PostgresDBRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT
			id, email, first_name, last_name, password, created_at, updated_at
		FROM
			users
		WHERE
			email = $1
	`

	var user models.User

	row := m.DB.QueryRowContext(ctx, query, email)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *PostgresDBRepo) GetUserByID(id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT
			id, email, first_name, last_name, password, created_at, updated_at
		FROM
			users
		WHERE
			id = $1
	`

	var user models.User

	row := m.DB.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *PostgresDBRepo) AllQuestions() ([]models.Question, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT
			id, author_email, question, answer, created_at, updated_at
		FROM
			questions
		ORDER BY
			created_at DESC
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []models.Question

	for rows.Next() {
		var question models.Question

		err = rows.Scan(
			&question.ID,
			&question.AuthorEmail,
			&question.Question,
			&question.Answer,
			&question.CreatedAt,
			&question.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		questions = append(questions, question)
	}

	return questions, nil
}

func (m *PostgresDBRepo) GetQuestion(id int) (*models.Question, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT
			id, author_email, question, answer, created_at, updated_at
		FROM
			questions
		WHERE
			id = $1
	`

	var question models.Question

	row := m.DB.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&question.ID,
		&question.AuthorEmail,
		&question.Question,
		&question.Answer,
		&question.CreatedAt,
		&question.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &question, nil
}

func (m *PostgresDBRepo) InsertQuestion(question models.Question) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `
		INSERT INTO questions
			(author_email, question, answer, created_at, updated_at)
		VALUES
			($1, $2, $3, $4, $5)
		RETURNING
			id
	`

	var newID int

	err := m.DB.QueryRowContext(ctx, stmt,
		question.AuthorEmail,
		question.Question,
		question.Answer,
		question.CreatedAt,
		question.UpdatedAt,
	).Scan(&newID)

	if err != nil {
		return -1, err
	}

	return newID, nil
}

func (m *PostgresDBRepo) UpdateQuestion(question models.Question) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `
		UPDATE
			questions
		SET
			author_email = $1,
			question = $2,
			answer = $3,
			updated_at = $4
		WHERE
			id = $5
	`

	_, err := m.DB.ExecContext(ctx, stmt,
		question.AuthorEmail,
		question.Question,
		question.Answer,
		question.UpdatedAt,
		question.ID)

	if err != nil {
		return err
	}

	return nil
}

func (m *PostgresDBRepo) DeleteQuestion(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `
		DELETE FROM 
			questions 
		WHERE 
			id = $1
		`

	_, err := m.DB.ExecContext(ctx, stmt, id)

	if err != nil {
		return err
	}

	return nil
}
