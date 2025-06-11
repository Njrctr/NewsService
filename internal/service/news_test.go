package service

import (
	"context"
	"fmt"
	"log"
	"news-service/internal/db"
	"news-service/internal/structs"
	"news-service/internal/tools"
	"testing"
)

func TestNewsService_GetNews(t *testing.T) {
	dbCfg := tools.GetTestDBCfg()
	ctx := context.Background()

	if err := db.InitDB(ctx, dbCfg); err != nil {
		log.Fatal(err)
	}

	filter := &structs.NewsFilter{
		CategoryID: 1,
		TagID:      0,
	}

	got, err := GetNews(ctx, filter, 0, 5)
	if err != nil {
		t.Errorf("NewsService.GetNews() error = %v", err)
		return
	}
	for _, news := range got {
		fmt.Printf("news: %d, newsTags:%v\n", news.ID, news.Tags)
	}
}
