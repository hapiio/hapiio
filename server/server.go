package server

import (
	"context"
	"sync"

	table "github.com/hapiio/hapiio/proto/table/v1"
)

type Server struct {
	m *sync.RWMutex
	// tables []*table.Table
}

func New() *Server {
	return &Server{
		m: &sync.RWMutex{},
	}
}

// CreateTable implements v1.TableServiceServer
func (s *Server) CreateTable(context.Context, *table.CreateTableRequest) (*table.CreateTableResponse, error) {
	s.m.Lock()
	defer s.m.Unlock()

	t := &table.Table{
		TableName: "Test",
		Column:    make([]*table.Column, 0),
	}

	return &table.CreateTableResponse{
		Table: t,
	}, nil
}
