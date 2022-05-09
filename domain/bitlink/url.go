package bitlink

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidURL error = errors.New("url is invalid")
)

type URL string

func (u URL) isValid() bool {
	pattern := "^(http:\\/\\/www\\.|https:\\/\\/www\\.|http:\\/\\/|https:\\/\\/)?[a-z0-9]+([\\-\\.]{1}[a-z0-9]+)*\\.[a-z]{2,5}(:[0-9]{1,5})?(\\/.*)?$"
	isValid, _ := regexp.Match(pattern, []byte(u))

	return isValid
}
