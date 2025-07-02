package participant_role

type Role int

const (
	ADMIN Role = iota + 1
	USER
)

func (r Role) String() string {
	return []string{"Admin", "User"}[r-1]
}
