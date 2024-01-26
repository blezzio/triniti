package usecases

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"math/rand"
	"testing"
)

func randomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func Test_ShortHash_Collision(t *testing.T) {
	seen := make(map[string]string)
	sh := NewHasher(sha1.New())
	i := 0
	for {
		str := randomString(rand.Intn(225))
		sh := sh.Hash(str)
		sha := sh.Next()
		if v, ok := seen[sha]; ok && v != str {
			t.Logf("SHA1 %s collided %s after %d with %s\n", str, v, i, sha)
			break
		}
		seen[sha] = str
		i++
		if i == 1_000_000_000 {
			break
		}
	}

	seen = make(map[string]string)
	sh = NewHasher(sha256.New())
	i = 0
	for {
		str := randomString(rand.Intn(225))
		sh := sh.Hash(str)
		sha := sh.Next()
		if v, ok := seen[sha]; ok && v != str {
			t.Logf("SHA256 %s collided %s after %d with %s\n", str, v, i, sha)
			break
		}
		seen[sha] = str
		i++
		if i == 1_000_000_000 {
			break
		}
	}

	seen = make(map[string]string)
	sh = NewHasher(sha512.New())
	i = 0
	for {
		str := randomString(rand.Intn(225))
		sh := sh.Hash(str)
		sha := sh.Next()
		if v, ok := seen[sha]; ok && v != str {
			t.Logf("SHA512 %s collided %s after %d with %s\n", str, v, i, sha)
			break
		}
		seen[sha] = str
		i++
		if i == 1_000_000_000 {
			break
		}
	}
}
