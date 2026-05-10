/*
package handler

import (

	"fmt"
	"task_ex/internal/service"
	"task_ex/service/pb"

)

	type GoldPrice struct{
		pb.UnimplementedTaskServiceServer
		Goldservice *service.GetGoldPrice
	}

	func newGoldPriceHandler (Goldservice *service.GetGoldPrice) *GoldPrice{
		fmt.Println("f")
		return &GoldPrice{Goldservice: Goldservice}

}
*/
package handler

import (
	"task_ex/internal/service"
	"task_ex/service/pb"
)

type GoldPrice struct {
	pb.UnimplementedTaskServiceServer
	Goldservice service.GetGoldPriceFunc
}

func newGoldPriceHandler(gs service.GetGoldPriceFunc) *GoldPrice {
	return &GoldPrice{Goldservice: gs}
}
