package rest

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http/httptest"
	"news-service/internal/db"
	"news-service/internal/newsportal"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getOneNews(t *testing.T) {
	testTable := []struct {
		name         string
		newsId       any
		expectedCode int
		expectErr    bool
		errMsg       string
	}{
		{
			name:         "Ok",
			newsId:       1,
			expectedCode: 200,
			expectErr:    false,
		},
		{
			name:         "not found",
			newsId:       5,
			expectedCode: 404,
			expectErr:    true,
			errMsg:       "news not found",
		},
		{
			name:         "invalid id param",
			newsId:       "qwe",
			expectedCode: 400,
			expectErr:    true,
			errMsg:       "invalid id param, should be int",
		},
	}

	cfgDb := db.TestDBCfg()
	ctx := context.Background()

	dbconn, err := db.New(ctx, cfgDb)
	if err != nil {
		log.Fatal(err)
	}
	repository := db.NewRepository(dbconn)
	services := newsportal.New(repository)
	handlers := New(services)

	router := echo.New()
	router.GET("/news/:id", handlers.getOneNews)

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			rec := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/news/%v", testCase.newsId), nil)

			router.ServeHTTP(rec, req)

			assert.Equal(t, testCase.expectedCode, rec.Code)
			if testCase.expectErr {
				assert.Equal(t, testCase.errMsg, rec.Body.String())
			}
		})
	}
}

func Test_getNews(t *testing.T) {
	testTable := []struct {
		name         string
		queryParams  string
		expectedCode int
		expectErr    bool
		errMsg       string
	}{
		{
			name:         "Ok",
			queryParams:  `tag=1&cat=1`,
			expectedCode: 200,
			expectErr:    false,
		},
		{
			name:         "not found",
			queryParams:  `tag=4&cat=6`,
			expectedCode: 404,
			expectErr:    true,
			errMsg:       "news not found",
		},
		{
			name:         "invalid query param",
			queryParams:  "tag=qwe&cat=1",
			expectedCode: 400,
			expectErr:    true,
			errMsg:       "invalid query param(s)",
		},
	}

	cfgDb := db.TestDBCfg()
	ctx := context.Background()

	dbconn, err := db.New(ctx, cfgDb)
	if err != nil {
		log.Fatal(err)
	}
	repository := db.NewRepository(dbconn)
	services := newsportal.New(repository)
	handlers := New(services)

	router := echo.New()
	router.GET("/news/", handlers.getNews)

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			rec := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/news/?%v", testCase.queryParams), nil)

			router.ServeHTTP(rec, req)

			assert.Equal(t, testCase.expectedCode, rec.Code)
			if testCase.expectErr {
				assert.Equal(t, testCase.errMsg, rec.Body.String())
			}
		})
	}

}
