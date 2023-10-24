package utils

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

var AccessKey = []byte("superSecretKey")
var RefreshKey = []byte("refreshSecretKey")

type JWTClaim struct {
	Email string `json:"email"`
	Uid   uint   `json:"uid"`
	jwt.StandardClaims
}

func CreateAccessToken(email string, uid uint) (accessToken string, err error) {
	// Set expiration time for the refresh token (e.g., 6 hours)
	expirationTime := time.Now().Add(6 * time.Hour)
	claims := &JWTClaim{
		Email: email,
		Uid:   uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// Create a new JWT token with the specified claims and signing method (ES256).
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	// Sign the token using the AccessKey and get the token string.
	accessToken, err = token.SignedString(AccessKey)
	return
}

func CreateRefreshToken(uid uint) (refreshAccesstoken string, err error) {
	// Set expiration time for the refresh token (e.g., 30 days)
	expirationTime := time.Now().Add(30 * 24 * time.Hour)
	claims := &JWTClaim{
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	refreshAccesstoken, err = token.SignedString(RefreshKey)
	return

}

var Id string
var Email string

// ValidateToken function for validating a JWT
func ValidateToken(signedToken string) (err error) {
	// Parse JWT with claims
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(AccessKey), nil
		},
	)
	if err != nil {
		return
	}
	// Set email and ID values from JWT claims
	claims, ok := token.Claims.(*JWTClaim)
	Id = strconv.Itoa(int(claims.Uid))
	Email = claims.Email
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
