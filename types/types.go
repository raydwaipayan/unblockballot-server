package types

// User user object
type User struct {
	Email     string `json:"email" form:"email"`
	FirstName string `json:"firstname" form:"firstname"`
	LastName  string `json:"lastname" form:"lastname"`
	Admin     int32  `json:"admin" form:"admin"`
	Password  string `json:"password" form:"password"`
}

// Poll poll object
type Poll struct {
	Questions string   `json:"questions" form:"questions"`
	Options   []string `json:"options" form:"options"`
	OpensAt   string   `json:"opensAt" form:"opensAt"`
	ClosesAt  string   `json:"closesAt" form:"closesAt"`
	OrgName   string   `json:"orgName" form:"orgName"`
	OrgImg    []byte   `json:"orgImg" form:"orgImg"`
}

//Election contains election details
type Election struct {
	ID         string `json:"id"`
	Candidates []string
	Active     bool
}

//Candidate contains the candidature information
type Candidate struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	ElectionID string
	Votes      uint64
}

//Vote structure for storing votes
type Vote struct {
	VoterID     string `json:"id"`
	CandidateID string `json:"candidateID"`
	Time        string `json:"time"`
}
