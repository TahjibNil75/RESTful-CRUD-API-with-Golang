package utils

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

var JwtToken = []byte("JWT_SECRET")

type JwtClaim struct {
	Email string `json:"email"`
	Uid   uint   `json:"uid"`
	jwt.StandardClaims
}

func CreateAccessToken(email string, uid uint) (accessToken string, refreshToken string, err error) {

	accessTokenExpirationTime := time.Now().Add(6 * time.Hour)
	refreshTokenExpirationTime := time.Now().Add(168 * time.Hour)

	claims := &JwtClaim{
		Email: email,
		Uid:   uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessTokenExpirationTime.Unix(),
		},
	}

	refreshClaims := &JwtClaim{
		Email: email,
		Uid:   uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshTokenExpirationTime.Unix(),
		},
	}

	// Create a new JWT token with the specified claims and signing method (HS256)
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(JwtToken))
	if err != nil {
		log.Panic(err)
		return "", "", err
	}

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(JwtToken))
	if err != nil {
		log.Panic(err)
		//return "", err 			// not enough return values; have (string, error);want (string, string, error)
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func UpdateAllTokens(accessToken string, refreshToken bool) (string, error) {
	// Initialize a claims struct to store token information
	claims := &JwtClaim{}

	// Parse the existing token without verifying the signature
	_, _, err := new(jwt.Parser).ParseUnverified(accessToken, claims)
	if err != nil {
		return "", err
	}

	// Determine the expiration time based on whether it's a refresh token or not
	var expirationTime time.Duration

	if refreshToken {
		expirationTime = time.Hour * 168
	} else {
		expirationTime = time.Hour * 24
	}

	claims.ExpiresAt = time.Now().Local().Add(expirationTime).Unix()

	// Create a new token with the updated expiration time
	newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedAcessToken, err := newAccessToken.SignedString([]byte(JwtToken))
	if err != nil {
		return "", err
	}
	return signedAcessToken, nil

}

var Email string

func ValidateToken(signedToken string) (UserId string, err error) {
	// Parse JWT with claims
	token, parseErr := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(JwtToken), nil
		},
	)
	if parseErr != nil {
		err = parseErr
		return
	}
	// Set email and ID values from JWT claims
	claims, ok := token.Claims.(*JwtClaim)
	UserId = strconv.Itoa(int(claims.Uid))
	Email = claims.Email
	if !ok {
		err = errors.New("failed to parse JWT claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return UserId, nil
}
