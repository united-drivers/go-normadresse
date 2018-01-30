# go-normadresse

[![GoDoc](https://godoc.org/github.com/united-drivers/go-normadresse?status.svg)](https://godoc.org/github.com/united-drivers/go-normadresse)
[![Go Report Card](https://goreportcard.com/badge/github.com/united-drivers/go-normadresse)](https://goreportcard.com/report/github.com/united-drivers/go-normadresse)
[![License](https://img.shields.io/github/license/united-drivers/go-normadresse.svg)](https://github.com/united-drivers/go-normadresse/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/united-drivers/go-normadresse.svg)](https://github.com/united-drivers/go-normadresse/releases)

French Postal Addresses Normalizer (Golang port of [etalab/normadresse](https://github.com/etalab/normadresse))

## Usage

Install with `go get github.com/united-drivers/go-normadresse`

```go
import "github.com/united-drivers/go-normadresse"

fmt.Println(strings.ToUpper(normadresse.Normalize("BOULEVARD DU MARECHAL JEAN MARIE DE LATTRE DE TASSIGNY")))
// Output: BD MAL J M DE LATTRE DE TASSIGNY
```
