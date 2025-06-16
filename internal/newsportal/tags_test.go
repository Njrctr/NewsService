package newsportal

import (
	"context"
	"fmt"
	"log"
	"news-service/internal/db"
	"testing"
)

func TestService_Tags(t *testing.T) {
	cfgDb := db.TestDBCfg()
	ctx := context.Background()

	dbconn, err := db.New(ctx, cfgDb)
	if err != nil {
		log.Fatal(err)
	}
	repository := db.NewRepository(dbconn)
	s := New(repository)

	got, err := s.Tags(ctx)
	if err != nil {
		t.Errorf("Service.Tags() error = %v", err)
		return
	}

	for _, t := range got {
		fmt.Println(t)

	}
}
