package gomock


type User struct {
	Person Person
}

func NewUser(p Person) *User {
	return &User{Person: p}
}

func (u *User) GetUserInfo(id int64) (string, error) {
	return u.Person.Get(id)
}
