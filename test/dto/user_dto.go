package dto

import "test/model"

type UserDto struct {
	ID        uint   `json:"id"`
	Telephone string `json:"telephone"`
	Role      uint   `json:"role"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		ID:        user.ID,
		Telephone: user.Telephone,
		Role:      user.Role,
	}
}
