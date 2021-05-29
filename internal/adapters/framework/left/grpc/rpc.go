package grpc

import (
	"context"

	"hex/internal/adapters/framework/left/grpc/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (grpca ServerAdapter) GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error){
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0{
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grpca.api.GetAddition(req.A, req.B)

	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}

func (grpca ServerAdapter) GetSubstraction(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error){
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0{
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grpca.api.GetSubstraction(req.A, req.B)

	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}

func (grpca ServerAdapter) GetMultiplication(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error){
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0{
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grpca.api.GetMultiplication(req.A, req.B)

	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb.Answer{
		Value: answer,
	}

	return ans, nil
}


func (grpca ServerAdapter) GetDivision(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error){
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0{
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grpca.api.GetDivision(req.A, req.B)

	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}

	ans = &pb.Answer{
		Value: answer,
	}
	
	return ans, nil
}