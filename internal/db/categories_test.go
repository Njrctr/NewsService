package db

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"log"
	"testing"
)

func TestGetCategories(t *testing.T) {
	cfgDb := TestDBCfg()
	ctx := context.Background()

	db := pg.Connect(cfgDb)
	if err := db.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	repository := NewNewsRepo(db)

	got, err := repository.CategoriesByFilters(ctx, nil, PagerNoLimit)
	if err != nil {
		t.Errorf("GetCategories() error = %v", err)
		return
	}

	for _, cat := range got {
		fmt.Println(cat)
	}
}
