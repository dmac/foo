package foo

import "github.com/dmac/bar"

func Foo() string {
	return "Foo" + bar.Bar()
}
