package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"

	"dac.sg/middleware"
	"dac.sg/server"
)

func TestLang(t *testing.T) {
	app := server.Server()

	tests := []struct {
		name           string
		cookieValue    string
		headerValue    string
		expectedResult string
	}{
		{
			name:           "Supported language cookie",
			cookieValue:    "en",
			headerValue:    "",
			expectedResult: "en",
		},
		{
			name:           "Unsupported language cookie",
			cookieValue:    "es",
			headerValue:    "",
			expectedResult: "es",
		},
		{
			name:           "Supported language in header",
			cookieValue:    "",
			headerValue:    "es",
			expectedResult: "es",
		},
		{
			name:           "Unsupported language in header",
			cookieValue:    "",
			headerValue:    "en",
			expectedResult: "en",
		},
		{
			name:           "No header and no cookie",
			cookieValue:    "",
			headerValue:    "",
			expectedResult: "en",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new fasthttp request and response
			req := fasthttp.AcquireRequest()
			defer fasthttp.ReleaseRequest(req)

			resp := fasthttp.AcquireResponse()
			defer fasthttp.ReleaseResponse(resp)

			if tt.cookieValue != "" {
				cookie := fasthttp.Cookie{}
				cookie.SetKey("language")
				cookie.SetValue(tt.cookieValue)
				req.Header.SetCookie(string(cookie.Key()), string(cookie.Value()))
			}
			if tt.headerValue != "" {
				req.Header.Set("Accept-Language", tt.headerValue)
			}

			// Create a new fasthttp request context
			ctx := &fasthttp.RequestCtx{}
			ctx.Init(req, nil, nil)

			c := app.AcquireCtx(ctx)
			defer app.ReleaseCtx(c)

			lang := middleware.Lang(c)
			assert.Equal(t, tt.expectedResult, lang)
		})
	}
}
