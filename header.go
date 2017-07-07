package main

type Algorithm string

const (
	ALG_RSA1_5 = Algorithm("RSA1_5")
)

type EncryptionMethod string

const (
	ENC_A128CBC_HS256 = EncryptionMethod("A128CBC-HS256")
)

type Header struct {
	Alg Algorithm        `json:"alg"`
	Enc EncryptionMethod `json:"enc"`
	Zip string           `json:"zip,omitempty"`
	Jku string           `json:"jku,omitempty"`
	Jwk string           `json:"jwk,omitempty"`
	Kid string           `json:"kid,omitempty"`
	X5u string           `json:"x5u,omitempty"`
	X5c string           `json:"x5c,omitempty"`
	X5t string           `json:"x5t,omitempty"`
}

func NewHeader() Header {
	return Header{}
}

func (h *Header) SetHeader(alg Algorithm, enc EncryptionMethod) {
	h.Alg = alg
	h.Enc = enc
}
