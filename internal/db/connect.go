package db

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func connect(ctx context.Context, cfg *Config) (*pgxpool.Pool, error) {
	pcfg, err := pgxpool.ParseConfig(formConnectionString(cfg))
	if err != nil {
		return nil, err
	}

	pcfg.AfterConnect = afterConnect(cfg)
	pcfg.MaxConns = cfg.MaxConn
	pcfg.MinConns = cfg.MinConn
	pcfg.MaxConnIdleTime = cfg.MaxIdleTime

	db, err := pgxpool.NewWithConfig(ctx, pcfg)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func formConnectionString(cfg *Config) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port,
		cfg.Database,
	)
}

func pingRetrying(retries int, ctx context.Context, fn func(context.Context) error) error {
	var err error
	for x := range retries {
		log.Printf("Try to connect DB. Try #%d", x+1)
		if err = fn(ctx); err != nil {
			fmt.Println(err)
			<-time.After(1 * time.Second)
			continue
		}
		return nil
	}
	return err
}

type afterFunc func(ctx context.Context, conn *pgx.Conn) error

func afterConnect(cfg *Config) afterFunc {
	return func(ctx context.Context, conn *pgx.Conn) error {
		if _, err := conn.Exec(ctx, fmt.Sprintf(`set time zone '%s';`, cfg.TimeZone)); err != nil {
			return err
		}

		return nil
	}
}

func checkTimeZone(ctx context.Context, cfg *Config, db IDB) error {
	var tz string

	if err := db.QueryRow(ctx, `check_time_zone`, `select current_setting('TIMEZONE')`).Scan(&tz); err != nil {
		return err
	}

	if !strings.EqualFold(tz, cfg.TimeZone) {
		return fmt.Errorf(`db server timezone not match with pgsql config. cfg tz: %s, server tz: %s`,
			cfg.TimeZone, tz)
	}

	return nil
}
