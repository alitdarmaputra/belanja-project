package product

import (
	"net/http"

	"github.com/alitdarmaputra/belanja-project/bussiness/product"
	"github.com/alitdarmaputra/belanja-project/cmd/api/common/response"
	"github.com/alitdarmaputra/belanja-project/cmd/api/middleware"
	"github.com/alitdarmaputra/belanja-project/cmd/api/request"
	"github.com/alitdarmaputra/belanja-project/utils"
	"github.com/gin-gonic/gin"
)

type ProductContorllerImpl struct {
	ProductService product.ProductService
	Middleware     middleware.Authetication
}

func NewProductController(
	productService product.ProductService,
	middleware middleware.Authetication,
) ProductController {
	return &ProductContorllerImpl{
		ProductService: productService,
		Middleware:     middleware,
	}
}

func (controller *ProductContorllerImpl) Create(ctx *gin.Context) {
	claims, err := controller.Middleware.ExtractJWTUser(ctx)
	utils.PanicIfError(err)

	productCreateRequest := request.ProductCreateRequest{}
	err = ctx.ShouldBind(&productCreateRequest)
	utils.PanicIfError(err)

	productResponse := controller.ProductService.Create(ctx, productCreateRequest, claims.Id)
	response.JsonBasicData(ctx, http.StatusCreated, "Created", productResponse)
}

func (controller *ProductContorllerImpl) Update(ctx *gin.Context) {
	claims, err := controller.Middleware.ExtractJWTUser(ctx)
	utils.PanicIfError(err)

	productUpdateRequest := request.ProductUpdateRequest{}
	err = ctx.ShouldBind(&productUpdateRequest)
	utils.PanicIfError(err)

	productResponse := controller.ProductService.Update(ctx, productUpdateRequest, claims.Id)
	response.JsonBasicData(ctx, http.StatusOK, "OK", productResponse)
}

func (controller *ProductContorllerImpl) Delete(ctx *gin.Context) {
	claims, err := controller.Middleware.ExtractJWTUser(ctx)
	utils.PanicIfError(err)

	uri := request.PathParam{}
	err = ctx.ShouldBindUri(&uri)
	utils.PanicIfError(err)

	controller.ProductService.Delete(ctx, uri.Id, claims.Id)
	response.JsonBasicResponse(ctx, http.StatusOK, "OK")
}

func (controller *ProductContorllerImpl) FindById(ctx *gin.Context) {
	uri := request.PathParam{}
	err := ctx.ShouldBindUri(&uri)
	utils.PanicIfError(err)

	productResponse := controller.ProductService.FindById(ctx, uri.Id)

	response.JsonBasicData(ctx, http.StatusOK, "OK", productResponse)
}

func (controller *ProductContorllerImpl) FindAll(ctx *gin.Context) {
	productResponses := controller.ProductService.FindAll(ctx)
	response.JsonBasicData(ctx, http.StatusOK, "OK", productResponses)
}
