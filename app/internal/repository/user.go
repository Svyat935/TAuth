package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type UserModel struct {
	ID       string `DB:"id"`
	Username string `DB:"username"`
	Password string `DB:"password"`
}

const creatingTable = `
CREATE TABLE IF NOT EXISTS Users 
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
)
`

type UserRepository struct {
	DB *sqlx.DB
}

func (rep *UserRepository) InitUserTable() (result sql.Result, err error) {
	defer rep.DB.Close()

	result, err = rep.DB.Exec(creatingTable)
	if err != nil {
		return nil, err
	}

	return result, nil
}
