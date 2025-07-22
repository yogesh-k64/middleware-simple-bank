package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/yogesh-k64/middleware-simple-bank/db/sqlc"
	"github.com/yogesh-k64/middleware-simple-bank/token"
	"github.com/yogesh-k64/middleware-simple-bank/utils"
)

type Server struct {
	store      db.Store
	config     utils.Config
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(store db.Store, config utils.Config) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmentricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		store:      store,
		config:     config,
		tokenMaker: tokenMaker,
	}
	server.SetupRouter()
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorHandler(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) SetupRouter() {
	router := gin.Default()
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authRouter := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRouter.POST("/accounts", server.createAccount)
	authRouter.GET("/accounts/:id", server.getAccount)
	authRouter.GET("/accounts", server.listAccounts)

	authRouter.POST("/transfers", server.createTransfer)

	server.router = router
}
