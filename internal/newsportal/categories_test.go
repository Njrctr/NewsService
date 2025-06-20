package newsportal

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"log"
	"news-service/internal/db"
	"testing"
)

func TestService_GetCategories(t *testing.T) {
	cfgDb := db.TestDBCfg()
	ctx := context.Background()

	dbconn := pg.Connect(cfgDb)
	if err := dbconn.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	repository := db.NewNewsRepo(dbconn)
	s := NewManager(repository)

	got, err := s.Categories(ctx)
	if err != nil {
		t.Errorf("Manager.GetCategories() error = %v", err)
		return
	}

	for _, c := range got {
		fmt.Println(c)
	}
}
