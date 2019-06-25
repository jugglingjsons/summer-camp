package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/jadczakd/summer-camp/summer_app/pkg/service"
)

// CreateRequest collects the request parameters for the Create method.
type CreateRequest struct {
	Payload string `json:"payload"`
}

// CreateResponse collects the response parameters for the Create method.
type CreateResponse struct {
	Response string `json:"response"`
	Err      error  `json:"err"`
}

// MakeCreateEndpoint returns an endpoint that invokes Create on the service.
func MakeCreateEndpoint(s service.SummerAppService) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		response, err := s.Create(ctx, req.Payload)
		return CreateResponse{Response: response, Err: err}, nil
	}
}

// Failed implements Failer.
func (r CreateResponse) Failed() error {
	return r.Err
}

// HealthRequest collects the request parameters for the Health method.
type HealthRequest struct{}

// HealthResponse collects the response parameters for the Health method.
type HealthResponse struct {
	Ok bool `json:"ok"`
}

// MakeHealthEndpoint returns an endpoint that invokes Health on the service.
func MakeHealthEndpoint(s service.SummerAppService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		ok := s.Health(ctx)
		return HealthResponse{
			Ok: ok,
		}, nil
	}
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Create implements Service. Primarily useful in a client.
func (e Endpoints) Create(ctx context.Context, payload string) (err error) {
	request := CreateRequest{Payload: payload}
	response, err := e.CreateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateResponse).Err
}

// Health implements Service. Primarily useful in a client.
func (e Endpoints) Health(ctx context.Context) (ok bool) {
	request := HealthRequest{}
	response, err := e.HealthEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(HealthResponse).Ok
}
