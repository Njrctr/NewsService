package db

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DBConfig struct {
	Host        string        `toml:"host"`
	Port        int           `toml:"port"`
	Username    string        `toml:"username"`
	Password    string        `toml:"password"`
	Database    string        `toml:"database"`
	MaxConn     int32         `toml:"maxConn"`
	MinConn     int32         `toml:"minConn"`
	MaxIdleTime time.Duration `toml:"maxIdleTime"`
	TimeZone    string        `toml:"timezone"`
	DisableTLS  bool          `toml:"disableTLS"`
}

func (c *DBConfig) Validate() error {
	if c.Host == `` {
		return errors.New(`empty host`)
	}
	if c.Port == 0 {
		return errors.New(`port is zero`)
	}
	if c.Username == `` {
		return errors.New(`empty username`)
	}
	if c.Password == `` {
		return errors.New(`empty password`)
	}
	if c.Database == `` {
		return errors.New(`empty database`)
	}
	if c.MaxConn == 0 {
		return errors.New(`maxCons is zero`)
	}
	if c.MaxIdleTime < time.Second {
		return errors.New(`maxIdleTime is less than 1 second`)
	}
	if c.TimeZone == `` {
		return errors.New(`empty timezone`)
	}

	return nil
}

type IDB interface {
	Ping(ctx context.Context) error
	QueryRow(ctx context.Context, label, query string, args ...any) pgx.Row
	Query(ctx context.Context, label, query string, args ...any) (pgx.Rows, error)
	BeginTx(ctx context.Context, opts pgx.TxOptions) (pgx.Tx, error)
}
type DbEngine struct {
	db *pgxpool.Pool
}

func New(ctx context.Context, cfg *DBConfig) (*DbEngine, error) {

	db, err := connect(ctx, cfg)
	if err != nil {
		return nil, err
	}

	newDB := &DbEngine{
		db: db,
	}

	if err := pingRetrying(10, ctx, newDB.Ping); err != nil {
		return nil, err
	}

	if err = checkTimeZone(ctx, cfg, newDB); err != nil {
		return nil, err
	}

	return newDB, nil
}

func (e *DbEngine) Ping(ctx context.Context) error {
	return e.db.Ping(ctx)
}

func (e *DbEngine) BeginTx(ctx context.Context, opts pgx.TxOptions) (pgx.Tx, error) {
	return e.db.BeginTx(ctx, opts)
}

func (e *DbEngine) Query(ctx context.Context, label, query string, args ...any) (pgx.Rows, error) {
	var err error

	r := &rows{
		label: label,
		start: time.Now(),
	}

	r.rows, _ = e.db.Query(ctx, query, args...)

	return r, err
}

func (e *DbEngine) QueryRow(ctx context.Context, label, query string, args ...any) pgx.Row {
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
