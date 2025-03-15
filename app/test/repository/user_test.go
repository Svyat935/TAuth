package repository

import (
	"TAuth/app/configs/database"
	"TAuth/app/internal/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitUserTableSuccess(t *testing.T) {
	dbInfo := database.LoadDatabaseSettings(".test_env")
	db, _ := dbInfo.DBSessionCreate()

	rep := repository.UserRepository{DB: db}
	result, err := rep.InitUserTable()

	assert.NotEmpty(t, result)
	assert.Nil(t, err)
}
