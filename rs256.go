package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

var publicKey = `-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAvirSgZYgExeWREezmBvZ
9PUxrmwDS5DACGhmXNgKpjQRenGC5vuDDXCC97rMHRvFCZZnyXAVxawYpN9wNVaF
sR4YmJHjRHeZrFpZHEqb6YteL3biNgyV6Uz6naflubXOa7goX6CNZ0a+MK6zD66L
3T9VcOSzLuxr/o+eaW84C58uDdLDY0RJXtD0nK5N0/I3l9C0GgrY6rlYs0POtJP1
FlxBnr7XhHi/xyUdXYrVpe720hz10u5LDw4Se4hA/P+0q5n7GJ43VBAMPq+tU74a
5HPtm8Cf2EgkAfgOGUbVs7hcy+mX0a9ZjDbDiFjBqELsZtNKJgWeeVsD8FOhvqbg
pDCXanEFiPkzm0+R2rrbrTHQC2fG3kv9y+H3gVdimaIQLVFBHj6HUyQeR73ahbuV
fWLpDx/pJ99EFymwxO2cBvlT9LqTOOCuePJxbsXJGTq30jGZswghbFx/NY1GIVSR
9BNvyOVjy0+AbriJMiVk0WDfHXDxd8kCxUJoMdwYa39JIuoIlVggesXdPijBXGT4
HRyC/oW2IrfbJwcYO12SbmbIix6aGVMcVlKvP0TW3tSmn1K7GRsI+y8WnLzPeboc
tsbMTp8ya8wmggGtbsBZQT2FhGiI0f4XGre3Y6NDoc0IZFuq7RCLnMVsoY66zICu
3pOseSoebApaotc7lhchtlECAwEAAQ==
-----END PUBLIC KEY-----`

var privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIJKQIBAAKCAgEAvirSgZYgExeWREezmBvZ9PUxrmwDS5DACGhmXNgKpjQRenGC
5vuDDXCC97rMHRvFCZZnyXAVxawYpN9wNVaFsR4YmJHjRHeZrFpZHEqb6YteL3bi
NgyV6Uz6naflubXOa7goX6CNZ0a+MK6zD66L3T9VcOSzLuxr/o+eaW84C58uDdLD
Y0RJXtD0nK5N0/I3l9C0GgrY6rlYs0POtJP1FlxBnr7XhHi/xyUdXYrVpe720hz1
0u5LDw4Se4hA/P+0q5n7GJ43VBAMPq+tU74a5HPtm8Cf2EgkAfgOGUbVs7hcy+mX
0a9ZjDbDiFjBqELsZtNKJgWeeVsD8FOhvqbgpDCXanEFiPkzm0+R2rrbrTHQC2fG
3kv9y+H3gVdimaIQLVFBHj6HUyQeR73ahbuVfWLpDx/pJ99EFymwxO2cBvlT9LqT
OOCuePJxbsXJGTq30jGZswghbFx/NY1GIVSR9BNvyOVjy0+AbriJMiVk0WDfHXDx
d8kCxUJoMdwYa39JIuoIlVggesXdPijBXGT4HRyC/oW2IrfbJwcYO12SbmbIix6a
GVMcVlKvP0TW3tSmn1K7GRsI+y8WnLzPeboctsbMTp8ya8wmggGtbsBZQT2FhGiI
0f4XGre3Y6NDoc0IZFuq7RCLnMVsoY66zICu3pOseSoebApaotc7lhchtlECAwEA
AQKCAgBIl14d22nI//L7g4dZ5B/SMxrQ4yhq2wmC7B9PB6UhBrU5UUVP2OiQ48cK
u8KYxfX0D/b0XRXijCwoG6bgpXOJRdzLuLzRcPo3YZGsjApyHyJH6hC14x4CncD5
F5NIzc7LLdQzlL0FlGqbeMSfktiPD1MVMif4HIWo+bfVtX/mZ9ATtMWjCfkb+ZW2
jY7l/gscp61oV4WwLCpg29x28BfZfkQKTf1E8zb51QAdqhaeLdcII8fuycnbKFsD
zuJH8XBNJQ1Fu3eRXkLeVv9J/UKUvHASSXh6/ibInaD8Ix8GaLT7neK3R0FelbUU
QhVCnrFRGwyt8O0qlASv46zVcyxCuCzs7qeb0WfOGvC/e5BE/sHZ9pYWVTZ6OlVE
e5k44NNy/aaz8Z/QYKhry7KAb/4ZKOu7/dBkibR8Hait+wX58amF8qRZ7pKqmQEX
D7XoEVJ3xwME9UpMMAxG3EsQdoM942VrZYk2RHvG6/qeaHi6R136ypBwGYQ0gbGI
LAsWFsZP0ej+WcgG5+j2z3KoTAxBoyUwvs4MYBabWA5MephSJj56zS6pGY629H4W
JyMpSzwsdugUSfd4W86Uz9587c6ql8oSaTcLmg6o+2OQzCc194Z/kk30aVxU79za
9x2Mre2rN0YggIorK8IVYtk8fr74P/pVLvSd2dnj9SGNDmBY8QKCAQEA5wExL+DH
wZ7OGXaLNk1+IscHd/fYZ71aQeTsUtaj4VMEvci4hOUfzVccOpYXBFKBI0BkO4zD
ws4oocThdfJKzsgOWJfa4dWhJGbPOte+mhNmMa5zNPoygFRZKDcCSS0bSwcf0nWx
zV81b5byrWWNtfM+eAP9AOeMnc1aCbPc1lfd07Kds0nLwmVtZRzQ6IoK9ctIurgx
OYT3Haoz42l/jU/a+5s9nWUNGCaQyZfn7qx0jMjDvzrzxClVdNBDSe6+sF0qLmBR
tiB1AnwAFPWy6HcXRtORPvhfaKSERoVjkzmWoZvvreCdIOL1BnqQE9CgR1T2QHoU
pIa/AZKx1/Wh/QKCAQEA0r5wOwykECMBRPoPI3eDDR+KiabWtDhDXReeX0QBryJ8
6m8PCaXyIRO6nEN6vBDSSGODd//RwdjLhbFWuagGxM4Ho3CKM/iLD3brN9cwvcfO
hLxUxMiOJWPr2d64sq0uGXOM7vxDoXFDMWo8kmWgxFPObiJzA+e1OJg+x+QcW7hG
po/stQNf4dAZSrQxnL2AxCzZT9bkZ29wGM/8EYTUDbmi/Z6nj11Km6UiD7QmGWyV
9qkakb7heIG5DC6Fpi73xAfozLLSRP2ZDlUXsQb3WK8vtwiNgvkM8jNlFoTL/+0a
KHBZZ4L1rS6ebooF4iJQRzgn2yoMXUPpvSV6GaS75QKCAQBsPU7C1FxFCRghLj6E
r0LPINsLB4LK+Rp1jcJ9/jzDs5ahJo/vFT0Vhh+gA/u89ruzvNQ6YvbHLLo1Mn9C
OMl8opi2QlE5SFQ9Lj8jnzucDkLwPIGW7TUElIFCKqRTjWMggLIUdzlctaPgKaaE
QPVpsBQxQA0og1aMClCKA1ESzhEOxL7H52gmKkhb/GiWzTfde+cUNoI6JWd06u9t
O12c5TICevcf7N1513g4PYlbeEsUPG4cCI939rYoCf29grSvBrhhCpi+8e5hv6B6
MbYm3sZ1VWTVUKPD8HJaaN4DehRunRZtLrXxVubgpkUkQ7kWM2U/SPSrMsgFAKt8
OYgxAoIBAQCZ0GcnohVItdPiXUSZVUwwUtoRl9TTGtGH64oq6/7yrZBQpxhpqVXs
8HJeRR4aj4h7Ty6QcgXNnwcEoCe1P78Y+2s1zhkAz6Hneps5WXV4vpGr7a1NJzgg
cCqfDODvgFjKOL0fXL2b4ofxVCG7lDNft+9OERSzP/XTxcLksEhGZVwji55vi3P6
46DiFlyzktid1kIR9L0mBX7ijULkIneHQPuGcrrHd4bRzRfsMEcyfT+DFW+P+qqk
AsJl8rqXTWaHDGLMrKOtuQ5yGIc/LN9xOgPwamymsFHToNFiOzCNradO5plKZJod
eY4nDdQ3dWX0ZTcpzinSFJRP+j+A4exFAoIBAQDPlS40epiI6+VIOcGUr7OiYwR6
RzP/azqE3DFdNzy4K96V35CUt2sn0uyVJxwCRN76344wDbmwSWuMZxqt6I8D/qrH
XV3gv7sCF0z44ORrTw+BibvftenS8dudTh1enPH3ORcSmRLZAoRi9Dfhbx9yCJue
R1kJJHsUreIJV3kTj7vrMwiy/lwchSlJQOw6WYQHmIOhUJJl4rFfS8V6u95fHw6V
i/tTAFnD9Qo9XtJltzsANvg+YCGfGQH7Q5wl8fdf8ks3M3QG5O8NbQiOjhr8nWQ4
USogEWLZiXJM+7ZHQPZbjaOr/Z13Of1tUEr/t65QrNsSGiYDcuHOl5sp+6wh
-----END RSA PRIVATE KEY-----`

func rs256Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "jon" && password == "shhh!" {
		// Create token
		token := jwt.New(jwt.SigningMethodRS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Jon Snow"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

		// Generate encoded token and send it as response.
		priKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
		if err != nil {
			return err
		}
		t, err := token.SignedString(priKey)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}

func rs256Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func rs256(e *echo.Echo) *echo.Echo {
	e.Post("/rs/login", rs256Login)

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		log.Fatal(err)
	}

	// Restricted group
	r := e.Group("/rs256")
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Skipper:       func(echo.Context) bool { return false },
		SigningMethod: "RS256",
		ContextKey:    "user",
		TokenLookup:   "header:" + echo.HeaderAuthorization,
		SigningKey:    pubKey,
	}))
	r.GET("/restricted", rs256Restricted)

	return e
}
