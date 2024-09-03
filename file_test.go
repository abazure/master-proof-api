package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"master-proof-api/database"
	"master-proof-api/dto"
	"testing"
)

func TestRepositoryFile(t *testing.T) {
	db := database.OpenConnection()
	var learningMaterial []dto.LearningMaterialResponse
	err := db.Table("learning_materials as lm").Select("lm.id ,lm.title, lm.description, f.url").Joins("join files as f on f.id = lm.file_id").Take(&learningMaterial).Error
	assert.Nil(t, err)
	fmt.Println(learningMaterial)
}
