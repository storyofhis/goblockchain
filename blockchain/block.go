package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	Nonce        int            `json:"nonce"`
	PrevHash     [32]byte       `json:"prevHash"`
	Timestamp    int64          `json:"timestamp"`
	Transactions []*Transaction `json:"transactions"`
}

func NewBlock(nonce int, prevHash [32]byte, transactions []*Transaction) *Block {
	return &Block{
		Timestamp:    time.Now().UnixNano(),
		Nonce:        nonce,
		PrevHash:     prevHash,
		Transactions: transactions,
	}
}

func (b *Block) Print() {
	fmt.Printf("timestamp\t: %d\n", b.Timestamp)
	fmt.Printf("nonce\t\t: %d\n", b.Nonce)
	fmt.Printf("previous_hash\t: %x\n", b.PrevHash)
	for _, t := range b.Transactions {
		t.Print()
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256(m)
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Nonce        int            `json:"nonce"`
		PrevHash     [32]byte       `json:"prevHash"`
		Timestamp    int64          `json:"timestamp"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Timestamp:    b.Timestamp,
		Nonce:        b.Nonce,
		PrevHash:     b.PrevHash,
		Transactions: b.Transactions,
	})
}
