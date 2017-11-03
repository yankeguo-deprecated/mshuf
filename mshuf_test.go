package mshuf

import (
	"encoding/binary"
	"math/rand"
	"testing"
)

func TestMatrix_ShuffleIdentity(t *testing.T) {
	m := Matrix{}
	// seed matrix with identity
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			m[i][j] = byte(j)
		}
	}
	for i := 0; i < 100; i++ {
		b := make([]byte, 8, 8)
		rand.Read(b)
		n := binary.BigEndian.Uint64(b)
		r := m.Shuffle(n)
		if r != n {
			t.Errorf("Shuffle failed, expect %v = %v", n, r)
		}
	}
}
