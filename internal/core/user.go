package core

type User struct {
	ID        int    `json:"ID" db:"id"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	UserName  string `json:"username" db:"username"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

func (u *User) Validate() error {
	if len(u.Email) == 0 {
		return ErrEmptyEmail
	}

	return nil
}
