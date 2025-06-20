package newsportal

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"log"
	"news-service/internal/db"
	"testing"
)

func TestGetNewsByID(t *testing.T) {
	cfgDb := db.TestDBCfg()
	ctx := context.Background()

	dbconn := pg.Connect(cfgDb)
	if err := dbconn.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	repository := db.NewNewsRepo(dbconn)
	services := NewManager(repository)

	got, err := services.NewsByID(ctx, 1)
	if err != nil {
		t.Errorf("GetNewsByID() error = %v", err)
		return
	}

	fmt.Printf("news: %v\n", got)
	if len(got.Tags) != 0 {
		fmt.Printf("newsTags: %v\n", got.Tags)
	}

}

func TestNewsService_GetNews(t *testing.T) {
	cfgDb := db.TestDBCfg()
	ctx := context.Background()

	dbconn := pg.Connect(cfgDb)
	if err := dbconn.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	repository := db.NewNewsRepo(dbconn)
	services := NewManager(repository)

	filter := &NewsFilter{
		CategoryID: 1,
		TagID:      1,
	}

	got, err := services.NewsByFilters(ctx, filter, 0, 5)
	if err != nil {
		t.Errorf("NewsService.GetNews() error = %v", err)
		return
	}
	for _, news := range got {
		fmt.Printf("news: %v\n", news)
		if len(news.Tags) != 0 {
			fmt.Printf("newsTags: %v\n", news.Tags)
		}
	}
}
