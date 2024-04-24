package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

// Um construtor de nova categoria, para toda vez que formos criar uma categoria passaremos uma conexão com o banco de dados.
func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

// Create- Cria/salva uma nova categoria no banco e retorna ela com o seu ID.
func (c *Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)", id, name, description)
	if err != nil {
		return Category{}, err
	}

	return Category{ID: id, Name: name, Description: description}, nil
}
