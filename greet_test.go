package greet

import (
    "testing"
)

func Test_sayHello(t *testing.T) {
    name := "Dude"
    got := SayHello(name)
    want := "Hello, Dude!"
    if got != want {
        t.Errorf("got '%s' want '%s'\n", got, want)
    }
}
