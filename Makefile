make ts-types:
	go run cmd/clean_ts.go apps/client/src/generated/db

make start:
	go build -o out && ./out
