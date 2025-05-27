package main3_test

import (
	"3/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	m := services.NewMap()

	m.Add("one", 1)
	if !m.Exists("one") {
		t.Error("expected key 'one'")
	}

	value, Ok := m.Get("one")
	assert.Equal(t, Ok, true)
	assert.Equal(t, value, 1)

	copiedMap := m.Copy()
	assert.Equal(t, copiedMap["one"], 1)

	m.Remove("one")
	assert.Equal(t, m.Exists("one"), false)

	_, Ok = m.Get("one")
	assert.Equal(t, Ok, false)

}
