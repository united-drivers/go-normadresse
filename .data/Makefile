build:
	go run ./convert.go > ./tmp.go
	gofmt -w ./tmp.go
	mv ./tmp.go ../rules.gen.go

update:
	wget -N https://raw.githubusercontent.com/etalab/normadresse/master/normadresse.csv
