package ozon

import (
	"context"
	"net/http"
	"testing"

	core "github.com/diphantxm/ozon-api-client"
)

func TestListChats(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ListChatsParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ListChatsParams{
				Filter: &ListChatsFilter{
					ChatStatus: "Opened",
					UnreadOnly: true,
				},
				Limit:  1,
				Offset: 0,
			},
			`{
				"chats": [
				  {
					"chat": {
					  "created_at": "2022-07-22T08:07:19.581Z",
					  "chat_id": "5e767w03-b400-4y1b-a841-75319ca8a5c8",
					  "chat_status": "Opened",
					  "chat_type": "Seller_Support"
					},
					"first_unread_message_id": 3000000000118021931,
					"last_message_id": 30000000001280042740,
					"unread_count": 1
				  }
				],
				"total_chats_count": 25,
				"total_unread_count": 5
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ListChatsParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Chats().List(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &ListChatsResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if len(resp.Chats) > 0 {
				if resp.Chats[0].Chat.ChatStatus == "" {
					t.Errorf("Chat status cannot be empty")
				}
				if resp.Chats[0].Chat.ChatType == "" {
					t.Errorf("Chat type cannot be empty")
				}
			}
		}
	}
}

func TestSendMessage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *SendMessageParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&SendMessageParams{
				ChatId: "99feb3fc-a474-469f-95d5-268b470cc607",
				Text:   "test",
			},
			`{
				"result": "success"
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&SendMessageParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Chats().SendMessage(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &SendMessageResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestSendFile(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *SendFileParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&SendFileParams{
				ChatId:        "99feb3fc-a474-469f-95d5-268b470cc607",
				Name:          "tempor",
				Base64Content: "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII=",
			},
			`{
				"result": "success"
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&SendFileParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Chats().SendFile(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &SendFileResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestChatHistory(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *ChatHistoryParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&ChatHistoryParams{
				ChatId:        "18b8e1f9-4ae7-461c-84ea-8e1f54d1a45e",
				Direction:     "Forward",
				FromMessageId: "3000000000118032000",
				Limit:         1,
			},
			`{
				"has_next": true,
				"messages": [
				  {
					"message_id": "3000000000817031942",
					"user": {
					  "id": "115568",
					  "type": "Сustomer"
					},
					"created_at": "2022-07-18T20:58:04.528Z",
					"is_read": true,
					"data": [
					  "Здравствуйте, у меня вопрос по вашему товару \"Стекло защитное для смартфонов\", артикул 11223. Подойдет ли он на данную [ модель ](https://www.ozon.ru/product/smartfon-samsung-galaxy-a03s-4-64-gb-chernyy) телефона?"
					]
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&ChatHistoryParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Chats().History(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &ChatHistoryResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestUpdateChat(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *UpdateChatParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&UpdateChatParams{
				ChatId:        "99feb3fc-a474-469f-95d5-268b470cc607",
				FromMessageId: 0,
				Limit:         1000,
			},
			`{
				"result": [
				  {
					"id": 3000000000012735500,
					"user": {
					  "id": "15",
					  "type": "seller"
					},
					"type": "file",
					"text": "",
					"file": {
					  "url": "https://cdn-stg.ozonru.me/s3/fs-chat-api/108a0370-4dfa-11ec-bd02-06a332735108.png",
					  "mime": "image/png",
					  "size": 68,
					  "name": "tempor"
					},
					"created_at": "2021-11-25T14:14:55Z",
					"context": null
				  }
				]
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&UpdateChatParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Chats().Update(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &UpdateChatResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}

func TestCreateNewChat(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *CreateNewChatParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&CreateNewChatParams{
				PostingNumber: "47873153-0052-1",
			},
			`{
				"result": {
				  "chat_id": "5969c331-2e64-44b7-8a0e-ff9526762c62"
				}
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&CreateNewChatParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Chats().Create(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &CreateNewChatResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}

		if resp.StatusCode == http.StatusOK {
			if resp.Result.ChatId == "" {
				t.Errorf("Chat id cannot be empty")
			}
		}
	}
}

func TestMarkAsRead(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		headers    map[string]string
		params     *MarkAsReadParams
		response   string
	}{
		// Test Ok
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			&MarkAsReadParams{
				ChatId:        "99feb3fc-a474-469f-95d5-268b470cc607",
				FromMessageId: 3000000000118032000,
			},
			`{
				"unread_count": 0
			}`,
		},
		// Test No Client-Id or Api-Key
		{
			http.StatusUnauthorized,
			map[string]string{},
			&MarkAsReadParams{},
			`{
				"code": 16,
				"message": "Client-Id and Api-Key headers are required"
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		ctx, _ := context.WithTimeout(context.Background(), testTimeout)
		resp, err := c.Chats().MarkAsRead(ctx, test.params)
		if err != nil {
			t.Error(err)
			continue
		}

		compareJsonResponse(t, test.response, &MarkAsReadResponse{})

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}
