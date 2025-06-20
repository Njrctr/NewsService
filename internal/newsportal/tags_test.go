package newsportal

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"log"
	"news-service/internal/db"
	"testing"
)

func TestService_Tags(t *testing.T) {
	cfgDb := db.TestDBCfg()
	ctx := context.Background()

	dbconn := pg.Connect(cfgDb)
	if err := dbconn.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	repository := db.NewNewsRepo(dbconn)
	s := NewManager(repository)

	got, err := s.Tags(ctx)
	if err != nil {
		t.Errorf("Manager.Tags() error = %v", err)
		return
	}

	for _, t := range got {
		fmt.Println(t)

	}
}
