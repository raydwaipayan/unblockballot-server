package types

// User user object
type User struct {
	Email     string `json:"email" form:"email"`
	FirstName string `json:"firstname" form:"firstname"`
	LastName  string `json:"lastname" form:"lastname"`
	Password  string `json:"password" form:"password"`
}
