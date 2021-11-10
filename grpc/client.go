package grpc

import (
	grpc1 "github.com/gogo/protobuf/grpc"
	"google.golang.org/grpc"

	abcitypes "github.com/BitOS-labs/bitos/grpc_abci/types"
	bibcclitypes "github.com/BitOS-labs/bitos/x/bibc/core/client/types"
	bibcpkttypes "github.com/BitOS-labs/bitos/x/bibc/core/packet/types"
	bibcruttypes "github.com/BitOS-labs/bitos/x/bibc/core/routing/types"

	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/types/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

type GClient struct {
	clientConn grpc1.ClientConn

	AuthQuery        authtypes.QueryClient
	BankQuery        banktypes.QueryClient
	GovQuery         govtypes.QueryClient
	BIBCClientQuery  bibcclitypes.QueryClient
	BIBCPacketQuery  bibcpkttypes.QueryClient
	BIBCRoutingQuery bibcruttypes.QueryClient
	ABCIQuery        abcitypes.ABCIQueryClient
	TMServiceQuery   tmservice.ServiceClient

	TxClient tx.ServiceClient
}

func NewGRPCClient(url string) (GClient, error) {
	dialOpts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	clientConn, err := grpc.Dial(url, dialOpts...)
	if err != nil {
		return GClient{}, err
	}
	return GClient{
		clientConn: clientConn,

		BIBCClientQuery:  bibcclitypes.NewQueryClient(clientConn),
		BIBCPacketQuery:  bibcpkttypes.NewQueryClient(clientConn),
		BIBCRoutingQuery: bibcruttypes.NewQueryClient(clientConn),
		ABCIQuery:        abcitypes.NewABCIQueryClient(clientConn),
		BankQuery:        banktypes.NewQueryClient(clientConn),
		AuthQuery:        authtypes.NewQueryClient(clientConn),
		TMServiceQuery:   tmservice.NewServiceClient(clientConn),

		TxClient: tx.NewServiceClient(clientConn),
	}, nil
}
