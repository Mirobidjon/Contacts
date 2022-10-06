package contact

// Contact struct
type Contact struct {
	ID     int    `json:"id" db:"id"`
	Name   string `json:"name" binding:"required" db:"name"`
	Phone  string `json:"phone" binding:"required" db:"phone"`
	UserID int    `json:"user_id" db:"user_id"`
}

// DefaultContact ...
type DefaultContact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
