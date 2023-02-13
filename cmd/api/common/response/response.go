package response

import (
	"time"

	"github.com/gin-gonic/gin"
)

type BasicResponse struct {
	Code      int    `json:"code"`
	Status    string `json:"status"`
	Timestamp int64  `json:"timestamp"`
}

type BasicData struct {
	Code      int         `json:"code"`
	Status    string      `json:"status"`
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data"`
}

func JsonBasicResponse(ctx *gin.Context, code int, status string) {
	ctx.JSON(
		code,
		BasicResponse{
			Code:      code,
			Status:    status,
			Timestamp: time.Now().UnixNano(),
		},
	)
}

func JsonBasicData(ctx *gin.Context, code int, status string, data interface{}) {
	ctx.JSON(
		code,
		BasicData{
			Code:      code,
			Status:    status,
			Timestamp: time.Now().UnixNano(),
			Data:      data,
		},
	)
}
