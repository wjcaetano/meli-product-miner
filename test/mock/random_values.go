package mock

import (
	"math/rand"
	"strconv"
	"time"
)

func RandomInt() int {
	return rand.Intn(99999) // nolint:gosec
}

func RandomInt64() int64 {
	return int64(rand.Intn(99999)) // nolint:gosec
}

func RandomBool() bool {
	return RandomInt() > 0
}

func RandomUint() uint {
	return uint(rand.Intn(99999)) // nolint:gosec
}

func RandomUint64() uint64 {
	return uint64(rand.Intn(99999)) // nolint:gosec
}

func RandomString() string {
	return strconv.Itoa(RandomInt())
}

func RandomFloat64() float64 {
	return float64(RandomInt())
}

func RandomDate() time.Time {
	return time.Now().UTC().AddDate(0, 0, RandomInt())
}
