package rsaz

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

const (
	privateKeyHeader = "PRIVATE KEY"
	publicKeyHeader  = "PUBLIC KEY"
)

// PEMToRSAPrivateKey decodes an RSA private key in PEM format.
func PEMToRSAPrivateKey(buf []byte) (*rsa.PrivateKey, error) {
	block, rest := pem.Decode(buf)
	if block == nil {
		return nil, fmt.Errorf("invalid PEM block")
	}
	if len(rest) > 0 {
		return nil, fmt.Errorf("unexpected number of PEM blocks")
	}
	if block.Type != privateKeyHeader {
		return nil, fmt.Errorf("unexpected PEM block type")
	}

	rawKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaKey, ok := rawKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("unexpected key type")
	}

	return rsaKey, nil
}

// MustPEMToRSAPrivateKey is like PEMToRSAPrivateKey but panics on error.
func MustPEMToRSAPrivateKey(buf []byte) *rsa.PrivateKey {
	key, err := PEMToRSAPrivateKey(buf)
	if err != nil {
		panic(err)
	}
	return key
}

// PEMToRSAPublicKey decodes an RSA public key in PEM format.
func PEMToRSAPublicKey(buf []byte) (*rsa.PublicKey, error) {
	block, rest := pem.Decode(buf)
	if block == nil {
		return nil, fmt.Errorf("invalid PEM block")
	}
	if len(rest) > 0 {
		return nil, fmt.Errorf("unexpected number of PEM blocks")
	}
	if block.Type != publicKeyHeader {
		return nil, fmt.Errorf("unexpected PEM block type")
	}

	rawKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaKey, ok := rawKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("unexpected key type")
	}

	return rsaKey, nil
}

// MustPEMToRSAPublicKey is like PEMToRSAPublicKey but panics on error.
func MustPEMToRSAPublicKey(buf []byte) *rsa.PublicKey {
	key, err := PEMToRSAPublicKey(buf)
	if err != nil {
		panic(err)
	}
	return key
}

// RSAPrivateKeyToPEM encodes an RSA private key to PEM format.
func RSAPrivateKeyToPEM(key *rsa.PrivateKey) []byte {
	buf, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		panic(err) // never happens because we already checked the key type
	}

	return pem.EncodeToMemory(&pem.Block{
		Type:  privateKeyHeader,
		Bytes: buf,
	})
}

// RSAPublicKeyToPEM encodes a RSA public key to PEM format.
func RSAPublicKeyToPEM(key *rsa.PublicKey) []byte {
	buf, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		panic(err) // never happens because we already checked the key type
	}

	return pem.EncodeToMemory(&pem.Block{
		Type:  publicKeyHeader,
		Bytes: buf,
	})
}
