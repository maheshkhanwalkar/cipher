package types

type Cipher interface {
	// Encrypt the data using the given key, returning the ciphertext
	Encrypt(data string, key string) string

	// Decrypt the data using the given key, returning the plaintext
	Decrypt(data string, key string) string
}

var ciphers = make(map[string]Cipher)

// Get the cipher by name
func GetCipherByName(name string) Cipher {
	cipher, good := ciphers[name]

	// No cipher by that name -- return nil
	if !good {
		return nil
	}

	return cipher
}

// Register a cipher with the provided name
func RegisterCipher(name string, cipher Cipher) bool {
	_, good := ciphers[name]

	// Cipher already exists by that name -- reject the registration
	if good {
		return false
	}

	ciphers[name] = cipher
	return true
}

// Setup and register the builtin ciphers
func SetupBuiltinCiphers() {
	RegisterCipher("caesar", &CaesarCipher{})
}
