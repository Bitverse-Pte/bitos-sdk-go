package client

import (
	"errors"

	"github.com/BitOS-labs/bitos/app"

	"github.com/BitOS-labs/bitos-sdk-go/common"
	"github.com/BitOS-labs/bitos-sdk-go/grpc"
	"github.com/BitOS-labs/bitos-sdk-go/types"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tharsis/ethermint/crypto/hd"
	"github.com/tharsis/ethermint/encoding"
)

func init() {
	sdk.GetConfig().SetBech32PrefixForAccount("bitos", "bitospub")
	sdk.GetConfig().SetBech32PrefixForValidator("bitosvaloper", "bitosvaloperpub")
	sdk.GetConfig().SetBech32PrefixForConsensusNode("bitosvalcons", "bitosvalconspub")
	sdk.GetConfig().SetCoinType(60)
	sdk.GetConfig().SetPurpose(44)
	sdk.GetConfig().Seal()
}

type BitosClient struct {
	grpc.GClient
	ctx sdkclient.Context

	accountRetriever *types.AccountRetriever
}

func NewClient(url string, chainId string) (*BitosClient, error) {
	if len(url) == 0 {
		return nil, errors.New("url can not be empty")
	}
	if len(chainId) == 0 {
		return nil, errors.New("chainId can not be empty")
	}
	encodingConfig := encoding.MakeConfig(app.ModuleBasics)
	grpcClient, err := grpc.NewGRPCClient(url)
	if err != nil {
		return nil, err
	}

	accountCache := common.NewCache(1000, true)
	ctx := sdkclient.Context{}.
		WithCodec(encodingConfig.Marshaler).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithBroadcastMode(flags.BroadcastSync).
		WithKeyringOptions(hd.EthSecp256k1Option()).
		WithChainID(chainId)

	return &BitosClient{
		ctx:              ctx,
		GClient:          grpcClient,
		accountRetriever: &types.AccountRetriever{QueryClient: grpcClient, Cache: accountCache},
	}, nil
}

func (client *BitosClient) WithChainId(chainId string) *BitosClient {
	client.ctx = client.ctx.WithChainID(chainId)
	return client
}

func (client *BitosClient) WithKeyring(k keyring.Keyring) *BitosClient {
	client.ctx = client.ctx.WithKeyring(k)
	return client
}

func (client *BitosClient) WithBroadcastMode(mode string) *BitosClient {
	client.ctx = client.ctx.WithBroadcastMode(mode)
	return client
}

func (client *BitosClient) DisableCache() {
	client.accountRetriever.Cache.Disable()
}

func (client *BitosClient) EnableCache() {
	client.accountRetriever.Cache.Enable()
}

func (client *BitosClient) ImportKey(name, armor, passphrase string) error {
	if client.ctx.Keyring == nil {
		return errors.New("no keyring found, please add keyring first")
	}
	return client.ctx.Keyring.ImportPrivKey(name, armor, passphrase)
}

func (client *BitosClient) ImportMnemonic(name, mnemonic string) error {
	if client.ctx.Keyring == nil {
		return errors.New("no keyring found, please add keyring first")
	}

	_, err := client.ctx.Keyring.NewAccount(name, mnemonic, "", sdk.GetConfig().GetFullBIP44Path(), hd.EthSecp256k1)
	return err
}

func (client *BitosClient) GetCtx() sdkclient.Context {
	return client.ctx
}
