// Package diffiehellman implements a solution for the exercise titled `Diffie Hellman'.
package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

var rnd *rand.Rand

func init() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// PrivateKey returns a private key k where 1 < k < p (given prime number).
func PrivateKey(p *big.Int) *big.Int {
	return new(big.Int).Add(new(big.Int).Rand(rnd, new(big.Int).Sub(p, big.NewInt(2))), big.NewInt(2))
}

// PublicKey returns a public key with private key , given p and g where P = g**private mod p.
func PublicKey(private, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), private, p)
}

// NewPair returns private and public key with given prime p and g.
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return
}

// SecretKey returns secret key with own private and peer public key.
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}
