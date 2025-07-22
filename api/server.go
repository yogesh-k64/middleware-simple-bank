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

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	router.POST("/transfers", server.createTransfer)

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	server.router = router
}
