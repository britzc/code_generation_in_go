//go:generate -command counter go run gen.go DEV
//go:generate counter DEV

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
