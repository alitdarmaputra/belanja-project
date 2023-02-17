package shipper

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/alitdarmaputra/belanja-project/bussiness"
	"github.com/alitdarmaputra/belanja-project/utils"
	"github.com/go-resty/resty/v2"
)

type Client = resty.Client

type ShipperServiceImpl struct {
	*resty.Client
}

func NewShipperService(baseUrl, key string) ShipperService {
	c := resty.New()
	c.SetBaseURL(baseUrl)
	c.SetHeader("X-API-Key", key)
	c.SetHeader("Content-Type", "application/json")
	c.SetTimeout(180 * time.Second)

	return &ShipperServiceImpl{c}
}

func (service *ShipperServiceImpl) CreateOrder(
	ctx context.Context,
	request ShipperCreateRequest,
) string {
	ctxWT, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	r := service.NewRequest()
	r.SetContext(ctxWT)
	r.SetBody(request)

	res, err := r.Post(service.BaseURL + "/v3/order")
	if err != nil {
		panic(bussiness.NewBadGateWayError(err.Error()))
	} else if res.StatusCode() != http.StatusCreated {
		panic(bussiness.NewBadGateWayError("Error from shipper api"))
	}

	result := ShipperCreateResposne{}
	err = json.Unmarshal(res.Body(), &result)
	utils.PanicIfError(err)

	return result.Data.OrderId
}

func (service *ShipperServiceImpl) CancelOrder(
	ctx context.Context,
	shipperId string,
) {
	ctxWT, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	r := service.NewRequest()
	r.SetContext(ctxWT)

	_, err := r.Delete(service.BaseURL + "/v3/order/" + shipperId)

	if err != nil {
		panic(bussiness.NewBadGateWayError(err.Error()))
	}
}
