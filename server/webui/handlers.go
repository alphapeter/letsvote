package webui

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsHandler(c *gin.Context) {
	c.Data(http.StatusOK, "application/javascript", Javascript)
}

func CssHandler(c *gin.Context) {
	c.Data(http.StatusOK, "text/css", Css)
}

func HtmlHandler(c *gin.Context) {
	c.Data(http.StatusOK, "text/html", Html)
}
