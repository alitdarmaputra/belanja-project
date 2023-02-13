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
	Id          int     `json:"id"           binding:"required,numeric"`
	Email       string  `json:"email"        binding:"required,email"`
	FullName    string  `json:"full_name"    binding:"required"`
	PhoneNumber string  `json:"phone_number" binding:"required,numeric"`
	Address     string  `json:"address"      binding:"required"`
	AreaId      int     `json:"area_id"      binding:"required,numeric"`
	RoleId      int8    `json:"role_id"      binding:"required,numeric"`
	Latitude    float64 `json:"latitude"     binding:"required,geolocation"`
	Longitude   float64 `json:"longitude"    binding:"required,geolocation"`
}
