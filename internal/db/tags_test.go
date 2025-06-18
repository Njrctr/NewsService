package db

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"log"
	"testing"
)

func TestGetTags(t *testing.T) {
	cfgDb := TestDBCfg()
	ctx := context.Background()

	db := pg.Connect(cfgDb)
	if err := db.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	repository := NewRepository(db)

	got, err := repository.Tags(ctx, []int{1, 2})
	if err != nil {
		t.Errorf("GetTags() error = %v", err)
		return
	}

	for _, tag := range got {
		fmt.Println(tag)
	}
}
