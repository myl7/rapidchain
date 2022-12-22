package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

type PrivKey struct {
	Priv ed25519.PrivateKey
	Pub  *PubKey
}

type PubKey struct {
	Pub   ed25519.PublicKey
	Bytes [32]byte // hash of x.bytes | y.bytes
}

type Sig struct {
	B []byte
}

func (s *Sig) bytes() []byte {
	return s.B
}

func (k *PrivKey) gen() {
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	ifErrFatal(err, "ed25519 genkey")
	k.Priv = privKey
	k.Pub = &PubKey{}
	k.Pub.Pub = pubKey

	k.Pub.init()
}

func (k *PrivKey) sign(hashedMsg [32]byte) *Sig {
	s := ed25519.Sign(k.Priv, hashedMsg[:])
	return &Sig{B: s}
}

func (k *PubKey) init() {
	copy(k.Bytes[:], k.Pub)
}

func (k *PubKey) verify(hashedMsg [32]byte, sig *Sig) bool {
	return verify(k.Pub, hashedMsg, sig)
}

func (k *PubKey) string() string {
	return hex.EncodeToString(k.Bytes[:])
}

func verify(pubKey ed25519.PublicKey, hashedMsg [32]byte, sig *Sig) bool {
	return ed25519.Verify(pubKey, hashedMsg[:], sig.B)
}

func hash(msg []byte) [32]byte {
	h := sha256.New()
	h.Write(msg)
	hashedMsg := h.Sum(nil)
	if len(hashedMsg) != 32 {
		errFatal(nil, "hash not 32 bytes")
	}
	return toByte32(hashedMsg)
}
