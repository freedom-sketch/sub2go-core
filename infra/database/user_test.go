package database

import (
	"math/rand"
	"testing"

	"github.com/freedom-sketch/sub2go-core/infra/database/models"
	"github.com/google/uuid"
)

func TestCreateUser(t *testing.T) {
	db := setupTestTx(t)

	user := &models.User{
		UUID:       uuid.New().String(),
		TelegramID: &[]int64{rand.Int63()}[0],
		IsActive:   true,
	}

	err := CreateUser(db, user)
	if err != nil {
		t.Errorf("CreateUser failed: %v", err)
	}

	var found models.User
	err = db.Where("uuid = ?", user.UUID).First(&found).Error
	if err != nil {
		t.Errorf("User not found: %v", err)
	}
	if found.UUID != user.UUID {
		t.Errorf("UUID mismatch: expected %s, got %s", user.UUID, found.UUID)
	}
}
