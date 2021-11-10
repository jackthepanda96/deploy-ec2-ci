package controllers

import (
	"fmt"
	"net/http"
	configs "project/mock_api/config"

	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

var s snap.Client

// func SetUp() {
// 	midtrans.ServerKey = configs.SERVER_KEY
// 	midtrans.Environment = midtrans.Sandbox
// }

func TestReq() {
	s.New(configs.SERVER_KEY, midtrans.Sandbox)
	s.Options.SetPaymentOverrideNotification("http://localhost:8000/ret")
	s.Options.SetPaymentAppendNotification("http://localhost:8000/ret")
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "YOUR-ORDER-ID-99",
			GrossAmt: 100000,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
	}

	// 3. Request create Snap transaction to Midtrans
	snapResp, _ := s.CreateTransaction(req)
	fmt.Println("Response :", snapResp)
}

func PaymentStatus(ec echo.Context) error {
	fmt.Println(ec.QueryParam("order_id"), ec.QueryParam("status_code"))

	return ec.JSON(http.StatusCreated, "OK")
}

// func TestCoreAPI() {
// 	c := coreapi.Client{}
// 	c.New(configs.SERVER_KEY, midtrans.Sandbox)

// 	chargeReq := &coreapi.ChargeReq{
// 		PaymentType: coreapi.PaymentTypeBCAKlikpay,
// 		TransactionDetails: midtrans.TransactionDetails{
// 			OrderID:  "123456",
// 			GrossAmt: 200000,
// 		},
// 		Items: &[]midtrans.ItemDetails{
// 			midtrans.ItemDetails{
// 				ID:    "ITEM1",
// 				Price: 200000,
// 				Qty:   1,
// 				Name:  "Someitem",
// 			},
// 		},
// 	}
// }
