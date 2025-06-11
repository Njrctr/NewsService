package repository

import (
	"context"
	"fmt"
	"log"
	"news-service/internal/db"
	"news-service/internal/tools"
	"testing"
)

func TestGetTags(t *testing.T) {
	cfgDb := tools.GetTestDBCfg()
	ctx := context.Background()

	if err := db.InitDB(ctx, cfgDb); err != nil {
		log.Fatal(err)
	}

	got, err := GetTags(ctx)
	if err != nil {
		t.Errorf("GetTags() error = %v", err)
		return
	}

	for _, tag := range got {
		fmt.Println(*tag)
	}
}

func TestGetTagsByIds(t *testing.T) {
	cfgDb := tools.GetTestDBCfg()
	ctx := context.Background()

	if err := db.InitDB(ctx, cfgDb); err != nil {
		log.Fatal(err)
	}

	ids := []int{1, 2}

	got, err := GetTagsByIds(ctx, ids)
	if err != nil {
		t.Errorf("GetTagsByIds() error = %v", err)
		return
	}

	for key, tag := range got {
		fmt.Printf("key:%d, tag:%v\n", key, *tag)
	}
}
