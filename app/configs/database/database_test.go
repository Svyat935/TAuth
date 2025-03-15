package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadSettings(t *testing.T) {
	dbInfo := *LoadDatabaseSettings(".test_env")
	expected := DBInfo{
		"localhost",
		"5432",
		"postgres",
		"postgres",
		"",
		"postgres",
	}

	assert.Equal(t, dbInfo, expected)
}

func TestCreateDBSessionSuccess(t *testing.T) {
	dbInfo := LoadDatabaseSettings(".test_env")
	session := dbInfo.DBSessionPing()

	assert.NotEmpty(t, session)
}
