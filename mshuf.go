package mshuf

import (
	"encoding/binary"
	"io"
	"math/rand"
)

// Matrix a matrix for mshuf
type Matrix [16][16]byte

// Shuffle shuffle a uint64
func (m Matrix) Shuffle(n uint64) uint64 {
	b := make([]byte, 8, 8)
	binary.BigEndian.PutUint64(b, n)
	for i := 0; i < 8; i++ {
		d := b[i]
		b[i] = m[i*2][d>>4]<<4 + m[i*2+1][d&0x0f]
	}
	return binary.BigEndian.Uint64(b)
}

// RandMatrix create a new randomized matrix for mshuf with crypto/rand
func RandMatrix(m *Matrix, r io.Reader) error {
	// seeds
	s := make([]byte, len(m)*8, len(m)*8)
	_, err := r.Read(s)
	if err != nil {
		return err
	}
	// make
	for i := 0; i < len(m); i++ {
		randSequence(int64(binary.BigEndian.Uint64(s[i*8:(i+1)*8])), &m[i])
	}
	return nil
}

func randSequence(seed int64, seq *[16]byte) {
	// fill sequence
	for i := 0; i < len(seq); i++ {
		seq[i] = byte(i)
	}
	// Fisher-Yates shuffle
	rnd := rand.New(rand.NewSource(seed))
	for i := len(seq) - 1; i > 0; i = i - 1 {
		j := rnd.Intn(i)
		k := seq[i]
		seq[i] = seq[j]
		seq[j] = k
	}
	return
}
