package rest

import (
	"context"
	"fmt"
	"log"
	"net/http/httptest"
	"news-service/internal/db"
	"news-service/internal/newsportal"
	"testing"

	"github.com/gin-gonic/gin"
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
		}, {
			name:         "тшдд",
			newsId:       3,
			expectedCode: 404,
			expectErr:    false,
		},
		{
			name:         "invalid id param",
			newsId:       "qwe",
			expectedCode: 400,
			expectErr:    true,
			errMsg:       `{"message":"invalid id param, should be int"}`,
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

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
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
