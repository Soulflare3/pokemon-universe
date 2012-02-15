package position

import (
	"testing"
)

type positionHashTest struct {
	x, y, z int
}

var positionHashTests = []positionHashTest {
	positionHashTest{ 6543, 6543, 1234 },
	positionHashTest{ 11, 22, 35 },
}

func TestHash(t *testing.T) {
	for _, ht := range positionHashTests {
		v := Hash(ht.x, ht.y, ht.z)
		pos := NewPositionFromHash(v)
		
		if pos.X != ht.x || pos.Y != ht.y || pos.Z != ht.z {
			t.Errorf("Failed to convert back (%d, %d, %d) to (%d, %d, %d)", pos.X, pos.Y, pos.Z, ht.x, ht.y, ht.z)
		}
	}
}