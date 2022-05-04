package postgresql

import "database/sql"

type WalkingSkeleStore struct {
	db *sql.DB
}

func NewWalkingSkeleStore(db *sql.DB) (*WalkingSkeleStore, error) {

	return &WalkingSkeleStore{db: db}, nil
}
