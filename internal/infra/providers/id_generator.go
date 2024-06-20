package providers

import (
	"crypto/rand"
	"github.com/oklog/ulid/v2"
	"time"
)

func NewID() string {
	entropy := ulid.Monotonic(rand.Reader, 0)
	return ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
}
