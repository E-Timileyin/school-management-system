package routes

import (
	"github.com/gin-gonic/gin"
	"school-management-backend/internal/handler"
)

// setupLibraryRoutes configures all library related routes
func setupLibraryRoutes(router *gin.RouterGroup, libraryHandler *handler.LibraryHandler) {
	// Library routes group
	library := router.Group("/library")
	{
		// Book routes
		books := library.Group("/books")
		{
			books.POST("", libraryHandler.CreateBook)
			books.GET("/:id", libraryHandler.GetBook)
		}

		// Library card routes
		cards := library.Group("/cards")
		{
			cards.POST("", libraryHandler.IssueLibraryCard)
		}

		// Book circulation routes
		circulation := library.Group("/circulation")
		{
			circulation.POST("/checkout", libraryHandler.CheckoutBook)
			circulation.PUT("/return", libraryHandler.ReturnBook)
		}

		// Fine routes
		fines := library.Group("/fines")
		{
			fines.POST("/pay", libraryHandler.PayFine)
		}
	}
}
