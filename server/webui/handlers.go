package webui

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsHandler(c *gin.Context) {
	c.Data(http.StatusOK, "gui.Javascript", Javascript)
}

func HtmlHandler(c *gin.Context) {
	c.Data(http.StatusOK, "text/html", Html)
}
