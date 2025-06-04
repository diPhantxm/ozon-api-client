package fbs

import "time"

type ListRequestLastChangedStatusDate struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

type ListRequestFilter struct {
	DeliveryMethodID      []string                         `json:"delivery_method_id"`
	IsQuantum             bool                             `json:"is_quantum"`
	LastChangedStatusDate ListRequestLastChangedStatusDate `json:"last_changed_status_date"`
	OrderID               int64                            `json:"order_id"`
	ProviderID            []string                         `json:"provider_id"`
	Since                 time.Time                        `json:"since"`
	Status                string                           `json:"status"`
	To                    time.Time                        `json:"to"`
	WarehouseID           []string                         `json:"warehouse_id"`
}

type ListRequestWith struct {
	AnalyticsData bool `json:"analytics_data"`
	Barcodes      bool `json:"barcodes"`
	FinancialData bool `json:"financial_data"`
	Translit      bool `json:"translit"`
}

type ListRequest struct {
	Dir    string            `json:"dir"`
	Filter ListRequestFilter `json:"filter"`
	Limit  int64             `json:"limit"`
	Offset int64             `json:"offset"`
	With   ListRequestWith   `json:"with"`
}

type ListResponseResultPostingDeliveryMethod struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	WarehouseID   int64  `json:"warehouse_id"`
	Warehouse     string `json:"warehouse"`
	TplProviderID int64  `json:"tpl_provider_id"`
	TplProvider   string `json:"tpl_provider"`
}

type ListResponseResultPostingOptional struct {
	ProductsWithPossibleMandatoryMark []int `json:"products_with_possible_mandatory_mark"`
}

type ListResponseResultPostingCancellation struct {
	CancelReasonID           int64  `json:"cancel_reason_id"`
	CancelReason             string `json:"cancel_reason"`
	CancellationType         string `json:"cancellation_type"`
	CancelledAfterShip       bool   `json:"cancelled_after_ship"`
	AffectCancellationRating bool   `json:"affect_cancellation_rating"`
	CancellationInitiator    string `json:"cancellation_initiator"`
}

type ListResponseResultPostingCustomerAddress struct {
	AddressTail     string  `json:"address_tail"`
	City            string  `json:"city"`
	Comment         string  `json:"comment"`
	Country         string  `json:"country"`
	District        string  `json:"district"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	ProviderPvzCode string  `json:"provider_pvz_code"`
	PvzCode         int64   `json:"pvz_code"`
	Region          string  `json:"region"`
	ZipCode         string  `json:"zip_code"`
}

type ListResponseResultPostingCustomer struct {
	Address    ListResponseResultPostingCustomerAddress `json:"address"`
	CustomerID int64                                    `json:"customer_id"`
	Name       string                                   `json:"name"`
	Phone      string                                   `json:"phone"`
}

type ListResponseResultPostingProduct struct {
	Price          string `json:"price"`
	CurrencyCode   string `json:"currency_code"`
	IsBlrTraceable bool   `json:"is_blr_traceable"`
	OfferID        string `json:"offer_id"`
	Name           string `json:"name"`
	SKU            int64  `json:"sku"`
	Quantity       int    `json:"quantity"`
}

type ListResponseResultPostingAdressee struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type ListResponseResultPostingBarcode struct {
	LowerBarcode string `json:"lower_barcode"`
	UpperBarcode string `json:"upper_barcode"`
}

type ListResponseResultPostingAnalyticsData struct {
	City                 string `json:"city"`
	DeliveryDateBegin    string `json:"delivery_date_begin"`
	DeliveryDateEnd      string `json:"delivery_date_end"`
	DeliveryType         string `json:"delivery_type"`
	IsLegal              bool   `json:"is_legal"`
	IsPremium            bool   `json:"is_premium"`
	PaymentTypeGroupName string `json:"payment_type_group_name"`
	Region               string `json:"region"`
	TplProvider          string `json:"tpl_provider"`
	TplProviderID        int64  `json:"tpl_provider_id"`
	Warehouse            string `json:"warehouse"`
	WarehouseID          int64  `json:"warehouse_id"`
}

type ListResponseResultPostingFinancialDataProductPicking struct {
	Amount float64 `json:"amount"`
	Moment string  `json:"moment"`
	Tag    string  `json:"tag"`
}

type ListResponseResultPostingFinancialDataProduct struct {
	Actions                []string                                              `json:"actions"`
	CurrencyCode           string                                                `json:"currencyCode"`
	CommissionAmount       float64                                               `json:"commission_amount"`
	CommissionPercent      int64                                                 `json:"commission_percent"`
	CommissionCurrencyCode string                                                `json:"commission_currency_code"`
	OldPrice               float64                                               `json:"old_price"`
	Payout                 float64                                               `json:"payout"`
	Picking                *ListResponseResultPostingFinancialDataProductPicking `json:"picking"`
	Price                  float64                                               `json:"price"`
	ProductID              int64                                                 `json:"product_id"`
	Quantity               int64                                                 `json:"quantity"`
	TotalDiscountPercent   float64                                               `json:"total_discount_percent"`
	TotalDiscountValue     float64                                               `json:"total_discount_value"`
}

type ListResponseResultPostingFinancialData struct {
	ClusterFrom string                                          `json:"cluster_from"`
	ClusterTo   string                                          `json:"cluster_to"`
	Products    []ListResponseResultPostingFinancialDataProduct `json:"products"`
}

type ListResponseResultPostingRequirement struct {
	ProductsRequiringGtd           []string `json:"products_requiring_gtd"`
	ProductsRequiringCountry       []string `json:"products_requiring_country"`
	ProductsRequiringMandatoryMark []string `json:"products_requiring_mandatory_mark"`
	ProductsRequiringJwUin         []string `json:"products_requiring_jw_uin"`
}

type ListResponseResultPostingTariffication struct {
	CurrentTariffRate               float64   `json:"current_tariff_rate"`
	CurrentTariffType               string    `json:"current_tariff_type"`
	CurrentTariffCharge             string    `json:"current_tariff_charge"`
	CurrentTariffChargeCurrencyCode string    `json:"current_tariff_charge_currency_code"`
	NextTariffRate                  float64   `json:"next_tariff_rate"`
	NextTariffType                  string    `json:"next_tariff_type"`
	NextTariffCharge                string    `json:"next_tariff_charge"`
	NextTariffStartsAt              time.Time `json:"next_tariff_starts_at"`
	NextTariffChargeCurrencyCode    string    `json:"next_tariff_charge_currency_code"`
}

type ListResponseResultPosting struct {
	PostingNumber        string                                  `json:"posting_number"`
	OrderID              int64                                   `json:"order_id"`
	OrderNumber          string                                  `json:"order_number"`
	PickupCodeVerifiedAt time.Time                               `json:"pickup_code_verified_at"`
	Status               string                                  `json:"status"`
	Substatus            string                                  `json:"substatus"`
	DeliveryMethod       ListResponseResultPostingDeliveryMethod `json:"delivery_method"`
	TrackingNumber       string                                  `json:"tracking_number"`
	TplIntegrationType   string                                  `json:"tpl_integration_type"`
	InProcessAt          time.Time                               `json:"in_process_at"`
	ShipmentDate         time.Time                               `json:"shipment_date"`
	DeliveringDate       *string                                 `json:"delivering_date"`
	Optional             ListResponseResultPostingOptional       `json:"optional"`
	Cancellation         ListResponseResultPostingCancellation   `json:"cancellation"`
	Customer             *ListResponseResultPostingCustomer      `json:"customer"`
	Products             []ListResponseResultPostingProduct      `json:"products"`
	Addressee            *ListResponseResultPostingAdressee      `json:"addressee"`
	Barcodes             *ListResponseResultPostingBarcode       `json:"barcodes"`
	AnalyticsData        *ListResponseResultPostingAnalyticsData `json:"analytics_data"`
	FinancialData        *ListResponseResultPostingFinancialData `json:"financial_data"`
	IsExpress            bool                                    `json:"is_express"`
	QuantumID            int64                                   `json:"quantum_id"`
}

type ListResponseResult struct {
	Postings []ListResponseResultPosting `json:"postings"`
	HasNext  bool                        `json:"has_next"`
}

type ListResponse struct {
	Result ListResponseResult `json:"result"`
}
