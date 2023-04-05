package bitlink

import (
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
)

type Identifier interface {
	Random() string
}

type Factory struct {
	identifier Identifier
}

func (f *Factory) Make(url URL) (*BitLink, error) {
	if !url.isValid() {
		return &BitLink{}, ErrInvalidURL
	}

	now := time.Now()

	return &BitLink{
		id:        uuid.NewString(),
		url:       url,
		createdAt: now,
		updatedAt: now,
	}, nil
}

var (
	ErrInvalidURL error = errors.New("url is invalid")
)

type URL string

func (u URL) isValid() bool {
	pattern := "^(http:\\/\\/www\\.|https:\\/\\/www\\.|http:\\/\\/|https:\\/\\/)?[a-z0-9]+([\\-\\.]{1}[a-z0-9]+)*\\.[a-z]{2,5}(:[0-9]{1,5})?(\\/.*)?$"
	isValid, _ := regexp.Match(pattern, []byte(u))

	return isValid
}

type BitLink struct {
	id        string
	url       URL
	createdAt time.Time
	updatedAt time.Time
}

func (b *BitLink) Id() string {
	return b.id
}

func (b *BitLink) URL() URL {
	return b.url
}

func (b *BitLink) CreatedAt() *time.Time {
	return &b.createdAt
}

func (b *BitLink) UpdatedAt() *time.Time {
	return &b.updatedAt
}
