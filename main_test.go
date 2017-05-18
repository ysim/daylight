package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtractCityFromTimezone(t *testing.T) {
	tables := []struct {
		Timezone       string
		ExpectedResult City
	}{
		{
			"Europe/Oslo",
			City{"Oslo"},
		},
		{
			"America/Kentucky/Louisville",
			City{"Louisville"},
		},
		{
			"America/St_Johns",
			City{"St Johns"},
		},
	}

	for _, table := range tables {
		actualResult := ExtractCityFromTimezone(table.Timezone)
		assert.Equal(t, table.ExpectedResult, actualResult)
	}
}
