// Copyright 2011-2014 Dmitry Chestnykh. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package captcha

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"io"
)

// rngKey is a secret key used to deterministically derive seeds for
// PRNGs used in image and audio. Generated once during initialization.
var rngKey [32]byte

func init() {
	if _, err := io.ReadFull(rand.Reader, rngKey[:]); err != nil {
		panic("captcha: error reading random source: " + err.Error())
	}
}

// Purposes for seed derivation. The goal is to make deterministic PRNG produce
// different outputs for images and audio by using different derived seeds.
const (
	imageSeedPurpose = 0x01
)

// deriveSeed returns a 16-byte PRNG seed from rngKey, purpose, id and digits.
// Same purpose, id and digits will result in the same derived seed for this
// instance of running application.
//
//   out = HMAC(rngKey, purpose || id || 0x00 || digits)  (cut to 16 bytes)
//
func deriveSeed(purpose byte, id string, digits []byte) (out [16]byte) {
	var buf [sha256.Size]byte
	h := hmac.New(sha256.New, rngKey[:])
	h.Write([]byte{purpose})
	io.WriteString(h, id)
	h.Write([]byte{0})
	h.Write(digits)
	sum := h.Sum(buf[:0])
	copy(out[:], sum)
	return
}
