package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitializeFirebase(t *testing.T) {

	app := InitializeFirebase()
	assert.NotNil(t, app)
}
