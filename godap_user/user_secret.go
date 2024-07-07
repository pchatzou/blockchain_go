package godap_user

// raw content, enc by public key
// enc content, decr by private key
// Encrypted user secret, has to be owned by the operation runner
type EncSecret struct {
	Content *[]byte
	User    *User
}

// make an encrypt, decrypt

func (es *EncSecret) ReEncrypt(pubKey []byte) error {

	return nil
}

func (es *EncSecret) Decrypt() error {
	return nil
}

func (es *EncSecret) CanDecrypt() error {
	return nil
}
