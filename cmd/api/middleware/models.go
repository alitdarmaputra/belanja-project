package middleware

type ErrResponse struct {
	Message     string `json:"message"`
	Status      uint16 `json:"status"`
	Description string `json:"description"`
}

type Role struct {
	Id          int
	Name        string
	Permissions map[string]string
}
