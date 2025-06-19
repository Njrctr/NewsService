genna:
	genna model -c "postgres://newsuser:akgj123cguygecuw3riu1y23@localhost:5432/news-db?sslmode=disable" -o ./internal/db/model.go -t "public.*" -f
lint:
	golangci-lint run