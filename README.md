# go-normadresse

[![CircleCI](https://circleci.com/gh/united-drivers/go-normadresse.svg?style=svg)](https://circleci.com/gh/united-drivers/go-normadresse)
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

More usage on [GoDoc](https://godoc.org/github.com/united-drivers/go-normadresse), or in the [test file](https://github.com/united-drivers/go-normadresse/blob/master/normadresse_test.go).

## Disclaimer

The original version written in Python uses [more efficient rule priorities](https://github.com/etalab/normadresse/blob/master/normadresse.py) (than [ours](https://github.com/united-drivers/go-normadresse/blob/master/normadresse.go)), which often results in better outputs

Example:
* Input: `SQUARE DES SOEURS DE SAINT VINCENT DE PAUL`
* Python: `SQ SOEURS DE ST VINCENT DE PAUL` (easy to read)
* Golang: `SQ DES SOEURS DE SAINT V DE P` (correct, but harder to read)

But it gives the exactly same result for others:
* `BOULEVARD DU MARECHAL JEAN MARIE DE LATTRE DE TASSIGNY` -> `BD MAL J M DE LATTRE DE TASSIGNY`
* `RUE DES FRERES MARC ET JEAN MARIE GAMON` -> `RUE DES FRERES M ET J M GAMON`
* `AVENUE GEORGES ET CLAUDE CAUSTIER` -> `AV GEORGES ET CLAUDE CAUSTIER`
* `AVENUE DE LA DIVISION DU GENERAL LECLERC` -> `AV DE LA DIVISION DU GAL LECLERC`
* `RUE EMMANUEL D ASTIER DE LA VIGERIE` -> `RUE E D ASTIER DE LA VIGERIE`
* ...

## License

[Apache 2.0](https://github.com/united-drivers/go-normadresse/blob/master/LICENSE)
