package database

import (
	"math/rand"
	"testing"
	"time"

	"github.com/freedom-sketch/sub2go-core/infra/database/models"
	"github.com/google/uuid"
)

func TestCreateSubscription(t *testing.T) {
	db := setupTestTx(t)

	user := &models.User{
		UUID:       uuid.New().String(),
		TelegramID: &[]int64{rand.Int63()}[0],
		IsActive:   true,
	}
	err := CreateUser(db, user)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	subscription := &models.Subscription{
		UserUUID:     user.UUID,
		Email:        uuid.New().String() + "@example.com",
		TotalTraffic: rand.Int63(),
		UsedTraffic:  0,
		StartDate:    time.Now(),
		EndDate:      time.Now().Add(30 * 24 * time.Hour),
		IsActive:     true,
	}

	err = CreateSubscription(db, subscription)
	if err != nil {
		t.Errorf("CreateSubscription failed: %v", err)
	}

	var found models.Subscription
	err = db.Where("user_uuid = ?", user.UUID).First(&found).Error
	if err != nil {
		t.Errorf("Subscription not found: %v", err)
	}
	if found.Email != subscription.Email {
		t.Errorf("Email mismatch: expected %s, got %s", subscription.Email, found.Email)
	}
}
