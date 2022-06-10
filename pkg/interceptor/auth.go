package interceptor

import (
	"context"
	"github.com/ShatteredRealms/GoUtils/pkg/service"
	"github.com/allegro/bigcache/v3"
	"google.golang.org/grpc"
	"log"
	"time"
)

type UserAuthorizations struct {
}

type authInterceptor struct {
	// The JWT service to use for verifying JWTs
	jwtService service.JWTService

	// Maps a user role to an array of gRPC method strings it has access to
	rolePermissions map[string][]string

	// Maps a gRPC method string (ie- /package.service/method) to a permission needed to run.
	// If the gRPC method is not in the map, then it is public
	authorizedPermissions map[string]string

	// Maps a gRPC method string (ie- /package.service/method) to a list of authorized roles
	// If a gRPC method is not in the map, then no roles have direct access to it.
	authorizedRoles map[string][]string

	// Cache to search for user permissions
	userAuthorizationCache *bigcache.BigCache
}

func NewAuthInterceptor(
	jwtService service.JWTService,
	rolePermissions map[string][]string,
	authorizedPermissions map[string]string,
	authorizedRoles map[string][]string,
) *authInterceptor {
	cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	if err != nil {
		panic(err)
	}

	return &authInterceptor{
		jwtService:             jwtService,
		rolePermissions:        rolePermissions,
		authorizedPermissions:  authorizedPermissions,
		authorizedRoles:        authorizedRoles,
		userAuthorizationCache: cache,
	}
}

func (interceptor *authInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Println("--> unary interceptor: ", info.FullMethod)
		return handler(ctx, req)
	}
}

func (interceptor *authInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		log.Println("--> stream interceptor: ", info.FullMethod)
		return handler(srv, stream)
	}
}

func (interceptor *authInterceptor) authorize(ctx context.Context, method string) error {
	//permission, ok := interceptor.authorizedPermissions[method]
	//
	//if !ok {
	//    return nil
	//}

	return nil
}

func (interceptor *authInterceptor) hasRole(userid uint64, role string) bool {
	return true
}

func (interceptor *authInterceptor) hasPermission(userid uint64, method string) bool {
	return true
}
