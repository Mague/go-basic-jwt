package main

import (
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
)

func main() {
	config := iris.Configuration{IsDevelopment: true, Gzip: true}
	var (
		server    = iris.New(config)
		middleJwt = jwtmiddleware.New(jwtmiddleware.Config{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				return []byte("this is secret key XD"), nil
			},
			SigningMethod: jwt.SigningMethodHS256,
		})
	)

	server.Post("/auth/log-in", login)
	server.Get("/zone/secured", middleJwt.Serve, securedZone)
	server.Listen(":3000")

}

func generateToken(id, user string) string {
	payload := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   id,
		"user": user,
	})
	payloadString, _ := payload.SignedString([]byte("this is secret key XD"))
	return payloadString
}

type Response struct {
	Text string `json:text`
}

func login(ctx *iris.Context) {
	user := ctx.FormValueString("user")
	pass := ctx.FormValueString("pass")

	if (user == "superman") && (pass == "criptonita") {

		response := Response{generateToken("1", user)}
		ctx.JSON(iris.StatusOK, iris.Map{
			"token": response,
		})
	} else {
		response := Response{"Username or password invalid"}
		ctx.JSON(iris.StatusOK, response)
	}
}

func securedZone(ctx *iris.Context) {
	ctx.Write("Authenticated")
}
