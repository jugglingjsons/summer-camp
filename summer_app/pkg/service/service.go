package service

import (
	"context"
	"fmt"
	"sync"
)

// SummerAppService describes the service.
type SummerAppService interface {
	// Add your methods here
	Create(ctx context.Context, payload string) (response string, err error)
	Health(ctx context.Context) (ok bool)
}

type basicSummerAppService struct {
	mtx sync.RWMutex
	db  map[string]int
}

func NewBasicSummerAppService() SummerAppService {
	return &basicSummerAppService{
		db: map[string]int{},
	}
}
func (b *basicSummerAppService) Create(ctx context.Context, payload string) (response string, err error) {
	// TODO implement the business logic of Create
	b.mtx.Lock()
	defer b.mtx.Unlock()
	b.db[payload] = b.db[payload] + 1
	return fmt.Sprintf("Number of invocations for %s, is exactly %v", payload, b.db[payload]), err
}

func (b *basicSummerAppService) Health(ctx context.Context) (ok bool) {
	// TODO implement the business logic of Health
	return true
}

// New returns a SummerAppService with all of the expected middleware wired in.
func New(middleware []Middleware) SummerAppService {
	var svc SummerAppService = NewBasicSummerAppService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
