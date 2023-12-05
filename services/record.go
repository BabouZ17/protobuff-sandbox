package services

import (
	"log"

	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	records []RecordResponse
	UnimplementedRecordServiceServer
}

func (s *Server) SaveRecord(ctx context.Context, record *RecordRequest) (*RecordResponse, error) {
	log.Printf("Received new record from sensor: %s at %s", record.GetSensorId(), record.GetCreatedAt())

	recordResponse := &RecordResponse{Record: record, SavedAt: timestamppb.Now()}
	s.records = append(s.records, *recordResponse)
	return recordResponse, nil
}

func (s *Server) ListRecords(listRecords *ListRecordsRequest, stream RecordService_ListRecordsServer) error {
	records := make([]RecordResponse, len(s.records))
	copy(records, s.records)

	limit := listRecords.GetLimit()
	if limit <= int32(len(records)) {
		records = records[:limit]
	}

	eg := new(errgroup.Group)
	for _, record := range records {
		localRecord := record
		eg.Go(func() error {
			if err := stream.Send(&localRecord); err != nil {
				return err
			}
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}
