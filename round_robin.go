package go_roundrobin

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"sync/atomic"
)

type RoundRobinDB interface {
	Next() *sqlx.DB
}

type roundRobinDB struct {
	databases []*sqlx.DB
	next      uint32
}

func New(databases []*sqlx.DB) (RoundRobinDB, error) {
	if len(databases) == 0 {
		return nil, errors.New("servers not found")
	}
	return &roundRobinDB{
		databases: databases,
	}, nil
}

func (r *roundRobinDB) Next() *sqlx.DB {
	next := atomic.AddUint32(&r.next, 1)
	return r.databases[(int(next)-1)%len(r.databases)]
}
