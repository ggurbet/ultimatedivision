package publicprivatekey

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

// GeneratePublicPrivateKey generate public and private key for casper login.
func GeneratePublicPrivateKey() ([]byte, []byte, error) {
	generatePrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(generatePrivateKey)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	generatePublicKey := &generatePrivateKey.PublicKey
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(generatePublicKey)
	if err != nil {
		return nil, nil, err
	}

	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	return pem.EncodeToMemory(publicKeyBlock), pem.EncodeToMemory(privateKeyBlock), nil
}

// DecryptCasperWalletAddress decrypt casper wallet address.
func DecryptCasperWalletAddress(signature string, privateKey []byte) ([]byte, error) {
	cipherText, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(privateKey)

	private, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	walletAddress, err := rsa.DecryptPKCS1v15(rand.Reader, private, cipherText)
	if err != nil {
		return nil, err
	}

	return walletAddress, nil
}
