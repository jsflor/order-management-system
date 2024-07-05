package main

import "context"

type store struct {
	// mongo db instance
}

func NewStore() *store {
	return &store{}
}

func (s *store) Create(ctx context.Context) error {
	return nil
}
