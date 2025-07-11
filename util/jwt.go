package util

import "github.com/golang-jwt/jwt/v5"

func GenerateJWT(userID int, userEmail string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   userEmail,
	}
	jwt, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("rahasia"))
	if err != nil {
		return "", err
	}
	return jwt, nil
}

func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("rahasia"), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}
