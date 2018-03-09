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

// admin section
func AdminJsHandler(c *gin.Context) {
	c.Data(http.StatusOK, "application/javascript", AdminJavascript)
}

func AdminCssHandler(c *gin.Context) {
	c.Data(http.StatusOK, "text/css", AdminCss)
}

func AdminHtmlHandler(c *gin.Context) {
	c.Data(http.StatusOK, "text/html", AdminHtml)
}
