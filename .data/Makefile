build:
	go run ./convert.go > ./tmp.go
	gofmt -w ./tmp.go
	mv ./tmp.go ../data.go

update:
	wget -N https://raw.githubusercontent.com/etalab/normadresse/master/normadresse.csv
