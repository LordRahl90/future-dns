package servers

import (
	"context"
	"dns/domains/maths"
	"dns/proto"
	"dns/requests"
)

// DNSGRPCServer grpc serer
type DNSGRPCServer struct {
	proto.DNSServer
	ms maths.IMathService
}

// NewGRPCServer returns a new instance of grpc server
func NewGRPCServer(ms maths.IMathService) *DNSGRPCServer {
	return &DNSGRPCServer{
		ms: ms,
	}
}

// Calculate prform the calculation
func (dns *DNSGRPCServer) Calculate(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	svcReq := &requests.Request{
		CoordX:   req.CoordX,
		CoordY:   req.CoordY,
		CoordZ:   req.CoordZ,
		Velocity: req.Velocity,
	}
	domainReq, err := svcReq.ToDomainRequest()
	if err != nil {
		return nil, err
	}
	result := dns.ms.Calculate(ctx, domainReq)
	return &proto.Response{
		Location: float32(result),
	}, nil
}
