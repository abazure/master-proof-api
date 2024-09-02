package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpenConnection(t *testing.T) {

	db := OpenConnection()
	assert.NotNil(t, db)

}
