package metago

import (
	"context"

	"github.com/go-resty/resty/v2"
)

// getRequest returns *resty.Request with context
func (m *Metago) getRequest(ctx context.Context) *resty.Request {
	return m.restyClient.R().SetContext(ctx)
}

// getRequestWithAuth returns *resty.Request with auth header set. Will panic when session is nil.
func (m *Metago) getRequestWithAuth(ctx context.Context) *resty.Request {
	if m.session == nil {
		panic("session is empty")
	}
	r := m.getRequest(ctx)
	switch m.authMethod {
	case USERPASS:

		r = r.SetHeader("x-metabase-session", *m.session)
	case APIKEY:
		r = r.SetHeader("x-api-key", *m.session)
	}
	return r
}
