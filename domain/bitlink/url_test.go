package bitlink

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidURL(t *testing.T) {
	url := URL("https://www.google.com")
	actual := url.isValid()

	assert.True(t, actual)
}

func TestInvalidURL(t *testing.T) {

	t.Run("given empty URL", func(t *testing.T) {
		url := URL("")
		actual := url.isValid()

		assert.False(t, actual)
	})

	t.Run("given URL with invalid protocol", func(t *testing.T) {
		url := URL("htt://www.google.com")
		actual := url.isValid()

		assert.False(t, actual)
	})

	t.Run("given URL with invalid format", func(t *testing.T) {
		url := URL("://www.google.com")
		actual := url.isValid()

		assert.False(t, actual)
	})
}
