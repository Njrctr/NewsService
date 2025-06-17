package newsportal

import (
	"context"
	"fmt"
	"log"
	"news-service/internal/db"
	"testing"
)

func TestGetNewsByID(t *testing.T) {
	cfgDb := db.TestDBCfg()
	ctx := context.Background()

	dbconn, err := db.New(ctx, cfgDb)
	if err != nil {
		log.Fatal(err)
	}
	repository := db.NewRepository(dbconn)
	services := New(repository)

	got, err := services.NewsByID(ctx, 1)
	if err != nil {
		t.Errorf("GetNewsByID() error = %v", err)
		return
	}

	fmt.Printf("news: %v, newsTags: %v", got, got.Tags)
}

func TestNewsService_GetNews(t *testing.T) {
	cfgDb := db.TestDBCfg()
	ctx := context.Background()

	dbconn, err := db.New(ctx, cfgDb)
	if err != nil {
		log.Fatal(err)
	}
	repository := db.NewRepository(dbconn)
	services := New(repository)

	filter := &NewsFilter{
		CategoryID: 1,
		TagID:      0,
	}

	got, err := services.NewsByFilters(ctx, filter, 0, 5)
	if err != nil {
		t.Errorf("NewsService.GetNews() error = %v", err)
		return
	}
	for _, news := range got {
		fmt.Printf("news: %d, tags: %v\n", news.ID, news.Tags)
	}
}
