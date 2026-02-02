package database

import (
	"testing"

	"github.com/freedom-sketch/sub2go-core/infra/config"
	"github.com/freedom-sketch/sub2go-core/tg_bot/utils"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	err := config.Load("config.json")
	if err != nil {
		t.Fatal(err)
	}

	cfg := config.Get()

	db, err := Connect(&cfg.DataBase)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}
func setupTestTx(t *testing.T) *gorm.DB {
	db := setupTestDB(t)
	tx := db.Begin()
	t.Cleanup(func() {
		tx.Rollback()
	})
	return tx
}

func TestUUID(t *testing.T) {
	uuid := utils.IntToUUID(5207913851)
	t.Log(uuid)
}
