package gincontext

import (
	"context"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ginKey struct{}

type GinData struct {
	Form        *multipart.Form
	Header      http.Header
	Accepted    []string
	ContentType string
	ClientIP    string
	RemoteIP    string
	UserAgent   string
}

// NewContext put gin data into context
func NewContext(ctx *gin.Context) context.Context {
	data := &GinData{
		Form:        ctx.Request.MultipartForm,
		Header:      ctx.Request.Header,
		Accepted:    ctx.Accepted,
		ContentType: ctx.ContentType(),
		ClientIP:    ctx.ClientIP(),
		RemoteIP:    ctx.RemoteIP(),
		UserAgent:   ctx.Request.UserAgent(),
	}
	return context.WithValue(ctx, ginKey{}, data)
}

// FromContext extract gin data from context
func FromContext(ctx context.Context) (data *GinData, ok bool) {
	data, ok = ctx.Value(ginKey{}).(*GinData)
	return
}
