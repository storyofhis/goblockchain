package blockchain

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Transaction struct {
	SenderBlockchainAddress    string
	RecipientBlockchainAddress string
	value                      float32
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" sender_blockchain_address\t%s\n", t.SenderBlockchainAddress)
	fmt.Printf(" recipient_blockchain_address\t%s\n", t.RecipientBlockchainAddress)
	fmt.Printf(" value\t\t\t\t%.1f\n", t.value)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"float"`
	}{
		Sender:    t.SenderBlockchainAddress,
		Recipient: t.RecipientBlockchainAddress,
		Value:     t.value,
	})
}
