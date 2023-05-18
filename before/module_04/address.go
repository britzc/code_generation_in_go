//go:generate go run gen.go

package module

type Address struct {
	ID      int
	Number  string
	Street  string
	Suburb  string
	City    string
	Country string
	Code    string
}

type Addresses []*Address
