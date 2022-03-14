package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

const (
	Difficulty = 17
)

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func (p *ProofOfWork) CreateHash(nonce int) []byte {
	btss := [][]byte{
		p.Block.PrevHash,
		p.Block.Data,
		ToHex(uint64(nonce)),
		ToHex(uint64(Difficulty)),
	}
	sep := []byte{}
	jbts := bytes.Join(btss, sep)
	hash := sha256.Sum256(jbts)
	return hash[:]
}

func (p *ProofOfWork) Run() ([]byte, int) {
	nonce := 0
	var hash []byte
	var intHash big.Int
	for nonce < math.MaxInt64 {
		hash = p.CreateHash(nonce)
		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash)
		if intHash.Cmp(p.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()
	return hash, nonce
}

func (p *ProofOfWork) Valid() bool {
	hash := p.CreateHash(p.Block.Nonce)
	var intHash big.Int
	intHash.SetBytes(hash)
	return intHash.Cmp(p.Target) == -1
}

func ToHex(n uint64) []byte {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, n); err != nil {
		log.Panic(err)
	}
	return buf.Bytes()
}

func NewProof(b *Block) *ProofOfWork {
	t := big.NewInt(1)
	t.Lsh(t, uint(256-Difficulty))
	return &ProofOfWork{b, t}
}
