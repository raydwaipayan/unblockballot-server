package types

// User user object
type User struct {
	FirstName string `json:"firstname" form:"firstname"`
	LastName  string `json:"lastname" form:"lastname"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Role      int32  `json:"role" form:"role"`
}

// Organization  body
type Organization struct {
	OrgName string `json:"orgName" form:"orgName"`
	OrgImg  []byte `json:"orgImg" form:"orgImg"`
}

// PollBody poll object
type PollBody struct {
	Questions string   `json:"questions" form:"questions"`
	Options   []string `json:"options" form:"options"`
	OpensAt   string   `json:"opensAt" form:"opensAt"`
	ClosesAt  string   `json:"closesAt" form:"closesAt"`
	OrgName   string   `json:"orgName" form:"orgName"`
	OrgImg    []byte   `json:"orgImg" form:"orgImg"`
}
