package polls

import "github.com/gin-gonic/gin"

type fetchError interface {
	message() string
	responseCode() int
}

type unsuccessfulFetch struct {
	msg  string
	code int
}

func (f unsuccessfulFetch) message() string {
	return f.msg
}
func (f unsuccessfulFetch) responseCode() int {
	return f.code
}

func errorResponse(c *gin.Context, reason string, responseCode int) {
	c.JSON(responseCode, struct {
		Success bool   `json:"success"`
		Reason  string `json:"reason"`
	}{false, reason})
}
