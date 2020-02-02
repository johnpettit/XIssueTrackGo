package user

type User struct {
  UserID string
  FirstName string
  LastName string
  Email string
}

func GetOneUser( userID string) *User {
  new := User{}
  new.UserID = "123"
  new.FirstName = "Jojo"
  new.Email = "scrap@gmail.com"

  return &new
}
