package db

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestGetCategories(t *testing.T) {
	cfgDb := TestDBCfg()
	ctx := context.Background()

	db, err := New(ctx, cfgDb)
	if err != nil {
		log.Fatal(err)
	}
	repository := NewRepository(db)

	got, err := repository.GetCategories(ctx)
	if err != nil {
		t.Errorf("GetCategories() error = %v", err)
		return
	}

	for _, cat := range got {
		fmt.Println(*cat)
	}
}
