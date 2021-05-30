package main

import (
	"hex/internal/adapters/core/app/api"
	"hex/internal/adapters/core/arithmetic"
	"hex/internal/adapters/framework/right/db"
	"hex/internal/ports"
	"log"
	"os"

	gRPC "hex/internal/adapters/framework/left/grpc"
)

func main() {
	var err error
	
	// ports
	var dbAdapter ports.DbPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort


	dbDriver := os.Getenv("DB_DRIVER");
	dbName := os.Getenv("DB_NAME")

	dbAdapter, err = db.NewDbAdapter(dbDriver, dbName);

	if err != nil {
		log.Fatalf("failed to initiate database connection %v", err)
	}
	
	defer dbAdapter.CloseDbConnection();

	core = arithmetic.NewArithmeticAdapter();

	appAdapter = api.NewApiAdapter(dbAdapter, core);

	gRPCAdapter = gRPC.NewServerAdapter(appAdapter);
	
	gRPCAdapter.Run();
	
}