//go:generate -command blah go run gen.go dev
//go:generate blah

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
