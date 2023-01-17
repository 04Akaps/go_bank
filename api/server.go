package api

import (
	"fmt"
	db "simple_bank/db/sqlc"
	util "simple_bank/db/utils"
	"simple_bank/token"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config util.Config
	store *db.Store
	tokenMaker token.Maker
	router *gin.Engine
}

func NewServer(config util.Config, store *db.Store) (*Server, error){
	tokenMaker, err := token.NewPasetoMaker(config.TokenKey)

	if err != nil {
		return nil, fmt.Errorf("create token error %w", err)
	}

	server := &Server{config:  config,store: store, tokenMaker: tokenMaker}
	server.setUpRouter()

	return server, nil
}

func (server *Server) setUpRouter	 () {
	router := gin.Default()

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/account/:id", server.getAccount)
	authRoutes.GET("/account", server.listAccount)

	server.router = router
}

//start server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error" : err.Error()}
}