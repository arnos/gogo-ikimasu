package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHarōWārudo(t *testing.T) {

	loadConfiguration("helloworld.toml")

	h := world{}

	helloWorld := h.world()

	assert.Equal(t, helloWorld, "HelloWorld!", "We didn't get HelloWorld")
}
