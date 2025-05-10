package entities

import "time"

type (
	User struct {
		ID           uint         `gorm:"primaryKey autoIncrement" json:"id"`
		Username     string       `gorm:"not null unique" json:"username"`
		Password     string       `gorm:"not null" json:"password"`
		ProfileURL   string       `json:"profile_url"`
		RefreshToken string       `json:"refresh_token"`
		CreatedAt    time.Time    `json:"created_at"`
		UpdatedAt    time.Time    `json:"updated_at"`
		Rooms        []Room       `gorm:"foreignKey:OwnerID;references:ID;constraint:OnDelete:CASCADE;" json:"-"`
		RoomMembers  []RoomMember `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;" json:"-"`
	}

	UserRepository interface {
		FindOneUser(username string) (*User, error)
		FindOneUserById(id uint) (*User, error)
		FindAllUsers() ([]User, error)
		Create(user *User) (*User, error)
		Update(user *User) (*User, error)
	}
)
