package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"

	"github.com/colinaaa/hackweek/database"

	"golang.org/x/crypto/bcrypt"
)

type response struct {
	Msg        string `json:"msg"`
	StatueCode int    `json:"status_code"`
}

func vertifyPasswd(passwd, phoneNum string) bool {
	hash := database.SelectByPhone(phoneNum).PasswdHash
	if hash == "" {
		log.Println("hash empty")
		return false
	}
	err := bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(passwd),
	)
	if err != nil {
		log.Printf("passwd wrong: phoneNum: %s", phoneNum)
		return false
	}
	return true
}
func generatePasswd(passwd string) (hash []byte) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(passwd),
		bcrypt.DefaultCost,
	)
	if err != nil {
		log.Println(err)
		return nil
	}
	return
}

func generateToken(phoneNum string) (string, error) {
	fmt.Println("generateToken: get phoneNum:", phoneNum)
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"phoneNum": phoneNum,
		},
	)
	return token.SignedString([]byte("eat something"))
}
