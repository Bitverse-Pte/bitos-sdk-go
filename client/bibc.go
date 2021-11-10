package client

import (
	"github.com/cosmos/cosmos-sdk/types/tx"

	clienttypes "github.com/BitOS-labs/bitos/x/bibc/core/client/types"
	packettypes "github.com/BitOS-labs/bitos/x/bibc/core/packet/types"
)

func (client *BitosClient) UpdateClient(msg clienttypes.MsgUpdateClient, options ...Option) (*tx.BroadcastTxResponse, error) {
	txf, err := Prepare(client, msg.GetSigners()[0], &msg, options...)
	if err != nil {
		return nil, err
	}
	return client.Broadcast(txf, &msg)
}

func (client *BitosClient) RecvPacket(msg packettypes.MsgRecvPacket, options ...Option) (*tx.BroadcastTxResponse, error) {
	txf, err := Prepare(client, msg.GetSigners()[0], &msg, options...)
	if err != nil {
		return nil, err
	}
	return client.Broadcast(txf, &msg)
}

func (client *BitosClient) Acknowledgement(msg packettypes.MsgAcknowledgement, options ...Option) (*tx.BroadcastTxResponse, error) {
	txf, err := Prepare(client, msg.GetSigners()[0], &msg, options...)
	if err != nil {
		return nil, err
	}
	return client.Broadcast(txf, &msg)
}

func (client *BitosClient) CleanPacket(msg packettypes.MsgCleanPacket, options ...Option) (*tx.BroadcastTxResponse, error) {
	txf, err := Prepare(client, msg.GetSigners()[0], &msg, options...)
	if err != nil {
		return nil, err
	}
	return client.Broadcast(txf, &msg)
}

func (client *BitosClient) RecvCleanPacket(msg packettypes.MsgRecvCleanPacket, options ...Option) (*tx.BroadcastTxResponse, error) {
	txf, err := Prepare(client, msg.GetSigners()[0], &msg, options...)
	if err != nil {
		return nil, err
	}
	return client.Broadcast(txf, &msg)
}
