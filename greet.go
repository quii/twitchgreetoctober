package twitchgreetoctober

import "fmt"

type GreetAdapter func(name string) string

func (g GreetAdapter) Greet(name string) (string, error) {
	return g(name), nil
}

func Greet(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
