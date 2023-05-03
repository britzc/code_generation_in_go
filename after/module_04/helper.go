//go:generate go run gen.go hello world

package module

type Person struct {
	ID        int
	FirstName string
	LastName  string
}
