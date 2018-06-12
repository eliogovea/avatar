package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/eliogovea/avatar/auth"
	"github.com/eliogovea/avatar/session"
)

type server struct {
	router           *http.ServeMux
	AuthMan          *auth.Auth `json:"auth_manager"`
	sessions         session.Manager
	Fs               *storage      `json:"storage"`
	Address          string        `json:"address"`
	LoginPath        string        `json:"login_path"`        // /login
	LogoutPath       string        `json:"logout_path"`       // /logout
	PersonalPath     string        `json:"personal_path"`     // /personal
	UploadPath       string        `json:upload_path`         // /upload
	LoginCookieName  string        `json:"login_cookie_name"` // login-avatar
	SessionDuration  time.Duration `json:"session_duration"`  // time.Hour
	LoginTemplate    string        `json:"login_template"`    // ./templates/login.html
	PersonalTemplate string        `json:"personal_template"` // ./templates/personal.html
	StaticFiles      string        `json:"static_files"`      // ./static/
	MaxUploadSize    int           `json:"max_upload_size"`   // 2MB
}

func saveConfig(s *server, w io.Writer) error {
	// return json.NewEncoder(w).Encode(s)
	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")
	return enc.Encode(s)
}

func loadFromConfig(configPath string) (*server, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	s := new(server)
	s.AuthMan = auth.New()
	s.router = http.NewServeMux()
	s.sessions = session.NewManager()
	s.Fs = new(storage)

	// json.NewDecoder(file).Deconde(s)

	dec := json.NewDecoder(file)
	dec.Decode(s)

	// s.buildHandlers()

	return s, nil
}

func (s *server) buildHandlers() error {
	s.router.HandleFunc("/", s.rootHandler())
	s.router.HandleFunc(s.LoginPath, s.notLoggedOnly(s.loginHandler(s.LoginTemplate)))
	s.router.HandleFunc(s.LogoutPath, s.loggedOnly(s.logoutHandler()))
	s.router.HandleFunc(s.PersonalPath, s.loggedOnly(s.personalHandler(s.PersonalTemplate)))
	s.router.HandleFunc("/api/approved/", s.getApprovedAvatar())
	s.router.HandleFunc("/api/pending/", s.loggedOnly(s.getPendingAvatar()))
	s.router.HandleFunc("/upload", s.loggedOnly(s.uploadHandler()))
	s.router.Handle(s.StaticFiles, http.StripPrefix(s.StaticFiles, http.FileServer(http.Dir("static"))))
	return nil
}
