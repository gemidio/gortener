package bitlink

import (
	"time"
)

type Factory struct {
	hash Hash
}

func (f *Factory) Make(url URL) (*BitLink, error) {
	if !url.isValid() {
		return &BitLink{}, ErrInvalidURL
	}

	now := time.Now()

	return &BitLink{
		hash:        f.hash.Random(),
		destination: url,
		createdAt:   now,
		updatedAt:   now,
	}, nil
}

type BitLink struct {
	hash        string
	destination URL
	createdAt   time.Time
	updatedAt   time.Time
}

func (b *BitLink) Hash() string {
	return b.hash
}

func (b *BitLink) Destination() URL {
	return b.destination
}

func (b *BitLink) CreatedAt() *time.Time {
	return &b.createdAt
}

func (b *BitLink) UpdatedAt() *time.Time {
	return &b.updatedAt
}
