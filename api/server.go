package api

import (
	db "github.com/anggavb/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	{
		routeAccounts := router.Group("/accounts")

		routeAccounts.POST("", server.createAccount)
		routeAccounts.GET("/:id", server.getAccount)
		routeAccounts.GET("", server.listAccounts)
	}

	server.router = router

	return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
