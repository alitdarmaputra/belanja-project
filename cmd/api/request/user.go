package request

type UserCreateRequest struct {
	Email       string  `json:"email"        binding:"required,email"`
	FullName    string  `json:"full_name"    binding:"required"`
	Password    string  `json:"password"     binding:"required"`
	PhoneNumber string  `json:"phone_number" binding:"required,numeric"`
	Address     string  `json:"address"      binding:"required"`
	AreaId      int     `json:"area_id"      binding:"required,numeric"`
	RoleId      int8    `json:"role_id"      binding:"required,numeric"`
	Latitude    float64 `json:"latitude"     binding:"required,latitude"`
	Longitude   float64 `json:"longitude"    binding:"required,longitude"`
}

type UserUpdateRequest struct {
	Email       string  `json:"email"        binding:"required,email"`
	FullName    string  `json:"full_name"    binding:"required"`
	PhoneNumber string  `json:"phone_number" binding:"required,numeric"`
	Address     string  `json:"address"      binding:"required"`
	AreaId      int     `json:"area_id"      binding:"required,numeric"`
	RoleId      int8    `json:"role_id"      binding:"required,numeric"`
	Latitude    float64 `json:"latitude"     binding:"required,latitude"`
	Longitude   float64 `json:"longitude"    binding:"required,longitude"`
}

type UserLoginRequest struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type ChangePasswordRequest struct {
	NewPassword string `json:"new_password" binding:"required"`
}
