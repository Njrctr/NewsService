package repository

import (
	"context"
	"fmt"
	"log"
	"news-service/internal/db"
	"news-service/internal/structs"
	"news-service/internal/tools"
	"testing"
)

func TestNewsRepository_GetNewsByID(t *testing.T) {
	cfgDb := tools.GetTestDBCfg()
	ctx := context.Background()

	if err := db.InitDB(ctx, cfgDb); err != nil {
		log.Fatal(err)
	}
	got, err := GetNewsByID(ctx, 1)
	if err != nil {

		t.Errorf("NewsRepository.GetNewsByID() error = %v", err)
		return
	}

	fmt.Println(got)

}

func TestGetNews(t *testing.T) {
	cfgDb := tools.GetTestDBCfg()
	ctx := context.Background()

	if err := db.InitDB(ctx, cfgDb); err != nil {
		log.Fatal(err)
	}

	filter := &structs.NewsFilter{
		CategoryID: 1,
		TagID:      0,
	}

	got, err := GetNews(ctx, filter, 0, 2)
	if err != nil {
		t.Errorf("GetNews() error = %v", err)
		return
	}

	for _, news := range got {
		fmt.Println(news)

	}
}

func TestGetNewsCount(t *testing.T) {
	cfgDb := tools.GetTestDBCfg()
	ctx := context.Background()

	if err := db.InitDB(ctx, cfgDb); err != nil {
		log.Fatal(err)
	}

	filter := &structs.NewsFilter{
		CategoryID: 1,
		TagID:      0,
	}

	got, err := GetNewsCount(ctx, filter)
	if err != nil {
		t.Errorf("GetNewsCount() error = %v", err)
		return
	}

	fmt.Println(got)
}
