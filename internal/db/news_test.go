package db

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestNewsRepository_GetNewsByID(t *testing.T) {
	cfgDb := TestDBCfg()
	ctx := context.Background()

	db, err := New(ctx, cfgDb)
	if err != nil {
		log.Fatal(err)
	}
	repository := NewRepository(db)
	got, err := repository.NewsByID(ctx, 4)
	if err != nil {

		t.Errorf("NewsRepository.GetNewsByID() error = %v", err)
		return
	}

	fmt.Println(got)

}

func TestGetNews(t *testing.T) {
	cfgDb := TestDBCfg()
	ctx := context.Background()

	db, err := New(ctx, cfgDb)
	if err != nil {
		log.Fatal(err)
	}
	repository := NewRepository(db)

	filter := &NewsFilter{
		CategoryID: 1,
		TagID:      1,
	}

	got, err := repository.NewsByFilters(ctx, filter, 0, 2)
	if err != nil {
		t.Errorf("GetNews() error = %v", err)
		return
	}

	for _, news := range got {
		fmt.Println(news)

	}
}

func TestGetNewsCount(t *testing.T) {
	cfgDb := TestDBCfg()
	ctx := context.Background()

	db, err := New(ctx, cfgDb)
	if err != nil {
		log.Fatal(err)
	}
	repository := NewRepository(db)

	filter := &NewsFilter{
		CategoryID: 1,
		TagID:      0,
	}

	got, err := repository.NewsCount(ctx, filter)
	if err != nil {
		t.Errorf("GetNewsCount() error = %v", err)
		return
	}

	fmt.Println(got)
}
