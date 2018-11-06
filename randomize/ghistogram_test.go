package randomize

import (
	"github.com/couchbaselabs/ghistogram"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestHistogram(t *testing.T) {
	random := rand.New(rand.NewSource(42));
	histo := ghistogram.NewHistogram(10, 100, 0)
	for j := 0; j < 100000; j++ {
		histo.Add(uint64(random.Int63n(1000)), 1)
	}

	assert.Equal(t, uint64(100000), histo.TotCount)
	for _, count := range histo.Counts {
		assert.True(t, count > 9700, "count %v is less than %v", count, 9700);
		assert.True(t, count < 10300, "count %v is higher than %v", count, 10300);
	}
}
