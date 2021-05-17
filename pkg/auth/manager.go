package auth

import (
	"fmt"
	"golang-core/config"
	"math/rand"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)


type TokenManager interface {
}

type Manager struct {

}
var jwtkey =[]byte (os.Getenv("JWT_SECRET_KEY"))

func (m *Manager) NewJWT(userId string) (token string, err error) {

	expiration :=  time.Now().Add(config.TTL *time.Minute).Unix()

	tokenData := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.StandardClaims{
		ExpiresAt:expiration,
		Subject: userId,
	})

	token, err = tokenData.SignedString(jwtkey)
	return
}

func (m *Manager) Parse(accessToken string)(subject string, err error){

	token, errToken := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtkey), nil
	})

	if errToken != nil {
		subject = ""
		err = errToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	
	if !ok {
		subject = ""
		err = fmt.Errorf("error get user claims from token")
	}
	subject = claims["sub"].(string)
	return
}

func(m *Manager)NewRefreshToken()(string, error){
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	_, err := r.Read(b)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x",b),nil
}