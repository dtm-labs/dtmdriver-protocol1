package sample

import (
	"google.golang.org/grpc/resolver"
)

type nopResolver struct {
	cc resolver.ClientConn
}

func (r *nopResolver) Close() {
}

func (r *nopResolver) ResolveNow(options resolver.ResolveNowOptions) {
}

type sampleBuilder struct{}

func (s *sampleBuilder) Scheme() string { return "protocol1" }

func (s *sampleBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {

	addrs := []resolver.Address{{Addr: target.Endpoint}}

	if err := cc.UpdateState(resolver.State{Addresses: addrs}); err != nil {
		return nil, err
	}
	return &nopResolver{cc: cc}, nil
}
