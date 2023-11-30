package gommand

import (
	"errors"
	"fmt"
	"testing"

	"github.com/sampgo/sampgo"
)

func TestEntryPoint(t *testing.T) {
	cmd := NewCommand(Command{
		Name: "f",
	})

	cmd.SetName("foo")

	cmd.SetAlias("bar", "baz")

	err := cmd.Handle(func(ctx Context) (err error) {
		fmt.Println("command fired!")
		fmt.Println(ctx)
		return
	})

	if err != nil {
		t.Error(err)
	}

	SetGeneralCommandBeforeFunc(func(ctx Context) (err error) {
		fmt.Println("before commadn gets fired!")
		fmt.Println(ctx)
		return
	})

	SetGeneralCommandAfterFunc(func(ctx Context) (err error) {
		fmt.Println("after command gets fired!")
		fmt.Println(ctx)
		fmt.Println()
		return
	})

	fail := errors.New("handler call failed")

	success := handler(sampgo.Player{ID: 0}, "/foo first handler call")
	if !success {
		t.Error(fail)
	}

}
