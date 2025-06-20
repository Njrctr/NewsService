package rpc

import (
	"context"
	"github.com/go-pg/pg/v10"
	middleware "github.com/vmkteam/zenrpc-middleware"
	"github.com/vmkteam/zenrpc/v2"
	"log"
	"net/http"
	"news-service/internal/newsportal"
	"os"
)

//go:generate zenrpc

var (
	errInternalError = zenrpc.NewStringError(http.StatusInternalServerError, `internal error`)
	errNotFound      = zenrpc.NewStringError(http.StatusNotFound, `not found`)
)

type NewsService struct {
	manager *newsportal.Manager
} //zenrpc

func NewNewsService(manager *newsportal.Manager) *NewsService {
	return &NewsService{
		manager: manager,
	}
}

type TagService struct {
	manager *newsportal.Manager
} //zenrpc

func NewTagService(manager *newsportal.Manager) *TagService {
	return &TagService{
		manager: manager,
	}
}

type CategoryService struct {
	manager *newsportal.Manager
} //zenrpc

func NewCategoryService(manager *newsportal.Manager) *CategoryService {
	return &CategoryService{
		manager: manager,
	}
}

func Init(service *newsportal.Manager, dbconn *pg.DB) zenrpc.Server {
	rpc := zenrpc.NewServer(zenrpc.Options{
		ExposeSMD: true,
		AllowCORS: true,
	})
	rpc.Register("news", NewNewsService(service))
	rpc.Register("categories", NewCategoryService(service))
	rpc.Register("tags", NewTagService(service))

	elog := log.New(os.Stderr, "E", log.LstdFlags|log.Lshortfile)
	dlog := log.New(os.Stderr, "D", log.LstdFlags|log.Lshortfile)
	allowDebug := func(param string) middleware.AllowDebugFunc {
		return func(req *http.Request) bool {
			return req.FormValue(param) == "true"
		}
	}

	rpc.Use(

		middleware.WithDevel(true),
		middleware.WithAPILogger(dlog.Printf, middleware.DefaultServerName),
		middleware.WithTiming(true, allowDebug("d")),
		middleware.WithMetrics(middleware.DefaultServerName),
		middleware.WithSQLLogger(dbconn, true, allowDebug("d"), allowDebug("s")),
		middleware.WithErrorLogger(elog.Printf, middleware.DefaultServerName),
	)

	return rpc
}

type PageRequest struct {
	PageSize   int `json:"page_size"`
	PageNumber int `json:"page_num"`
}

// Get возвращает список новостей по указанным фильтрам
//
//zenrpc:500 		server error
//zenrpc:404 		news not found
func (ns NewsService) Get(ctx context.Context, filter NewsFilter, page PageRequest) ([]News, error) {
	news, err := ns.manager.NewsByFilters(ctx, &newsportal.NewsFilter{
		CategoryID: filter.CategoryID,
		TagID:      filter.TagID,
	}, page.PageNumber, page.PageSize)
	if err != nil {
		return nil, errInternalError
	}

	return newNewsSlice(news), nil
}

// GetByID возвращает новость по ID
//
//zenrpc:500 		server error
//zenrpc:404 		news not found
func (ns NewsService) GetByID(ctx context.Context, id int) (News, error) {
	news, err := ns.manager.NewsByID(ctx, id)
	if err != nil {
		return News{}, errInternalError
	} else if news == nil {
		return News{}, errNotFound
	}

	return newNews(*news), nil
}

// Count возвращает количество новостей по указанному фильтру
//
//zenrpc:500 		server error
func (ns NewsService) Count(ctx context.Context, filter NewsFilter) (int, error) {
	count, err := ns.manager.NewsCount(ctx, &newsportal.NewsFilter{
		CategoryID: filter.CategoryID,
		TagID:      filter.TagID,
	})
	if err != nil {
		return 0, errInternalError
	}

	return count, nil
}

// Get возвращает список тэгов
//
//zenrpc:500 		server error
//zenrpc:404 		tags not found
func (ts TagService) Get(ctx context.Context) ([]Tag, error) {
	tags, err := ts.manager.Tags(ctx)
	if err != nil {
		return nil, errInternalError
	} else if len(tags) == 0 {
		return nil, errNotFound
	}

	return newTags(tags), nil
}

// Get возвращает список категорий
//
//zenrpc:500 		server error
//zenrpc:404 		categories not found
func (cs CategoryService) Get(ctx context.Context) ([]Category, error) {
	cats, err := cs.manager.Categories(ctx)
	if err != nil {
		return nil, errInternalError
	} else if len(cats) == 0 {
		return nil, errNotFound
	}

	return newCategorySlice(cats), nil
}
