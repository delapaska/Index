package main

import (
	"strings"
	"unicode/utf8"
)

const (
	MaxSliding = 64
	PrimeRK1   = 16777619
	PrimeRK2   = 2166136261
)

func SlidingWindow(s, substr string) int {

	i, j := 0, 0
	ans1, ans2 := 0, 0
	var str rune
	var str_size int
	for i < len(s) && j < len(substr) {
		str, str_size = utf8.DecodeRuneInString(s[i:])
		subst, substr_size := utf8.DecodeRuneInString(substr[j:])
		if str == subst && str_size == substr_size {
			i += str_size
			j += substr_size
			ans1++
			ans2++

		} else {
			ans1 = ans1 - ans2 + 1
			ans2 = 0
			i = i - j + str_size
			j = 0
		}
	}
	if j == len(substr) {

		return ans1 - ans2

	}
	return -1
}

func IndexRabinKarp(s, substr string) int {

	ln := utf8.RuneCountInString(s)

	subRunes := []rune(substr)
	hash1, pow1, hash2, pow2 := HashRunesDouble(subRunes)

	n := len(subRunes)
	var h1, h2 uint32
	var a rune

	var size int
	k := 0
	var ans []rune

	for i := 0; i < ln; {
		a, size = utf8.DecodeRuneInString(s[k:])
		if i < n {
			h1 = h1*PrimeRK1 + uint32(a)
			h2 = h2*PrimeRK2 + uint32(a)
			ans = append(ans, a)
		} else {
			h1 *= PrimeRK1
			h2 *= PrimeRK2

			h1 += uint32(a) - pow1*uint32(ans[0])
			h2 += uint32(a) - pow2*uint32(ans[0])

			ans = append(ans[:0], ans[1:]...)
			ans = append(ans, a)
		}
		i++
		k += size

		if h1 == hash1 && h2 == hash2 && string(ans) == substr {
			return i - n
		}
	}

	return -1
}

func Index(s, substr string) int {
	n := len(substr)
	switch {

	case n == 0:
		return 0

	case n == 1:
		return strings.IndexRune(s, rune(substr[0]))

	case n == len(s):
		if substr == s {
			return 0
		}
		return -1
	case n > len(s):
		return -1
	case len(s) <= MaxSliding:

		return SlidingWindow(s, substr)
	default:
		return IndexRabinKarp(s, substr)
	}
}

func HashRunesDouble(runes []rune) (uint32, uint32, uint32, uint32) {
	hash1, pow1 := HashRunesWithPrime(runes, PrimeRK1)
	hash2, pow2 := HashRunesWithPrime(runes, PrimeRK2)
	return hash1, pow1, hash2, pow2
}

func HashRunesWithPrime(runes []rune, PrimeRK uint32) (uint32, uint32) {
	hash := uint32(0)

	for i := 0; i < len(runes); i++ {
		hash = hash*PrimeRK + uint32(runes[i])
	}

	var pow, sq uint32 = 1, PrimeRK

	for i := len(runes); i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return hash, pow
}
