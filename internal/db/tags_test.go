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

	got, err := repository.Tags(ctx)
	if err != nil {
		t.Errorf("GetTags() error = %v", err)
		return
	}

	for _, tag := range got {
		fmt.Println(*tag)
	}
}

func TestGetTagsByIds(t *testing.T) {
	cfgDb := TestDBCfg()
	ctx := context.Background()

	db, err := New(ctx, cfgDb)
	if err != nil {
		log.Fatal(err)
	}
	repository := NewRepository(db)

	ids := []int{1, 2}

	got, err := repository.TagsByIds(ctx, ids)
	if err != nil {
		t.Errorf("GetTagsByIds() error = %v", err)
		return
	}

	for key, tag := range got {
		fmt.Printf("key:%d, tag:%v\n", key, *tag)
	}
}
