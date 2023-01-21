package cruds

import (
	_ "nokogiri-api/db"
	"nokogiri-api/utils"
	"time"

	"github.com/golang-jwt/jwt"
	_ "golang.org/x/crypto/bcrypt"
)

// func GenerateJWT(email string, password string) (jwtInfo db.JWTInfo, err error) {
// 	var (
// 		u     db.User
// 		token string
// 	)

// 	if err = db.Db.Where("email = ?", email).First(&u).Error; err != nil {
// 		return
// 	}

// 	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
// 	if err != nil {
// 		return
// 	}

// 	token, err = generateToken(u.ID)
// 	if err != nil {
// 		return
// 	}

// 	jwtInfo = db.JWTInfo{JWT: token}
// 	return
// }

func generateToken(userID string) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"iat": now.Unix(),
		"exp": now.Add(7 * 24 * time.Hour).Unix(),
	})
	return token.SignedString(utils.SigningKey)
}
