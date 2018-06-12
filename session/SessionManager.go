package session

type Manager interface {
	GetInfo(string) (string, bool, error)
	IsIdActive(string) bool
	IsUsernameActive(string) bool
	AddSession(string, bool) (string, error)
	DeleteSession(string) error
	UpdateSession(string) error
}
