package repository

import (
	"context"
	"fmt"
	"log"
	"news-service/internal/db"
	"news-service/internal/tools"
	"testing"
)

func TestGetCategories(t *testing.T) {
	cfgDb := tools.GetTestDBCfg()
	ctx := context.Background()

	if err := db.InitDB(ctx, cfgDb); err != nil {
		log.Fatal(err)
	}

	got, err := GetCategories(ctx)
	if err != nil {
		t.Errorf("GetCategories() error = %v", err)
		return
	}

	for _, cat := range got {
		fmt.Println(*cat)
	}
}
