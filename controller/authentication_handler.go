package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("TutupLapak")
var tokenName = "token"

type Claims struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func generateToken(c *gin.Context, id int, name string, email string) {
	tokenExpiryTime := time.Now().Add(60 * time.Minute)

	claims := &Claims{
		ID:    id,
		Name:  name,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpiryTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return
	}

	// c.SetCookie(tokenName, signedToken, tokenExpiryTime, "/", "localhost", false, true)
	c.SetCookie(tokenName, signedToken, 1000, "/", "localhost", false, true)
}

func resetUserToken(c *gin.Context) {
	c.SetCookie(tokenName, "", -1, "/", "localhost", false, true)
}

// func Authenticate(accessType int) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		isValidToken := validateUserToken(c, accessType)
// 		if !isValidToken {
// 			var response model.UserResponse
// 			response.Message = "Unauthorized Access"
// 			fmt.Println("Unauthorized Access")
// 			// sendUnAuthorizedResponse(c, response)
// 			c.Abort()
// 			return
// 		} else {
// 			c.Next()
// 		}
// 	}
// }

// func validateUserToken(c *gin.Context, accessType int) bool {
// 	isAccessTokenValid, id, name, email := validateTokenFromCookies(c)
// 	fmt.Print(id, email, userType, accessType, isAccessTokenValid)

// 	if isAccessTokenValid {
// 		isUserValid := name == accessType
// 		fmt.Print(isUserValid)
// 		if isUserValid {
// 			return true
// 		}
// 	}
// 	return false
// }

func validateTokenFromCookies(c *gin.Context) (bool, int, string, string) {
	if cookie, err := c.Cookie(tokenName); err == nil {
		accessToken := cookie
		accessClaims := &Claims{}
		parsedToken, err := jwt.ParseWithClaims(accessToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err == nil && parsedToken.Valid {
			return true, accessClaims.ID, accessClaims.Name, accessClaims.Email
		}
	}
	return false, -1, "", ""
}
