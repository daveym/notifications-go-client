package authorisation

// IRequestTokenCreator - Creates a authorisation token that is appended to request as header
type IRequestTokenCreator interface {
	Create(string, string, string, string) string
}

// JWTReqestTokenCreator - JWT implementation of IRequestTokenCreator
type JWTReqestTokenCreator struct {
	_claimKeyRequestSignature string
	_claimKeyPayloadSignature string
}
