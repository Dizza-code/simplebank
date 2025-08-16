package api

import (
	db "github.com/Dizza-code/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

// This will create a new server instance and setup all api routes for our service on that server
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//add routes to the router
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
