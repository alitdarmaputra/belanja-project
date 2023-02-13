package request

type PathParam struct {
	Id int `uri:"id" binding:"required,numeric"`
}
