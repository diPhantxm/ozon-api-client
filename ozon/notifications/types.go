package notifications

import "time"

// Checking if the service is ready at initial connection and periodically after it
type pingRequest struct {
	Common

	// Date and time when the notification was sent in UTC format
	Time time.Time `json:"time"`
}

type pingResponse struct {
	// Application version
	Version string `json:"version"`

	// Application name
	Name string `json:"name"`

	// Date and time when notification processing started in UTC format
	Time time.Time `json:"time"`
}

type Common struct {
	MessageType MessageType `json:"message_type"`
}

// New shipment
type NewPosting struct {
	Common

	// Shipment number
	PostingNumber string `json:"posting_number"`

	// Products information
	Products []Product `json:"products"`

	// Date and time when the shipment processing started in the UTC format
	InProccessAt time.Time `json:"in_process_at"`

	// Warehouse identifier where the products for this shipment are stored
	WarehouseId int64 `json:"warehouse_id"`

	// Seller identifier
	SellerId string `json:"seller_id"`
}

type Product struct {
	// Product SKU
	SKU int64 `json:"sku"`

	// Product quantity
	Quantity int64 `json:"quantity"`
}

// Shipment cancellation
type PostingCancelled struct {
	Common

	// Shipment number
	PostingNumber string `json:"posting_number"`

	// Products information
	Products []Product `json:"products"`

	// Previous shipment status
	OldState string `json:"old_state"`

	// New shipment status: posting_canceled—canceled
	NewState string `json:"new_state"`

	// Date and time when the shipment status was changed in UTC format
	ChangedStateDate time.Time `json:"changed_state_date"`

	// Information about cancellation reason
	Reason Reason `json:"reason"`

	// Warehouse identifier where the products for this shipment are stored
	WarehouseId int64 `json:"warehouse_id"`

	// Seller identifier
	SellerId string `json:"seller_id"`
}

type Reason struct {
	// Cancellation reason identifier
	Id int64 `json:"id"`

	// Cancellation reason
	Message string `json:"message"`
}

// Shipment status change
type StateChanged struct {
	Common

	// Shipment number
	PostingNumber string `json:"posting_number"`

	// New shipment status
	NewState string `json:"new_state"`

	// Date and time when the shipment status was changed in UTC format
	ChangedStateDate time.Time `json:"chagned_state_date"`

	// Warehouse identifier where the products for this shipment are stored
	WarehouseId int64 `json:"warehouse_id"`

	// Seller identifier
	SellerId string `json:"seller_id"`
}

// Shipment shipping date change
type CutoffDateChanged struct {
	Common

	// Shipment number
	PostingNumber string `json:"posting_number"`

	// New shipping date and time in UTC format
	NewCutoffDate time.Time `json:"new_cutoff_date"`

	// Previous shipping date and time in UTC format
	OldCutoffDate time.Time `json:"old_cutoff_date"`

	// Warehouse identifier where the products for this shipment are stored
	WarehouseId int64 `json:"warehouse_id"`

	// Seller identifier
	SellerId string `json:"seller_id"`
}

// Shipment delivery date change
type DeliveryDateChanged struct {
	Common

	// Shipment number
	PostingNumber string `json:"posting_number"`

	// New delivery start date and time in UTC format
	NewDeliveryDateBegin time.Time `json:"new_delivery_date_begin"`

	// New delivery end date and time in UTC format
	NewDeliveryDateEnd time.Time `json:"new_delivery_date_end"`

	// Previous delivery start date and time in UTC format
	OldDeliveryDateBegin time.Time `json:"old_delivery_date_begin"`

	// Previous delivery end date and time in UTC format
	OldDeliveryDateEnd time.Time `json:"old_delivery_date_end"`

	// Warehouse identifier where the products for this shipment are stored
	WarehouseId int64 `json:"warehouse_id"`

	// Seller identifier
	SellerId string `json:"seller_id"`
}

// Product creation and update or processing error
type CreateOrUpdateItem struct {
	Common

	// Product identifier in the seller's system
	OfferId string `json:"offer_id"`

	// Product identifier
	ProductId int64 `json:"product_id"`

	// An indication that errors occurred during the product creation or update
	IsError bool `json:"is_error"`

	// Update date and time
	ChangedAt time.Time `json:"changed_at"`

	// Seller identifier
	SellerId string `json:"seller_id"`
}

// Product price index change
type PriceIndexChanged struct {
	Common

	// Date and time of price index change
	UpdatedAt time.Time `json:"updated_at"`

	// Product SKU
	SKU int64 `json:"sku"`

	// Product identifier
	ProductId int64 `json:"product_id"`

	// Price index
	PriceIndex int64 `json:"price_index"`

	// Seller identifier
	SellerId string `json:"seller_id"`
}

// Stock change at the seller's warehouse
type StocksChanged struct {
	Common

	// Array with products data
	Items []Item `json:"items"`

	// Seller identifier
	SellerId string `json:"seller_id"`
}

type Item struct {
	// Update date and time
	UpdatedAt time.Time `json:"updated_at"`

	// Product SKU when working under the FBS or rFBS schemes
	SKU int64 `json:"sku"`

	// Product identifier
	ProductId int64 `json:"product_id"`

	// Array with product stocks data
	Stocks []Stock `json:"stocks"`
}

type Stock struct {
	// Warehouse identifier
	WarehouseId int64 `json:"warehouse_id"`

	// Total product stocks at the warehouse
	Present int64 `json:"present"`

	// Number of reserved products at the warehouse
	Reserved int64 `json:"reserved"`
}

// New message in chat
type NewMessage struct {
	Common

	// Chat identifier
	ChatId string `json:"chat_id"`

	// Chat type
	ChatType string `json:"chat_type"`

	// Message identifier
	MessageId string `json:"message_id"`

	// Message creation date
	CreatedAt time.Time `json:"created_at"`

	// Information about message sender
	User User `json:"user"`

	// Array with message content in Markdown format
	Data []string `json:"data"`

	// Seller identifier
	SellerId string `json:"seller_id"`
}

type User struct {
	// Sender identifier
	Id string `json:"id"`

	// Sender type
	Type string `json:"type"`
}

// Message in chat has changed
type UpdateMessage struct {
	NewMessage

	// Message update date
	UpdatedAt time.Time `json:"updated_at"`
}

// Customer or support read your message
type MessageRead struct {
	NewMessage

	// Last read message identifier
	LastReadMessageId string `json:"last_read_message_id"`
}

// Chat is closed
type ChatClosed struct {
	Common

	// Chat identifier
	ChatId string `json:"chat_id"`

	// Chat type
	ChatType string `json:"chat_type"`

	// Information about the user who closed the chat
	User User `json:"user"`

	// User identifier
	Id string `json:"id"`

	// User type
	Type string `json:"type"`

	// Seller identifier
	SellerId string `json:"seller_id"`
}

type Response struct {
	// Notification is received
	Result bool `json:"result"`
}

type errorResponse struct {
	// Information about the error
	Data errorData `json:"error"`
}

type errorData struct {
	// Error code
	Code string `json:"code"`

	// Detailed error description
	Message string `json:"message"`

	// Additional information
	Details string `json:"details"`
}
