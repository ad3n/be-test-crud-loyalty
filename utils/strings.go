package utils

import "math/rand"

const LETTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Strings struct {
}

func (s Strings) Random(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = LETTERS[rand.Intn(len(LETTERS))]
	}

	return string(b)
}
