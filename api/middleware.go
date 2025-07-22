package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yogesh-k64/middleware-simple-bank/token"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationType       = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorHandler(err))
			return
		}

		fields := strings.Fields(authorizationHeader)

		if len(fields) < 2 {
			err := errors.New("authorization header is invalid")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorHandler(err))
			return
		}

		if authorizationType != strings.ToLower(fields[0]) {
			err := errors.New("authorization header is not supported")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorHandler(err))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorHandler(err))
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
