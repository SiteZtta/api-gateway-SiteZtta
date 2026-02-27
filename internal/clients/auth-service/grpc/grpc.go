package grpc

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	sitezttav2 "github.com/SiteZtta/protos-SiteZtta/gen/go/auth"
	grpclog "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	api sitezttav2.AuthServiceClient
	log *slog.Logger
}

func New(addr string,
	log *slog.Logger,
	timeout time.Duration,
	retriesCount int,
) (*Client, error) {
	const fn = "api-gateway-SiteZtta.internal.clients.auth-service.grpc.grpc"
	retryOpts := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(retriesCount)),
		grpcretry.WithPerRetryTimeout(timeout),
	}
	lopOpts := []grpclog.Option{
		grpclog.WithLogOnEvents(grpclog.PayloadReceived, grpclog.PayloadSent),
	}

	cc, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			grpclog.UnaryClientInterceptor(InterceptorLogger(log), lopOpts...),
			grpcretry.UnaryClientInterceptor(retryOpts...),
		))
	if err != nil {
		return nil, fmt.Errorf("%s:%w", fn, err)
	}
	return &Client{
		api: sitezttav2.NewAuthServiceClient(cc),
		log: log,
	}, nil
}

func (c *Client) CreateUser(ctx context.Context, req *sitezttav2.SignUpRequest) (*sitezttav2.UserIdResponse, error) {
	const fn = "api-gateway-SiteZtta.internal.clients.auth-service.grpc.createUser"

	resp, err := c.api.CreateUser(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%s:%w", fn, err)
	}

	return resp, nil
}

func (c *Client) GenerateToken(ctx context.Context, req *sitezttav2.SignInRequest) (*sitezttav2.TokenResponse, error) {
	const fn = "api-gateway-SiteZtta.internal.clients.auth-service.grpc.generateToken"

	resp, err := c.api.GenerateToken(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%s:%w", fn, err)
	}

	return resp, nil
}

func (c *Client) ValidateToken(ctx context.Context, req *sitezttav2.TokenRequest) (*sitezttav2.AuthInfo, error) {
	const fn = "api-gateway-SiteZtta.internal.clients.auth-service.grpc.validateToken"

	resp, err := c.api.ValidateToken(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%s:%w", fn, err)
	}

	return resp, nil
}

// InterceptorLogger is a logger for grpc interceptors.
// This code is simple enough to be copied and not imported.
func InterceptorLogger(log *slog.Logger) grpclog.Logger {
	return grpclog.LoggerFunc(func(ctx context.Context, lvl grpclog.Level, msg string, fields ...any) {
		log.Log(ctx, slog.Level(lvl), msg, fields...)
	})

}
