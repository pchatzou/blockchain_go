package godap_user

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"log"
)

const version = byte(0x00)
const addressChecksumLen = 4

// User stores private and public keys
type User struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
	Username   string
}

// NewUser creates and returns a User
func NewUser(username string) *User {
	private, public := newKeyPair()
	user := User{private, public, username}

	return &user
}

// GetAddress returns user address
func (w User) GetAddress() string {
	pubKeyHash := HashPubKey(w.PublicKey)

	versionedPayload := append([]byte{version}, pubKeyHash[:]...)
	checksum := checksum(versionedPayload)

	fullPayload := append(versionedPayload, checksum...)
	address := Encode(fullPayload)

	return address
}

// HashPubKey hashes public key
func HashPubKey(pubKey []byte) [sha256.Size]byte {
	hash := sha256.Sum256(pubKey)
	return sha256.Sum256(hash[:])

}

// ValidateAddress check if address if valid
func ValidateAddress(address string) bool {
	pubKeyHash := Decode(address)
	actualChecksum := pubKeyHash[len(pubKeyHash)-addressChecksumLen:]
	version := pubKeyHash[0]
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-addressChecksumLen]
	targetChecksum := checksum(append([]byte{version}, pubKeyHash...))

	return bytes.Equal(actualChecksum, targetChecksum)
}

// Checksum generates a checksum for a public key
func checksum(payload []byte) []byte {
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])

	return secondSHA[:addressChecksumLen]
}

func newKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return *private, pubKey
}
