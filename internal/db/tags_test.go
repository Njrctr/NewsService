package db

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestGetTags(t *testing.T) {
	cfgDb := TestDBCfg()
	ctx := context.Background()

	db, err := New(ctx, cfgDb)
	if err != nil {
		log.Fatal(err)
	}
	repository := NewRepository(db)

	got, err := repository.Tags(ctx, nil)
	if err != nil {
		t.Errorf("GetTags() error = %v", err)
		return
	}

	for _, tag := range got {
		fmt.Println(tag)
	}
}
