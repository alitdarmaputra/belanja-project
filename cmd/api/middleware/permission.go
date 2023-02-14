package middleware

import (
	"net/http"

	"github.com/alitdarmaputra/belanja-project/constant"
	"github.com/alitdarmaputra/belanja-project/utils"
	"github.com/gin-gonic/gin"
)

var roles = []Role{
	{
		Id:   1,
		Name: "user",
		Permissions: map[string]string{
			constant.PermissionShowUser:    constant.PermissionShowUser,
			constant.PermissionUpdateUser:  constant.PermissionUpdateUser,
			constant.PermissionShowProduct: constant.PermissionShowProduct,
		},
	},
	{
		Id:   2,
		Name: "outlet",
		Permissions: map[string]string{
			constant.PermissionShowUser:      constant.PermissionShowUser,
			constant.PermissionUpdateUser:    constant.PermissionUpdateUser,
			constant.PermissionCreateProduct: constant.PermissionCreateProduct,
			constant.PermissionUpdateProduct: constant.PermissionUpdateProduct,
			constant.PermissionDeleteProduct: constant.PermissionDeleteProduct,
			constant.PermissionShowProduct:   constant.PermissionShowProduct,
		},
	},
}

func PermissionMiddleware(authetication Authetication, permissions ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, err := authetication.ExtractJWTUser(ctx)
		utils.PanicIfError(err)

		for _, permission := range permissions {
			if _, ok := roles[claims.RoleId-1].Permissions[permission]; !ok {
				ctx.AbortWithStatusJSON(
					http.StatusBadRequest,
					ErrResponse{
						Message:     "Unauthorized",
						Status:      http.StatusUnauthorized,
						Description: "Role does not have access",
					},
				)
			}
		}

		ctx.Next()
	}
}
