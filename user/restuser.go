package user

type UserREST struct {
  UserID string `json:"userid"`
  FirstName string `json:"firstname"`
  LastName string `json:"lastname"`
  Email string `json:"email"`
  Password string `json:"password"`
}
