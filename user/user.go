package user

// User struct
type User struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

// PrivateUserDetails struct
type PrivateUserDetails struct {
	ID       int64
	Password string
	Salt     string
}
