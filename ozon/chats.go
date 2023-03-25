package ozon

import (
	"net/http"
	"time"

	core "github.com/diphantxm/ozon-api-client"
)

type Chats struct {
	client *core.Client
}

type ListChatsParams struct {
	// Chats filter
	Filter ListChatsFilter `json:"filter"`

	// Number of values in the response. Default value is 1
	Limit int64 `json:"limit" default:"1"`

	// Number of elements that will be skipped in the response.
	// For example, if offset=10, the response will start with the 11th element found
	Offset int64 `json:"offset"`
}

type ListChatsFilter struct {
	// Filter by chat status:
	//   - All
	//   - Opened
	//   - Closed
	ChatStatus string `json:"chat_status" default:"ALL"`

	// Filter by chats with unread messages
	UnreadOnly bool `json:"unread_only"`
}

type ListChatsResponse struct {
	core.CommonResponse

	// Chats data
	Chats []struct {
		// Chat data
		Chat struct {
			// Chat identifier
			ChatId string `json:"chat_id"`

			// Chat status:
			//   - All
			//   - Opened
			//   - Closed
			ChatStatus string `json:"chat_status"`

			// Chat type:
			//   - Seller_Support — support chat
			//   - Buyer_Seller — chat with a customer
			ChatType string `json:"chat_type"`

			// Chat creation date
			CreatedAt time.Time `json:"created_at"`
		} `json:"chat"`

		// Identifier of the first unread chat message
		FirstUnreadMessageId string `json:"first_unread_message_id"`

		// Identifier of the last message in the chat
		LastMessageId string `json:"last_message_id"`

		// Number of unread messages in the chat
		UnreadCount int64 `json:"unread_count"`
	} `json:"chats"`

	// Total number of chats
	TotalChatsCount int64 `json:"total_chats_count"`

	// Total number of unread messages
	TotalUnreadCount int64 `json:"total_unread_count"`
}

// Returns information about chats by specified filters
func (c Chats) List(params *ListChatsParams) (*ListChatsResponse, error) {
	url := "/v2/chat/list"

	resp := &ListChatsResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type SendMessageParams struct {
	// Chat identifier
	ChatId string `json:"chat_id"`

	// Message text in the plain text format
	Text string `json:"text"`
}

type SendMessageResponse struct {
	core.CommonResponse

	// Method result
	Result string `json:"result"`
}

// Sends a message to an existing chat by its identifier
func (c Chats) SendMessage(params *SendMessageParams) (*SendMessageResponse, error) {
	url := "/v1/chat/send/message"

	resp := &SendMessageResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type SendFileParams struct {
	// File as a base64 string
	Base64Content string `json:"base64_content"`

	// Chat identifier
	ChatId string `json:"chat_id"`

	// File name with extension
	Name string `json:"name"`
}

type SendFileResponse struct {
	core.CommonResponse

	// Method result
	Result string `json:"result"`
}

// Sends a file to an existing chat by its identifier
func (c Chats) SendFile(params *SendFileParams) (*SendFileResponse, error) {
	url := "/v1/chat/send/file"

	resp := &SendFileResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type ChatHistoryParams struct {
	// Chat idenitifier
	ChatId string `json:"chat_id"`

	// Messages sorting direction:
	//   - Forward—from old messages to new ones.
	//   - Backward—from new messages to old ones.
	// The default value is `Backward`. You can set the number of messages in the limit parameter
	Direction string `json:"direction" default:"Backward"`

	// Identifier of the message from which the chat history will be displayed.
	// Default value is the last visible message
	FromMessageId string `json:"from_message_id"`

	// Number of messages in the response. The default value is 50
	Limit int64 `json:"limit" default:"50"`
}

type ChatHistoryResponse struct {
	core.CommonResponse

	// Indicates that the response returned only a part of messages
	HasNext bool `json:"has_next"`

	// An array of messages sorted according to the direction parameter in the request body
	Messages []struct {
		// Message creation date
		CreatedAt time.Time `json:"created_at"`

		// Array with message content in Markdown format
		Data []string `json:"data"`

		// Indication of the read message
		IsRead bool `json:"is_read"`

		// Message identifier
		MessageId string `json:"message_id"`

		// Chat participant identifier
		User struct {
			// Chat participant identifier
			Id string `json:"id"`

			// Chat participant type:
			//   - customer
			//   - seller
			//   - crm—system messages
			//   - courier
			//   - support
			Type string `json:"type"`
		} `json:"user"`
	} `json:"messages"`
}

// Chat history
func (c Chats) History(params *ChatHistoryParams) (*ChatHistoryResponse, error) {
	url := "/v2/chat/history"

	resp := &ChatHistoryResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type UpdateChatParams struct {
	// Chat identifier
	ChatId string `json:"chat_id"`

	// Message identifier
	FromMessageId uint64 `json:"from_message_id"`

	// Number of messages in the response
	Limit int64 `json:"limit"`
}

type UpdateChatResponse struct {
	core.CommonResponse

	// Method result
	Result []struct {
		// An order or a product user wrote about in the chat
		Context struct {
			// Product inforamtion
			Item struct {
				// Product identifier in the Ozon system, SKU
				SKU int64 `json:"sku"`
			} `json:"item"`

			// Order information
			Order struct {
				// Order number
				OrderNumber string `json:"order_number"`

				// Shipment information
				Postings []struct {
					// Delivery scheme:
					//   - FBO
					//   - FBS
					//   - RFBS
					//   - Crossborder
					DeliverySchema string `json:"delivery_schema"`

					// Shipment number
					PostingNumber string `json:"posting_number"`

					// List of product identifiers in the shipment
					SKUList []int64 `json:"sku_list"`
				} `json:"postings"`
			} `json:"order"`
		} `json:"context"`

		// Creation date and time
		CreatedAt time.Time `json:"created_at"`

		// Information about the file in the chat. Displayed only for `type = file`
		File struct {
			// File type
			Mime string `json:"mime"`

			// File name
			Name string `json:"name"`

			// File size in bytes
			Size int64 `json:"size"`

			// File URL
			URL string `json:"url"`
		} `json:"file"`

		// File identifier
		Id uint64 `json:"id"`

		// Message. Displayed only for `type = text`
		Text string `json:"text"`

		// Message type:
		//   - text
		//   - file
		Type string `json:"type"`

		// Chat participant information
		User struct {
			// Chat participant identifier
			Id string `json:"id"`

			// Chat participant chat:
			//   - customer
			//   - seller
			//   - crm—system messages
			//   - courier
			Type string `json:"type"`
		} `json:"user"`
	} `json:"result"`
}

// Update chat
func (c Chats) Update(params *UpdateChatParams) (*UpdateChatResponse, error) {
	url := "/v1/chat/updates"

	resp := &UpdateChatResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type CreateNewChatParams struct {
	// Shipment identifier
	PostingNumber string `json:"posting_number"`
}

type CreateNewChatResponse struct {
	core.CommonResponse

	//Method result
	Result struct {
		// Chat identifier
		ChatId string `json:"chat_id"`
	} `json:"result"`
}

// Creates a new chat on the shipment with the customer. For example, to clarify the address or the product model
func (c Chats) Create(params *CreateNewChatParams) (*CreateNewChatResponse, error) {
	url := "/v1/chat/start"

	resp := &CreateNewChatResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}

type MarkAsReadParams struct {
	// Chat identifier
	Chatid string `json:"chat_id"`

	// Message identifier
	FromMessageId uint64 `json:"from_message_id"`
}

type MarkAsReadResponse struct {
	core.CommonResponse

	// Number of unread messages in the chat
	UnreadCount int64 `json:"unread_count"`
}

// A method for marking the selected message and messages before it as read
func (c Chats) MarkAsRead(params *MarkAsReadParams) (*MarkAsReadResponse, error) {
	url := "/v2/chat/read"

	resp := &MarkAsReadResponse{}

	response, err := c.client.Request(http.MethodPost, url, params, resp, nil)
	if err != nil {
		return nil, err
	}
	response.CopyCommonResponse(&resp.CommonResponse)

	return resp, nil
}
