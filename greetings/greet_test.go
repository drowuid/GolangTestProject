package greetings

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Pedro")
	want := "Hello from the greetings package, Pedro"

	if got !=want{
	t.Errorf("Hello() = %q, want %q", got, want)
}
}
