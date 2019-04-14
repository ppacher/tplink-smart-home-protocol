package tpshp

import (
	"encoding/binary"
	"io"
)

// SendRaw encryptes the given payload and sends it to the TP-Link device
func SendRaw(conn io.Writer, payload []byte) error {
	cipher := Encrypt(payload)

	// the first 4 bytes contain the length of the payload in big endian
	if err := binary.Write(conn, binary.BigEndian, uint32(len(cipher))); err != nil {
		return err
	}

	n, err := conn.Write(cipher)
	if n != len(cipher) {
		// err should be set
		return err
	}

	return nil
}

// RecvRaw receives a response from the TP-Link device and decrypts it
func RecvRaw(conn io.Reader) ([]byte, error) {
	var length uint32

	// first we need to read the length of the response payload
	if err := binary.Read(conn, binary.BigEndian, &length); err != nil {
		return nil, err
	}

	var payload = make([]byte, length)

	if n, err := conn.Read(payload); n != int(length) {
		return nil, err
	}

	return Decrypt(payload), nil
}

// Encrypt a plaintext byte slice using the encryption algorithm used by TP-Link's Smart-Home Protocol
// it will also prepend the length of the ciphertext to the returned buffer
func Encrypt(plaintext []byte) []byte {
	n := len(plaintext)
	ciphertext := []byte{}

	key := byte(0xAB)
	payload := make([]byte, n)
	for i := 0; i < n; i++ {
		payload[i] = plaintext[i] ^ key
		key = payload[i]
	}

	for i := 0; i < len(payload); i++ {
		ciphertext = append(ciphertext, payload[i])
	}

	return ciphertext
}

// Decrypt a ciphertext byte slice using the decryption algorithm used by TP-Link's Smart-Home Protocol
func Decrypt(ciphertext []byte) []byte {
	n := len(ciphertext)
	key := byte(0xAB)
	var nextKey byte
	for i := 0; i < n; i++ {
		nextKey = ciphertext[i]
		ciphertext[i] = ciphertext[i] ^ key
		key = nextKey
	}
	return ciphertext
}
