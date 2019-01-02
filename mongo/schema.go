package mongo

// Team contains information about a team
type Team struct {
	ID        uint32
	Name      string `json:"name"`
	Record    string `json:"record"`
	HeadCoach string `json:"headCoach"`
}
