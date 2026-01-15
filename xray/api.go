package xray

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"time"

	"regexp"

	"github.com/xtls/xray-core/app/proxyman/command"
	statsService "github.com/xtls/xray-core/app/stats/command"
	"github.com/xtls/xray-core/common/protocol"
	"github.com/xtls/xray-core/common/serial"
	"github.com/xtls/xray-core/infra/conf"
	"github.com/xtls/xray-core/proxy/vless"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

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

// AddUser adds a user to an inbound in the Xray core
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

// RemoveUser removes a user from an inbound in the Xray core by email
func (x *XrayAPI) RemoveUser(inboundTag, email string) error {
	op := &command.RemoveUserOperation{Email: email}
	req := &command.AlterInboundRequest{
		Tag:       inboundTag,
		Operation: serial.ToTypedMessage(op),
	}
	_, err := (*x.HandlerServiceClient).AlterInbound(context.Background(), req)
	return err
}

// GetTraffic queries traffic statistics from the Xray core, optionally resetting counters.
func (x *XrayAPI) GetTraffic(reset bool) ([]*Traffic, []*ClientTraffic, error) {
	if !x.isConnected {
		return nil, nil, fmt.Errorf("xray api not connected")
	}

	if x.StatsServiceClient == nil {
		return nil, nil, fmt.Errorf("stats service client not initialized")
	}

	trafficRegex := regexp.MustCompile(`(inbound|outbound)>>>([^>]+)>>>traffic>>>(downlink|uplink)`)
	clientTrafficRegex := regexp.MustCompile(`user>>>([^>]+)>>>traffic>>>(downlink|uplink)`)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := (*x.StatsServiceClient).QueryStats(ctx, &statsService.QueryStatsRequest{Reset_: reset})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to query stats: %w", err)
	}

	tagTrafficMap := make(map[string]*Traffic)
	emailTrafficMap := make(map[string]*ClientTraffic)

	for _, stat := range resp.GetStat() {
		name := stat.Name
		value := stat.Value

		if matches := trafficRegex.FindStringSubmatch(name); len(matches) == 4 {
			processTraffic(matches, value, tagTrafficMap)
		}

		if matches := clientTrafficRegex.FindStringSubmatch(name); len(matches) == 3 {
			processClientTraffic(matches, value, emailTrafficMap)
		}
	}

	return mapToSlice(tagTrafficMap), mapToSlice(emailTrafficMap), nil
}

// processTraffic aggregates a traffic stat into trafficMap using regex matches and value.
func processTraffic(matches []string, value int64, trafficMap map[string]*Traffic) {
	isInbound := matches[1] == "inbound"
	tag := matches[2]
	isDown := matches[3] == "downlink"

	if tag == "api" {
		return
	}

	traffic, ok := trafficMap[tag]
	if !ok {
		traffic = &Traffic{
			IsInbound:  isInbound,
			IsOutbound: !isInbound,
			Tag:        tag,
		}
		trafficMap[tag] = traffic
	}

	if isDown {
		traffic.Down = value
	} else {
		traffic.Up = value
	}
}

// processClientTraffic updates clientTrafficMap with upload/download values for a client email.
func processClientTraffic(matches []string, value int64, clientTrafficMap map[string]*ClientTraffic) {
	email := matches[1]
	isDown := matches[2] == "downlink"

	traffic, ok := clientTrafficMap[email]
	if !ok {
		traffic = &ClientTraffic{Email: email}
		clientTrafficMap[email] = traffic
	}

	if isDown {
		traffic.Down = value
	} else {
		traffic.Up = value
	}
}

// mapToSlice converts a map of pointers to a slice of pointers.
func mapToSlice[T any](m map[string]*T) []*T {
	result := make([]*T, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}
