package client

import (
	"context"
	"exc8/pb"
	"fmt"
	//"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

// GrpcClient wraps the generated gRPC client
type GrpcClient struct {
	client pb.OrderServiceClient
}

// Creates and connects a new client to localhost:4000
func NewGrpcClient() (*GrpcClient, error) {
	conn, err := grpc.NewClient(":4000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewOrderServiceClient(conn)
	return &GrpcClient{client: client}, nil
}

func (c *GrpcClient) Run() error {
	ctx := context.Background()
	// todo
	// 1. List drinks
	fmt.Println("Requesting drinks 🍹🍺☕")
	drinksResponse, err := c.client.GetDrinks(ctx, &emptypb.Empty{})
	if err != nil {
		return fmt.Errorf("Error listing drinks: %w", err)
	}

	// Create a mp of drink IDs to names for later
	drinksMap := make(map[int32]string)
	fmt.Println("Available drinks:")
	for _, drink := range drinksResponse.GetDrinks() {
		drinksMap[drink.GetId()] = drink.GetName()

		if drink.GetPrice() == 0 {
			fmt.Printf("\t> id:%d  name:\"%s\"  description:\"%s\"\n",
				drink.GetId(), drink.GetName(), drink.GetDescription())
		} else {
			fmt.Printf("\t> id:%d  name:\"%s\"  price:%v  description:\"%s\"\n",
				drink.GetId(), drink.GetName(), drink.GetPrice(), drink.GetDescription())
		}
	}

	// 2. Order a few drinks
	fmt.Println("Ordering drinks 👨‍🍳⏱️🍻🍻")
	round1 := map[int32]int32{
		1: 2,
		2: 2,
		3: 2,
	}
	for id, qty := range round1 {
		fmt.Printf("\t> Ordering: %d x %s\n", qty, drinksMap[id])
		c.client.OrderDrink(ctx, &pb.Order{DrinkId: id, Quantity: qty})
	}

	// 3. Order more drinks
	fmt.Println("Ordering another round of drinks 👨‍🍳⏱️🍻🍻")
	round2 := map[int32]int32{
		1: 6,
		2: 6,
		3: 6,
	}
	for id, qty := range round2 {
		fmt.Printf("\t> Ordering: %d x %s\n", qty, drinksMap[id])
		c.client.OrderDrink(ctx, &pb.Order{DrinkId: id, Quantity: qty})
	}

	// 4. Get order total
	fmt.Println("Getting the bill 💹💹💹")
	summary, err := c.client.GetOrders(ctx, &emptypb.Empty{})
	if err != nil {
		return err
	}

	for _, order := range summary.GetOrders() {
		fmt.Printf("\t> Total: %d x %s\n", order.GetQuantity(), drinksMap[order.GetDrinkId()])
	}

	return nil
}
