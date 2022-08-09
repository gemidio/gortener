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

func (h IdentifierStub) Random() string {
	return "some-identifier"
}
