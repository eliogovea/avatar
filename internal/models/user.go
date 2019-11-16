package models

type User struct {
  Username string `json:"username"`
  Avatar   string `json:"avatar"`
  IsAdmin  bool   `json:"is_admin"`
}
