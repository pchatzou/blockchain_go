package godap_blockchain

import (
	"bytes"

	"github.com/pchatzou/godap/godap_user"
)

// TXInput represents a transaction input
type TXInput struct {
	Txid       []byte
	UserSecret godap_user.EncSecret
	Signature  []byte
	PubKey     []byte
}

// UsesKey checks whether the address initiated the transaction
func (in *TXInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := godap_user.HashPubKey(in.PubKey)

	return bytes.Equal(lockingHash[:], pubKeyHash)
}
