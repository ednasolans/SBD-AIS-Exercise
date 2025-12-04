package server

import (
	"context"
	"exc8/pb"
	"net"
	//"fmt"
	//"log/slog"
	//"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

type GRPCService struct {
	pb.UnimplementedOrderServiceServer
	drinks map[int32]*pb.Drink // Drink menu
	orders map[int32]int32     // Accumulated orders
}

// Create a new gRPC server instance
func NewGRPCService() *GRPCService {
	srv := &GRPCService{
		drinks: make(map[int32]*pb.Drink),
		orders: make(map[int32]int32),
	}

	srv.drinks[1] = &pb.Drink{Id: 1, Name: "Spritzer", Price: 2, Description: "Wine with soda"}
	srv.drinks[2] = &pb.Drink{Id: 2, Name: "Beer", Price: 3, Description: "Hagenbergner Gold"}
	srv.drinks[3] = &pb.Drink{Id: 3, Name: "Coffee", Price: 0, Description: "Mifare isn't that secure"}

	return srv
}

func StartGrpcServer() error {
	// Create a new gRPC server.
	srv := grpc.NewServer()
	// Create grpc service
	grpcService := NewGRPCService()
	// Register our service implementation with the gRPC server.
	pb.RegisterOrderServiceServer(srv, grpcService)
	// Serve gRPC server on port 4000.
	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		return err
	}
	err = srv.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}

// todo implement functions
// Return all drinks on the menu
func (s *GRPCService) GetDrinks(ctx context.Context, in *emptypb.Empty) (*pb.DrinkList, error) {
	list := &pb.DrinkList{}
	for _, drink := range s.drinks {
		list.Drinks = append(list.Drinks, drink)
	}
	return list, nil
}

// Stores a new order
func (s *GRPCService) OrderDrink(ctx context.Context, order *pb.Order) (*wrapperspb.BoolValue, error) {
	_, exists := s.drinks[order.DrinkId]
	if !exists {
		return wrapperspb.Bool(false), status.Errorf(codes.NotFound, "Drink ID %d not found!", order.DrinkId)
	}
	if order.Quantity <= 0 {
		return wrapperspb.Bool(false), status.Error(codes.InvalidArgument, "Quantity must be greater than 0!")
	}

	s.orders[order.DrinkId] += order.Quantity

	return wrapperspb.Bool(true), nil
}

// Returns all accumulated orders
func (s *GRPCService) GetOrders(ctx context.Context, in *emptypb.Empty) (*pb.Orders, error) {
	result := &pb.Orders{}
	for id, quantity := range s.orders {
		result.Orders = append(result.Orders, &pb.Order{
			DrinkId:  id,
			Quantity: quantity,
		})
	}
	return result, nil
}
