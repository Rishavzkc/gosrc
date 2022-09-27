package Routes

import (
	"gingorm/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/newdb")
	{
		grp1.GET("book", Controllers.GetBooks)
		grp1.POST("book", Controllers.CreateBook)
		grp1.GET("book/:id", Controllers.GetBookByID)
		grp1.PUT("book/:id", Controllers.UpdateBook)
		grp1.DELETE("book/:id", Controllers.DeleteBook)
	}
	return r
}
