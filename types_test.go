package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestNewAccount(t *testing.T) {
	acc, err := NewAccount("a", "b", "hunter")
	assert.Nil(t, err)

	fmt.Printf("%+v\n)", acc)
}
