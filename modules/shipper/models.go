package shipper

type Consignee struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type Destination struct {
	Address string `json:"address" binding:"required"`
	AreaId  int    `json:"area_id" binding:"required"`
	Lat     string `json:"lat"     binding:"required"`
	Lng     string `json:"lng"     binding:"required"`
}

type Origin struct {
	AreaId  int    `json:"area_id"`
	Address string `json:"address"`
	Lat     string `json:"lat"`
	Lng     string `json:"lng"`
}

type Item struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Qty   int    `json:"qty"`
}

type Package struct {
	Height      int    `json:"height"`
	Length      int    `json:"length"`
	Width       int    `json:"width"`
	Weight      int    `json:"weight"`
	Items       []Item `json:"items"`
	Price       int    `json:"price"`
	PackageType int    `json:"package_type"`
}

type ShipperCreateRequest struct {
	Consignee   Consignee   `json:"consignee"`
	Coverage    string      `json:"coverage"`
	Destination Destination `json:"destination"`
	Origin      Origin      `json:"origin"`
	Package     Package     `json:"package"`
	Consigner   Consignee   `json:"consigner"`
	PaymentType string      `json:"payment_type"`
}

type Metadata struct {
	Path       string `json:"path"`
	StatusCode uint8  `json:"http_status_code"`
	Status     string `json:"http_status"`
	Timestamp  int64  `json:"timestamp"`
}

type ShipperCreateData struct {
	OrderId string `json:"order_id"`
}

type ShipperCreateResposne struct {
	Metadata Metadata          `json:"metadata"`
	Data     ShipperCreateData `json:"data"`
}
