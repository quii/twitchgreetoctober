package twitchgreetoctober

import (
	"testing"
	"twitchgreetoctober/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(t, GreetAdapter(Greet))
}
