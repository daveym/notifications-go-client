package authorisation

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// IRequestTokenCreator - Creates a authorisation token that is appended to request as header
type IRequestTokenCreator interface {
	Create(string, string, string, string) (string, error)
}

// JWTReqestTokenCreator - JWT implementation of IRequestTokenCreator
type JWTReqestTokenCreator struct {
	_claimKeyRequestSignature string
	_claimKeyPayloadSignature string
}

// Create - Creates a JWT token. JwtTokenCreator uses HMAC-SHA256 signature, by default.
func (j *JWTReqestTokenCreator) Create(resourcePath string, body string, issuer string, secret string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims["ISSUER"] = issuer
	token.Claims["ISSUED_AT"] = time.Now()

	tokenString, err := token.SignedString(j._claimKeyRequestSignature)

	return tokenString, err

}
