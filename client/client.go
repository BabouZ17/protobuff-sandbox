package main

import (
	"context"
	"log"

	"github.com/BabouZ17/protobuff-sandbox/services"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not start dialing on :9000")
	}
	defer conn.Close()

	client := services.NewRecordServiceClient(conn)

	records := []services.RecordRequest{
		{
			Id:        uuid.New().String(),
			SensorId:  "1",
			Value:     12.5,
			CreatedAt: timestamppb.Now(),
		},
		{
			Id:        uuid.New().String(),
			SensorId:  "2",
			Value:     14.4,
			CreatedAt: timestamppb.Now(),
		},
		{
			Id:        uuid.New().String(),
			SensorId:  "1",
			Value:     10.0,
			CreatedAt: timestamppb.Now(),
		},
	}

	for _, record := range records {
		result, err := client.SaveRecord(context.Background(), &record)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println(result)
		}
	}
}
