package bitlink

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewBitLink(t *testing.T) {

	factory := Factory{
		identifier: IdentifierStub{},
	}

	t.Run("given an empty url", func(t *testing.T) {
		_, err := factory.Make("")

		assert.Error(t, err, ErrInvalidURL)
	})

	t.Run("given a valid url", func(t *testing.T) {
		b, err := factory.Make("www.google.com")

		assert.Nil(t, err)
		assert.IsType(t, &BitLink{}, b)
		assert.Equal(t, URL("www.google.com"), b.URL())
		assert.Equal(t, "some-identifier", b.Id())
		assert.IsType(t, &time.Time{}, b.CreatedAt())
		assert.IsType(t, &time.Time{}, b.UpdatedAt())
	})
}

type IdentifierStub struct{}

func (i IdentifierStub) Random() string {
	return "some-identifier"
}

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
