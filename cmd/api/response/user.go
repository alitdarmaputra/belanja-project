package response

import "github.com/alitdarmaputra/belanja-project/modules/database/user"

type UserResponse struct {
	Id          int     `json:"id"`
	Email       string  `json:"email"`
	FullName    string  `json:"full_name"`
	PhoneNumber string  `json:"phone_number"`
	Address     string  `json:"address"`
	AreaId      int     `json:"area_id"`
	RoleId      int8    `json:"role_id"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

func ToUserResponse(user user.User) UserResponse {
	return UserResponse{
		Id:          user.Id,
		Email:       user.Email,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		AreaId:      user.AreaId,
		RoleId:      user.RoleId,
		Latitude:    user.Latitude,
		Longitude:   user.Longitude,
	}
}

func ToUserResponses(users []user.User) []UserResponse {
	var usersResponses []UserResponse
	for _, user := range users {
		usersResponses = append(usersResponses, ToUserResponse(user))
	}
	return usersResponses
}
