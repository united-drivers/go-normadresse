# go-normadresse

[![GoDoc](https://godoc.org/github.com/united-drivers/go-normadresse?status.svg)](https://godoc.org/github.com/united-drivers/go-normadresse)

French Postal Addresses Normalizer (Golang port of [etalab/normadresse](https://github.com/etalab/normadresse))

## Example

```go
import "github.com/united-drivers/go-normadresse"

fmt.Println(strings.ToUpper(normadresse.Normalize("BOULEVARD DU MARECHAL JEAN MARIE DE LATTRE DE TASSIGNY")))
// Output: BD MAL J M DE LATTRE DE TASSIGNY
```
