package greetings_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/example/module/src/greetings"
)

func TestGreeting(t *testing.T) {
	assert.NotEqual(t, greetings.Greeting(), "")
}
