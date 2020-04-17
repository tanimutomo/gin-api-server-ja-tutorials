package domain

type User struct {
	ID          int
	ScreenName  string
	DisplayName string
	Password    string
	Email       *string
	CreatedAt   int64
	UpdatedAt   int64
}

type UserForGet struct {
	ID          int     `json:"id"`
	ScreenName  string  `json:"screenName`
	DisplayName string  `json:"displayName`
	Email       *string `json:"email"`
}

func (u *User) BuildForGet() UserForGet {
	user := UserForGet{}
	user.ID = u.ID
	user.ScreenName = u.ScreenName
	user.DisplayName = u.DisplayName
	if u.Email != nil {
		user.Email = u.Email
	} else {
		empty := ""
		user.Email = &empty
	}
	return user
}
