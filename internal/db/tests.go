package db

func TestDBCfg() *DBConfig {
	cfg := &DBConfig{
		Host:        "localhost",
		Port:        5432,
		Username:    "newsuser",
		Password:    "akgj123cguygecuw3riu1y23",
		Database:    "news-db",
		MaxConn:     300,
		MinConn:     10,
		MaxIdleTime: 10,
		TimeZone:    "Europe/Moscow",
		DisableTLS:  true,
	}

	return cfg
}
