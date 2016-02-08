package statelessPassword //import "go.iondynamics.net/statelessPassword"

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"

	"golang.org/x/crypto/scrypt"
)

type Algorithm struct {
	variant        uint8
	fullname       []byte
	masterpassword []byte
	n              int
	r              int
	p              int
	keyLen         int
	saltPrefix     []byte
	key            []byte
}

func New(fullname, masterpassword []byte, variant uint8) *Algorithm {
	algo := &Algorithm{
		variant:    variant,
		fullname:   fullname,
		r:          8,
		keyLen:     128,
		p:          4,
		saltPrefix: []byte("go.iondynamics.net/statelessPassword"),
	}

	switch variant {
	case 0:
		algo.n = 32768
		algo.p = 2
		algo.keyLen = 64
		algo.saltPrefix = []byte("com.lyndir.masterpassword")
	case 1:
		algo.n = 32768
		algo.p = 2
	case 2:
		algo.n = 32768
	case 3:
		algo.n = 65536
	case 4:
		algo.n = 131072
	case 5:
		algo.n = 131072
		algo.r = 12
	case 6:
		algo.n = 1048576
		algo.p = 1
	case 7:
		algo.n = 1048576
	}

	var err error
	algo.key, err = scrypt.Key(masterpassword, append(algo.saltPrefix, []byte(fmt.Sprint(uint32(len(algo.fullname)), algo.fullname))...), algo.n, algo.r, algo.p, algo.keyLen)
	if err != nil {
		return nil
	}

	return algo
}

func (algo *Algorithm) Password(site string, version string, templates []string) (string, error) {
	if len(templates) < 1 {
		return "", fmt.Errorf("%s", "invalid template")
	}

	h := hmac.New(sha512.New, algo.key)
	if algo.variant == 0 {
		h = hmac.New(sha256.New, algo.key)
	}

	_, err := h.Write([]byte(fmt.Sprint(string(algo.saltPrefix), uint32(len(site)), site, version)))
	if err != nil {
		return "", err
	}
	seed := h.Sum(nil)

	template := templates[uint32(seed[0])%uint32(len(templates))]

	password := []byte(template)

	for i, tplRune := range template {
		var passchars string
		switch tplRune {
		case 'V':
			passchars = V
		case 'C':
			passchars = C
		case 'v':
			passchars = v
		case 'c':
			passchars = c
		case 'A':
			passchars = A
		case 'a':
			passchars = a
		case 'n':
			passchars = n
		case 'o':
			passchars = o
		case 'X':
			passchars = X
		case 'x':
			passchars = x
		case 'p':
			passchars = p
		default:
			return "", fmt.Errorf("%s", "invalid template")
		}

		if i+1 < len(seed) {
			password[i] = passchars[uint32(seed[i+1])%uint32(len(passchars))]
		} else {
			password[i] = passchars[0]
		}

	}

	return string(password), err

}
