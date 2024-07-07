package godap_blockchain

import (
	"bytes"
	"os"
	"testing"
)

func setup() string {
	var err error
	// Set up a temporary directory
	temp_dir, err := os.MkdirTemp("", "_godap_test")
	if err != nil {
		panic("Failed to create temporary directory: " + err.Error())
	}
	return temp_dir
}

// Teardown function
func teardown(t string) {
	// Clean up the directory after the tests
	if err := os.RemoveAll(t); err != nil {
		panic("Failed to remove temporary directory: " + err.Error())
	}
}

func TestMain(m *testing.M) {
	temp_dir := setup()
	code := m.Run()
	teardown(temp_dir)
	os.Exit(code)
}

func TestCreateBlockChain(t *testing.T) {
	bc := CreateBlockchain("123", "hey")
	exp := []byte{1, 2, 3, 4}
	if !(bytes.Equal(bc.Iterator().Next().Serialize(), exp)) {
		t.Errorf("Fooobar")

	}

}
