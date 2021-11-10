package client

import (
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
)

func (client *BitosClient) SubmitProposal(msg types.MsgSubmitProposal, options ...Option) (*tx.BroadcastTxResponse, error) {
	txf, err := Prepare(client, msg.GetSigners()[0], &msg, options...)
	if err != nil {
		return nil, err
	}
	return client.Broadcast(txf, &msg)
}

func (client *BitosClient) Deposit(msg types.MsgDeposit, options ...Option) (*tx.BroadcastTxResponse, error) {
	txf, err := Prepare(client, msg.GetSigners()[0], &msg, options...)
	if err != nil {
		return nil, err
	}
	return client.Broadcast(txf, &msg)
}

func (client *BitosClient) Vote(msg types.MsgVote, options ...Option) (*tx.BroadcastTxResponse, error) {
	txf, err := Prepare(client, msg.GetSigners()[0], &msg, options...)
	if err != nil {
		return nil, err
	}
	return client.Broadcast(txf, &msg)
}

func (client *BitosClient) VoteWeighted(msg types.MsgVoteWeighted, options ...Option) (*tx.BroadcastTxResponse, error) {
	txf, err := Prepare(client, msg.GetSigners()[0], &msg, options...)
	if err != nil {
		return nil, err
	}
	return client.Broadcast(txf, &msg)
}
