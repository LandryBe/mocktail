package a

import (
	"bytes"
	"context"
	"time"

	"golang.org/x/mod/module"
)

type Pineapple interface {
	Hello(bar Water) string
	World() string
	Goo() (string, int, Water)
	Coo(context.Context, string, Water) Water
}

type Coconut interface {
	Boo(src *bytes.Buffer) time.Duration
	Doo(src time.Duration) time.Duration
	Foo(st Strawberry) string
	Goo(st string) Strawberry
	Hoo(string, int, Water)
	Joo(string, int, Water) (string, int)
	Koo(src string) (dst string)
	Loo(st string, values ...int) string
	Too(src string) time.Duration
	Voo(src *module.Version) time.Duration
	Yoo(st string) interface{}
	Zoo(st interface{}) string
}

type Water struct{}

type Strawberry interface {
	Bar(string) int
}
