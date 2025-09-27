package main

import (
	"fmt"
	"math"
)

type KarpRabin struct {
	prime float64
}

func NewKarpRabin() *KarpRabin {
	return &KarpRabin{
		prime: 101,
	}
}

// calculate hash of the string
func (k *KarpRabin) calculateHash(str string) float64 {
	var hash float64
	for i := 0; i < len(str); i++ {
		hash += float64(str[i]) * math.Pow(k.prime, float64(i))
	}
	return hash
}

// rolls the hash to next substring
func (k *KarpRabin) updateHash(prevHash float64, oldChar byte, newChar byte, patternLength int) float64 {
	newHash := (prevHash - float64(oldChar)) / k.prime
	newHash += float64(newChar) * math.Pow(k.prime, float64(patternLength-1))
	return newHash
}

// search looks for pattern in text
func (k *KarpRabin) Search(text, pattern string) {
	patternLength := len(pattern)
	if patternLength > len(text) {
		return
	}

	patternHash := k.calculateHash(pattern)
	textHash := k.calculateHash(text[:patternLength])

	for i := 0; i <= len(text)-patternLength; i++ {
		if textHash == patternHash {
			if text[i:i+patternLength] == pattern {
				fmt.Println("Pattern found at index", i)
			}
		}
		if i < len(text)-patternLength {
			textHash = k.updateHash(textHash, text[i], text[i+patternLength], patternLength)
		}
	}
}

func main() {
	k := NewKarpRabin()
	text := "abracadabra"
	pattern := "ca"
	k.Search(text, pattern)
}
