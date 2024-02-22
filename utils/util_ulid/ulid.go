package utilulid

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

// Generate Ulid to create new ID
func GetUlid() string {
	time := time.Now().UTC()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(time), entropy).String()
	
	return id
}