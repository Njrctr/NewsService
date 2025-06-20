package rpc

import (
	"context"
	"fmt"
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

func TestNewsService_Get(t *testing.T) {

	tests := []struct {
		name      string
		filter    NewsFilter
		wantCount int
		page      PageRequest
		wantErr   bool
		wantedErr error
	}{
		{
			name: "OK2",
			filter: NewsFilter{
				CategoryID: 1,
				TagID:      1,
			},
			wantCount: 2,
			page:      PageRequest{},
			wantErr:   false,
		},
		{
			name:      "OK3",
			filter:    NewsFilter{},
			wantCount: 3,
			page:      PageRequest{},
		},
		{
			name: "Not Found",
			filter: NewsFilter{
				CategoryID: 10,
				TagID:      12,
			},
			page:    PageRequest{},
			wantErr: false,
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

			got, err := rpc.Get(ctx, tt.filter, tt.page)

			if tt.wantErr {
				assert.ErrorIs(t, err, tt.wantedErr)
				return
			}
			assert.Equalf(t, tt.wantCount, len(got), "Get(%v, %v)", tt.filter, tt.page)
		})
	}
}

func TestNewsService_GetByID(t *testing.T) {

	tests := []struct {
		name      string
		id        int
		wantErr   bool
		wantedErr error
	}{
		{
			name: "OK",
			id:   1,
		},
		{
			name:      "Not Found",
			id:        10,
			wantErr:   true,
			wantedErr: errNotFound,
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

			got, err := rpc.GetByID(ctx, tt.id)

			if !tt.wantErr {
				fmt.Println(got)
				return
			}
			assert.ErrorIs(t, err, errNotFound)
		})
	}
}

func TestTagService_Get(t *testing.T) {
	type fields struct {
		manager *newsportal.Manager
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name      string
		wantCount int
		wantErr   bool
		wantedErr error
	}{
		{
			name:      "OK",
			wantCount: 2,
		},
	}

	// TODO Test MAIN
	cfgDb := db.TestDBCfg()
	ctx := context.Background()

	dbconn := pg.Connect(cfgDb)
	if err := dbconn.Ping(ctx); err != nil {
		log.Fatal(err)
	}
	repository := db.NewRepository(dbconn)
	services := newsportal.NewManager(repository)

	rpc := NewCategoryService(services)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := rpc.Get(ctx)
			if !tt.wantErr {
				assert.Equal(t, tt.wantCount, len(got))
				return
			}
			assert.ErrorIs(t, err, errNotFound)
		})
	}
}
