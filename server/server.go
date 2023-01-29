package server

import (
	"sync"

	"github.com/hapiio/hapiio/core"
	"github.com/hapiio/hapiio/db"
	"gorm.io/gorm"
)

type Server struct {
	m  *sync.RWMutex
	db *gorm.DB
	// tables []*table_pb.Table
}

func New() *Server {
	conf := core.Configure()
	return &Server{
		m:  &sync.RWMutex{},
		db: db.Connect(conf),
	}
}
