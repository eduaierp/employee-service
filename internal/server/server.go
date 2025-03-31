package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"employee-service/internal/db"
	"employee-service/proto"
	"syscall"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedEmployeeServiceServer
}

// ✅ Create Employee
func (s *server) CreateEmployee(ctx context.Context, req *proto.Employee) (*proto.EmployeeResponse, error) {
	log.Printf("Received new employee: %+v", req)

	err := db.CreateEmployee(req)
	if err != nil {
		return nil, fmt.Errorf("could not create employee: %v", err)
	}

	message := fmt.Sprintf("Employee %s created successfully!", req.FullName)
	return &proto.EmployeeResponse{Message: message}, nil
}

// ✅ Get All Employees
func (s *server) GetAllEmployees(ctx context.Context, req *proto.EmptyRequest) (*proto.EmployeeList, error) {
	employees, err := db.GetAllEmployees()
	if err != nil {
		return nil, fmt.Errorf("could not fetch employees: %v", err)
	}

	return &proto.EmployeeList{Employees: employees}, nil
}

// ✅ Get Employee by ID
func (s *server) GetEmployee(ctx context.Context, req *proto.EmployeeRequest) (*proto.Employee, error) {
	employee, err := db.GetEmployeeByID(req.GetEmployeeId())
	if err != nil {
		return nil, fmt.Errorf("could not fetch employee: %v", err)
	}

	return employee, nil
}

// ✅ Update Employee
func (s *server) UpdateEmployee(ctx context.Context, req *proto.Employee) (*proto.EmployeeResponse, error) {
	err := db.UpdateEmployee(req.EmployeeId, req)
	if err != nil {
		return nil, fmt.Errorf("could not update employee: %v", err)
	}

	message := fmt.Sprintf("Employee %s updated successfully!", req.FullName)
	return &proto.EmployeeResponse{Message: message}, nil
}

// ✅ Delete Employee
func (s *server) DeleteEmployee(ctx context.Context, req *proto.EmployeeRequest) (*proto.EmployeeResponse, error) {
	err := db.DeleteEmployee(req.GetEmployeeId())
	if err != nil {
		return nil, fmt.Errorf("could not delete employee: %v", err)
	}

	message := fmt.Sprintf("Employee %s deleted successfully!", req.EmployeeId)
	return &proto.EmployeeResponse{Message: message}, nil
}

// StartServer function
func StartServer() {
	port := ":50052" // Different port from student service
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterEmployeeServiceServer(s, &server{})

	go func() {
		log.Printf("gRPC Employee service listening on %v", port)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down Employee service...")
	s.GracefulStop()
}
