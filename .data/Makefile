build:
	go run ./convert.go > ./tmp.go
	gofmt -w ./tmp.go
	mv ./tmp.go ../data.go
