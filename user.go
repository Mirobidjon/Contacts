package contact

// User struct
type User struct {
	ID       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ID struct {
	ID int `json:"id"`
}

type Token struct {
	Token string `json:"token"`
}
