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

	repository := NewNewsRepo(db)

	got, err := repository.TagsByFilters(ctx, &TagSearch{IDs: []int{1, 2}}, PagerNoLimit)
	if err != nil {
		t.Errorf("GetTags() error = %v", err)
		return
	}

	for _, tag := range got {
		fmt.Println(tag)
	}
}
