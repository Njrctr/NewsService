package rpc

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/stretchr/testify/assert"
	"log"
	"news-service/internal/db"
	"news-service/internal/newsportal"
	"testing"
)

func TestNewsService_GetNewsCount(t *testing.T) {

	tests := []struct {
		name string
		cat  int
		tag  int
		want int
	}{
		{
			name: "2",
			cat:  1,
			tag:  1,
			want: 2,
		},
		{
			name: "3",
			cat:  0,
			tag:  0,
			want: 3,
		},
	}

	cfgDb := db.TestDBCfg()
	ctx := context.Background()

	dbconn := pg.Connect(cfgDb)
	if err := dbconn.Ping(ctx); err != nil {
		log.Fatal(err)
	}
	repository := db.NewRepository(dbconn)
	services := newsportal.NewManager(repository)

	rpc := NewNewsService(services)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := rpc.Count(ctx, NewsFilter{
				CategoryID: tt.cat,
				TagID:      tt.tag,
			})

			assert.Equal(t, tt.want, got)
		})
	}
}
