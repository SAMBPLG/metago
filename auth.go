package metago

import "context"

type AuthMethod uint

const (
	APIKEY   AuthMethod = iota // auth header: x-api-key
	USERPASS                   // Auth Header: // x-metabase-session
)

// getSession get session from metabase using the desired auth method
func (m *Metago) getSession(ctx context.Context, key, usr, pass string) error {
	switch m.authMethod {
	case APIKEY:
		m.session = &key
	case USERPASS:
		ses, err := m.Session.Login(ctx, usr, pass)
		if err != nil {
			return err
		}
		m.session = ses
	}
	return nil
}

// UserPermission represents common fields for user permission
type UserPermission struct {
	CanWrite   bool `json:"can_write,omitempty"`
	CanDelete  bool `json:"can_delete,omitempty"`
	CanUpdate  bool `json:"can_update,omitempty"`
	CanRestore bool `json:"can_restore,omitempty"`
}
