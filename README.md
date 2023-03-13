# Ozon API Client
A Ozon API client written in Golang

[Ozon](https://ozon.ru) is a marketplace for small and medium enterprises to launch and grow their businesses in Russia.

Read full [documentation](https://docs.ozon.ru/api/seller/en/#tag/Introduction)

## How to start
Just add dependency to your project and you're ready to go.
```bash
go get github.com/diphantxm/ozon-api-client
```
A simple example on how to use this library:
```Golang
package integrations

import (
	"fmt"
	"log"
	"net/http"

	"github.com/diphantxm/ozon-api-client/ozon"
)

func t() {
	// Create a client with your Client-Id and Api-Key
	// [Documentation]: https://docs.ozon.ru/api/seller/en/#tag/Auth
	client := ozon.NewClient("my-client-id", "my-api-key")

	// Send request with parameters
	resp, err := client.GetProductDetails(&ozon.GetProductDetailsParams{
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