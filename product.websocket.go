// package product

// import (
// 	"WebServices/database"
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"
// 	// "golang.org/x/net/websocket"
// )

// type message struct {
// 	Data string `json:"data"`
// 	Type string `json:"type"`
// }

// func productSocket(ws *websocket.Conn) {
// 	go func(c *websocket.Conn) {
// 		for {
// 			var msg message
// 			if err := websocket.JSON.Receive(ws, &msg); err != nil {
// 				log.Println(err)
// 				break
// 			}
// 			fmt.Println("Received message %s\n", msg.Data)
// 		}
// 	}(ws)
// 	for {
// 		products, err := GetTopTenProducts()
// 		if err != nil {
// 			log.Println(err)
// 			break
// 		}
// 		if err := websocket.JSON.Send(ws, products); err != nil {
// 			log.Println(err)
// 			break
// 		}
// 		time.Sleep(10 * time.Second)
// 	}
// }
// func GetTopTenProducts() ([]Product, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
// 	defer cancel()
// 	result, err := database.DbConn.QueryContext(ctx, `SELECT productId,manufacturer,sku,upc,pricePerUnit,quantityOnHand,productName FROM products ORDER BY quantityOnHand DESC LIMIT 10`)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer result.Close()
// 	products := make([]Product, 0)
// 	for result.Next() {
// 		var product Product
// 		result.Scan(&product.ProductID, &product.Manufacturer, &product.Sku, &product.Upc, &product.PricePerUnit, &product.QuantityOnHand, &product.ProductName)
// 		products = append(products, product)
// 	}
// 	return products, nil
// }
