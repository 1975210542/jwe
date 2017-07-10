package main

type Algorithm string

const (
	ALG_RSA1_5       = Algorithm("RSA1_5")
	ALG_RSA_OAEP     = Algorithm("RSA-OAEP")
	ALG_RSA_OAEP_256 = Algorithm("RSA-OAEP-256")
	ALG_A128KW       = Algorithm("A128KW")
	ALG_A256KW       = Algorithm("A256KW")
)

type EncryptionMethod string

const (
	ENC_A128CBC_HS256_v7 = EncryptionMethod("A128CBC+HS256")
	ENC_A256CBC_HS512_v7 = EncryptionMethod("A256CBC+H512")
	ENC_A128CBC_HS256    = EncryptionMethod("A128CBC-HS256")
	ENC_A256CBC_HS512    = EncryptionMethod("A256CBC-HS512")
	ENC_A128GCM          = EncryptionMethod("A128GCM")
	ENC_A256GCM          = EncryptionMethod("A256GCM")
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
