package server

import (
	"context"
	"fmt"

	table_pb "github.com/hapiio/hapiio/proto/table/v1"
)

// CreateTable implements v1.TableServiceServer
func (s *Server) CreateTable(ctx context.Context, req *table_pb.CreateTableRequest) (*table_pb.CreateTableResponse, error) {
	s.m.Lock()
	defer s.m.Unlock()

	tableName := req.Table.GetTableName()
	fields := req.Table.GetFields()

	columns := ""
	for i, field := range fields {
		if i > 0 {
			columns += ","
		}
		columns += field.GetName() + " " + field.GetType()
	}
	response := table_pb.Table{TableName: tableName, Fields: fields}
	err := s.db.Create(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to create table: %v", err)
	}

	return &table_pb.CreateTableResponse{
		Table: &response,
	}, nil
}
