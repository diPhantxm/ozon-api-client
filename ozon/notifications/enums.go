package notifications

type MessageType string

const (
	PingType                MessageType = "TYPE_PING"
	NewPostingType          MessageType = "TYPE_NEW_POSTING"
	PostingCancelledType    MessageType = "TYPE_POSTING_CANCELLED"
	StateChangedType        MessageType = "TYPE_STATE_CHANGED"
	CutoffDateChangedType   MessageType = "TYPE_CUTOFF_DATE_CHANGED"
	DeliveryDateChangedType MessageType = "TYPE_DELIVERY_DATE_CHANGED"
	CreateOrUpdateType      MessageType = "TYPE_CREATE_OR_UPDATE_ITEM"
	PriceIndexChangedType   MessageType = "TYPE_PRICE_INDEX_CHANGED"
	StocksChangedType       MessageType = "TYPE_STOCKS_CHANGED"
	NewMessageType          MessageType = "TYPE_NEW_MESSAGE"
	UpdateMessageType       MessageType = "TYPE_UPDATE_MESSAGE"
	MessageReadType         MessageType = "TYPE_MESSAGE_READ"
	ChatClosedType          MessageType = "TYPE_CHAT_CLOSED"
)
