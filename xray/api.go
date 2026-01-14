package xray

import (
	"context"
	"encoding/json"
	"fmt"
	"math"

	"github.com/xtls/xray-core/app/proxyman/command"
	statsService "github.com/xtls/xray-core/app/stats/command"
	"github.com/xtls/xray-core/common/protocol"
	"github.com/xtls/xray-core/common/serial"
	"github.com/xtls/xray-core/infra/conf"
	"github.com/xtls/xray-core/proxy/vless"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

// Init connects to the Xray API server and initializes handler and stats service clients
func (x *XrayAPI) Init(apiPort int) error {
	if apiPort <= 0 || apiPort > math.MaxUint16 {
		return fmt.Errorf("invalid Xray API port: %d", apiPort)
	}

	addr := fmt.Sprintf("127.0.0.1:%d", apiPort)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to Xray API: %w", err)
	}
	x.grpcClient = conn
	x.isConnected = true

	hsClient := command.NewHandlerServiceClient(conn)
	ssClient := statsService.NewStatsServiceClient(conn)
	x.HandlerServiceClient = &hsClient
	x.StatsServiceClient = &ssClient
	return nil
}

// Close closes the gRPC connection and resets the XrayAPI client state
func (x *XrayAPI) Close() {
	if x.grpcClient != nil {
		x.grpcClient.Close()
	}
	x.HandlerServiceClient = nil
	x.StatsServiceClient = nil
	x.isConnected = false
}

// AddInbound adds a new inbound configuration to the Xray core via gRPC
func (x *XrayAPI) AddInbound(inbound []byte) error {
	conf := new(conf.InboundDetourConfig)
	err := json.Unmarshal(inbound, conf)
	if err != nil {
		return err
	}
	config, err := conf.Build()
	if err != nil {
		return err
	}
	inboundConfig := command.AddInboundRequest{Inbound: config}
	_, err = (*x.HandlerServiceClient).AddInbound(context.Background(), &inboundConfig)
	return err
}

func (x *XrayAPI) AddUser(inboundTag string, user *VlessUser) error {
	var account *serial.TypedMessage
	account = serial.ToTypedMessage(&vless.Account{
		Id:   user.ID,
		Flow: user.Flow,
	})

	client := *x.HandlerServiceClient
	_, err := client.AlterInbound(context.Background(), &command.AlterInboundRequest{
		Tag: inboundTag,
		Operation: serial.ToTypedMessage(&command.AddUserOperation{
			User: &protocol.User{
				Email:   user.Email,
				Account: account,
			},
		}),
	})
	return err
}

func (x *XrayAPI) RemoveUser(inboundTag, email string) error {
	op := &command.RemoveUserOperation{Email: email}
	req := &command.AlterInboundRequest{
		Tag:       inboundTag,
		Operation: serial.ToTypedMessage(op),
	}
	_, err := (*x.HandlerServiceClient).AlterInbound(context.Background(), req)
	return err
}
