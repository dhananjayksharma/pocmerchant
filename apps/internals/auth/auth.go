package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("supersecretkey")
var jwtRefreshKey = []byte("supersecretkeyrefresh")

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

// getDuration
func getDuration(durationtype string, duration int) time.Time {
	var timetoexpire time.Duration
	timetoexpire = time.Duration(duration)
	expirationTime := time.Now().Add(timetoexpire * time.Millisecond)
	if durationtype == "minute" {
		expirationTime = time.Now().Add(timetoexpire * time.Minute)
	} else if durationtype == "hour" {
		expirationTime = time.Now().Add(timetoexpire * time.Hour)
	}

	return expirationTime
}

// getToken
func getToken(email string, username string, expirationTime time.Time, jwtKeyName []byte) (string, error) {
	claims := &JWTClaim{
		Email:    email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKeyName)
	return tokenString, err
}

// GenerateJWT
func GenerateJWT(email string, username string) (string, string, error) {
	timetoexpire := getDuration("minute", 2)
	token, err := getToken(email, username, timetoexpire, jwtKey) // token
	// fmt.Println("token:", token)

	timetoexpire = getDuration("hour", 1)
	refresh_token, err := getToken(email, username, timetoexpire, jwtRefreshKey) // refresh_token
	// fmt.Println("refresh_token:", refresh_token)
	return token, refresh_token, err
}

func ValidateRefreshToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtRefreshKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
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

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
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

func GetClaim(signedToken string) (*JWTClaim, error) {
	var claims *JWTClaim
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return claims, err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return claims, err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return claims, err
	}
	return claims, nil
}

func GetRefreshClaim(signedToken string) (*JWTClaim, error) {
	var claims *JWTClaim
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtRefreshKey), nil
		},
	)
	if err != nil {
		return claims, err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return claims, err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return claims, err
	}
	return claims, nil
}
