package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/richardktran/KMY-Drug-Store/pkg/app"
)

func Recovery() func(*gin.Context) {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				if err, oke := r.(error); oke {
					c.AbortWithStatusJSON(
						http.StatusInternalServerError,
						app.ResponseInternalServer(app.ThrowInternalServerError(err)),
					)
				}

				panic(r)
			}
		}()

		c.Next()
	}
}
