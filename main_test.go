package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtractCityFromTimezone(t *testing.T) {
	tables := []struct {
		Timezone       string
		ExpectedResult string
	}{
		{
			"Europe/Oslo",
			"Oslo",
		},
		{
			"America/Kentucky/Louisville",
			"Louisville",
		},
		{
			"America/St_Johns",
			"St_Johns",
		},
	}

	for _, table := range tables {
		actualResult := ExtractCityFromTimezone(table.Timezone)
		assert.Equal(t, table.ExpectedResult, actualResult)
	}
}
