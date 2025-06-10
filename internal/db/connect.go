package db

import (
	"context"
	"fmt"
	"news-service/internal/configs"
	g "news-service/internal/global"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func connect(ctx context.Context) (*pgxpool.Pool, error) {
	pcfg, err := pgxpool.ParseConfig(formConnectionString(g.Cfg.DB))
	if err != nil {
		return nil, err
	}

	pcfg.AfterConnect = afterConnect(g.Cfg.DB)
	pcfg.MaxConns = g.Cfg.DB.MaxConn
	pcfg.MinConns = g.Cfg.DB.MinConn
	pcfg.MaxConnIdleTime = g.Cfg.DB.MaxIdleTime

	db, err := pgxpool.NewWithConfig(ctx, pcfg)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func formConnectionString(cfg *configs.DB) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port,
		cfg.Database,
	)
}

type afterFunc func(ctx context.Context, conn *pgx.Conn) error

func afterConnect(cfg *configs.DB) afterFunc {
	return func(ctx context.Context, conn *pgx.Conn) error {
		if _, err := conn.Exec(ctx, fmt.Sprintf(`set time zone '%s';`, cfg.TimeZone)); err != nil {
			return err
		}

		return nil
	}
}

func checkTimeZone(ctx context.Context, cfg *configs.DB, db IDB) error {
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
