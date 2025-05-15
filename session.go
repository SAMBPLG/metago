package metago

import (
	"context"
	"fmt"
)

// session represents Metabase's Session API
type session struct {
	sdk *Metago
}

// Login get session from metabase using username and password
func (m *session) Login(ctx context.Context, usr, pass string) (*string, error) {
	var errResp MetabaseErr
	var authResp AuthResponse
	resp, err := m.sdk.getRequest(ctx).
		SetError(errResp).
		SetBody(map[string]string{
			"username": usr,
			"password": pass,
		}).
		SetResult(&authResp).
		Post(m.sdk.endpoint.SessionEndpoint)
	if err := checkForError(resp, err, "failed to get session"); err != nil {
		return nil, err
	}
	return authResp.SessionID, nil
}

// Logout request delete metabase session
func (m *session) Logout(ctx context.Context) error {
	var errResp MetabaseErr
	var logoutResp any
	resp, err := m.sdk.getRequestWithAuth(ctx).
		SetError(&errResp).
		SetResult(&logoutResp).
		Delete(m.sdk.endpoint.SessionEndpoint)
	if err := checkForError(resp, err, "failed to delete session"); err != nil {
		return err
	}
	return nil
}

func (m *session) ResetPassword(ctx context.Context) error {
	return fmt.Errorf("not implemented")
}
