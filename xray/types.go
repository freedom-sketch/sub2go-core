package xray

import (
	"github.com/xtls/xray-core/app/proxyman/command"
	statsService "github.com/xtls/xray-core/app/stats/command"
	"google.golang.org/grpc"
)

// XrayAPI is a gRPC client for managing Xray core configuration, inbounds, outbounds, and statistics
type XrayAPI struct {
	HandlerServiceClient *command.HandlerServiceClient
	StatsServiceClient   *statsService.StatsServiceClient
	grpcClient           *grpc.ClientConn
	isConnected          bool
}

type VlessUser struct {
	ID    string
	Email string
	Flow  string
}

// Traffic represents network traffic statistics for Xray connections.
// It tracks upload and download bytes for inbound or outbound traffic.
type Traffic struct {
	IsInbound  bool
	IsOutbound bool
	Tag        string
	Up         int64
	Down       int64
}

type ClientTraffic struct {
	Email string
	Up    int64
	Down  int64
}
