package user

type User struct {
	Email      string
	Name       string
	Age        int
	Subscribed bool
}

type Role string

const (
	Unidentified Role = "unidentified"
	Visitor      Role = "visitor"
	Customer     Role = "customer"
	Manager      Role = "manager"
	Admin        Role = "admin"
)

type UserStore interface {
	Get(string) (*User, error)
	Store(*User) error
}
