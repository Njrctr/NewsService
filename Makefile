lint:
	golangci-lint run
mfd-xml:
	@mfd-generator xml -c "postgres://newsuser:akgj123cguygecuw3riu1y23@localhost:5432/news-db?sslmode=disable" -m ./docs/model/newsportal.mfd -n "news:news,categories,tags"

mfd-model:
	@mfd-generator model -m ./docs/model/newsportal.mfd -p db -o ./internal/db

mfd-repo:
	@mfd-generator repo -m ./docs/model/newsportal.mfd -p db -o ./internal/db