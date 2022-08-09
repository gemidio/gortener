package bitlink

import (
	"time"
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
		id:        f.identifier.Random(),
		url:       url,
		createdAt: now,
		updatedAt: now,
	}, nil
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
