package models

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(map[string]interface{}{"username": username, "password": password}).First(&auth)
	if auth.ID > 0 {
		return true
	}

	return false
}
