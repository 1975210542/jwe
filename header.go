package main

//type Algorithm string

const (
	ALG_RSA1_5       = "RSA1_5"
	ALG_RSA_OAEP     = "RSA-OAEP"
	ALG_RSA_OAEP_256 = "RSA-OAEP-256"
	ALG_A128KW       = "A128KW"
	ALG_A256KW       = "A256KW"
)

//type EncryptionMethod1 string

const (
	ENC_A128CBC_HS256_v7 = "A128CBC+HS256"
	ENC_A256CBC_HS512_v7 = "A256CBC+H512"
	ENC_A128CBC_HS256    = "A128CBC-HS256"
	ENC_A256CBC_HS512    = "A256CBC-HS512"

	ENC_A128GCM = "A128GCM"
	ENC_A256GCM = "A256GCM"
)

type Header struct {
	Alg string `json:"alg"`
	Enc string `json:"enc"`
	Zip string `json:"zip,omitempty"`
	Jku string `json:"jku,omitempty"`
	Jwk string `json:"jwk,omitempty"`
	Kid string `json:"kid,omitempty"`
	X5u string `json:"x5u,omitempty"`
	X5c string `json:"x5c,omitempty"`
	X5t string `json:"x5t,omitempty"`
}

func NewHeader(alg, enc string) Header {
	header := Header{}
	header.Alg = alg
	header.Enc = enc
	return header
}

func (h *Header) SetHeader(alg, enc string) {
	h.Alg = alg
	h.Enc = enc
}
