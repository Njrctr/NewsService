package db

import (
	"context"
	g "news-service/internal/global"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IDB interface {
	Ping(ctx context.Context) error
	QueryRow(ctx context.Context, label, query string, args ...any) pgx.Row
	Query(ctx context.Context, label, query string, args ...any) (pgx.Rows, error)
	BeginTx(ctx context.Context, opts pgx.TxOptions) (pgx.Tx, error)
}

type dbEngine struct {
	db *pgxpool.Pool
}

func NewDB(ctx context.Context) (*dbEngine, error) {

	db, err := connect(ctx)
	if err != nil {
		return nil, err
	}

	newDB := &dbEngine{
		db: db,
	}

	// TODO Добавить пинг БД

	if err = checkTimeZone(ctx, g.Cfg.DB, newDB); err != nil {
		return nil, err
	}

	return newDB, nil
}

func (e *dbEngine) Ping(ctx context.Context) error {
	return e.db.Ping(ctx)
}

func (e *dbEngine) BeginTx(ctx context.Context, opts pgx.TxOptions) (pgx.Tx, error) {
	return e.db.BeginTx(ctx, opts)
}

func (e *dbEngine) Query(ctx context.Context, label, query string, args ...any) (pgx.Rows, error) {
	var err error

	r := &rows{
		label: label,
		start: time.Now(),
	}

	r.rows, _ = e.db.Query(ctx, query, args...)

	return r, err
}

func (e *dbEngine) QueryRow(ctx context.Context, label, query string, args ...any) pgx.Row {
	s := &scanner{
		label: label,
		start: time.Now(),
	}

	s.row = e.db.QueryRow(ctx, query, args...)
	return s
}

type scanner struct {
	label string
	start time.Time
	row   pgx.Row
}

func (s *scanner) Scan(targets ...any) error {
	return s.row.Scan(targets...)
}

type rows struct {
	label string
	start time.Time
	rows  pgx.Rows
}

func (r *rows) Close() {
	r.rows.Close()
}

func (r *rows) Err() error {
	return r.rows.Err()
}

func (r *rows) CommandTag() pgconn.CommandTag {
	return r.rows.CommandTag()
}

func (r *rows) FieldDescriptions() []pgconn.FieldDescription {
	return r.rows.FieldDescriptions()
}

func (r *rows) Next() bool {
	next := r.rows.Next()
	if !next {
	}

	return next
}

func (r *rows) Scan(dest ...any) error {
	err := r.rows.Scan(dest...)
	if err != nil {
	}

	return err
}

func (r *rows) Values() ([]any, error) {
	res, err := r.rows.Values()
	if err != nil {
	}

	return res, err
}

func (r *rows) RawValues() [][]byte {
	return r.rows.RawValues()
}

func (r *rows) Conn() *pgx.Conn {
	return r.rows.Conn()
}
