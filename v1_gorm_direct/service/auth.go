package service

import (
	"os"
	"time"
	"v1_gorm_direct/controller/service"
	"v1_gorm_direct/graph/model"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtKey = os.Getenv("API_KEY")

const (
	expRefreshToken = 24
	expToken        = 15
	company         = "Hybrid Technologies Viet Nam"
	hostMail        = "humghuy201280@gmail.com"
	subjectMail     = "Email reset password"
	textContent     = "Struction to reset your password:"
)

func NewAuthService() service.AuthService {
	return &AuthService{}
}

type AuthService struct{}

func (m *AuthService) GenerateToken(useInfo model.User) (string, error) {
	//generate access token
	tokenJwt := jwt.New(jwt.SigningMethodHS256)
	claims := tokenJwt.Claims.(jwt.MapClaims)
	claims["id"] = useInfo.ID
	claims["exp"] = time.Now().Add(time.Hour * expToken).Unix()
	token, err := tokenJwt.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	//generate refresh token
	// refreshTokenJWT := jwt.New(jwt.SigningMethodHS256)
	// rtClaims := refreshTokenJWT.Claims.(jwt.MapClaims)
	// rtClaims["id"] = useInfo.Id
	// rtClaims["exp"] = time.Now().Add(time.Hour * expRefreshToken).Unix()

	// refreshToken, err := refreshTokenJWT.SignedString([]byte(jwtKey))
	// if err != nil {
	// 	return nil, err
	// }

	return token, nil
}

// func (m *AuthService) ValidateToken(ctx *gin.Context, tokenString string) error {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}

// 		return []byte(jwtKey), nil
// 	})

// 	if err != nil {
// 		return err
// 	}

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		userID := uint(claims["id"].(float64))
// 		ctx.Set("userID", int(userID))
// 		return nil
// 	} else {
// 		return fmt.Errorf("account invalid")
// 	}
// }
