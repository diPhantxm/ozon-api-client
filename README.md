# Ozon Seller API Client
A Ozon Seller API client written in Golang

[![Coverage Status](https://coveralls.io/repos/github/diPhantxm/ozon-api-client/badge.svg)](https://coveralls.io/github/diPhantxm/ozon-api-client)
![example workflow](https://github.com/diPhantxm/ozon-api-client/actions/workflows/tests.yml/badge.svg)

[Ozon](https://ozon.ru) is a marketplace for small and medium enterprises to launch and grow their businesses in Russia.

Read full [documentation](https://docs.ozon.ru/api/seller/en/#tag/Introduction)

## How to start
### API
Get Client-Id and Api-Key in your seller profile [here](https://seller.ozon.ru/app/settings/api-keys?locale=en)

Just add dependency to your project and you're ready to go.
```bash
go get github.com/diphantxm/ozon-api-client
```
A simple example on how to use this library:
```Golang
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/diphantxm/ozon-api-client/ozon"
)

func main() {
	// Create a client with your Client-Id and Api-Key
	// [Documentation]: https://docs.ozon.ru/api/seller/en/#tag/Auth
	opts := []ozon.ClientOption{
		ozon.WithAPIKey("api-key"),
		ozon.WithClientId("client-id"),
	}
	c := ozon.NewClient(opts...)

	// Send request with parameters
	resp, err := client.Products().GetProductDetails(context.Background(), &ozon.GetProductDetailsParams{
		ProductId: 123456789,
	})
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Fatalf("error when getting product details: %s", err)
	}

	// Do some stuff
	for _, d := range resp.Result.Barcodes {
		fmt.Printf("Barcode %s\n", d)
	}
}
```

### Notifications
Ozon can send push-notifications to your REST server. There is an implementation of REST server that handles notifications in this library.

[Official documentation](https://docs.ozon.ru/api/seller/en/#tag/push_intro)

How to use:
```Golang
package main

import (
	"log"

	"github.com/diphantxm/ozon-api-client/ozon/notifications"
)

func main() {
	// Create server
	port := 5000
	server := notifications.NewNotificationServer(port)

	// Register handlers passing message type and handler itself
	server.Register(notifications.ChatClosedType, func(req interface{}) error {
		notification := req.(*notifications.ChatClosed)

		// Do something with the notification here...
		log.Printf("chat %s has been closed\n", notification.ChatId)

		return nil
	})

	// Run server
	if err := server.Run(); err != nil {
		log.Printf("error while running notification server: %s", err)
	}
}
```
