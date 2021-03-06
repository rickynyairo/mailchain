package ethereum

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func NewRPC(address string) (*Client, error) {
	client, err := ethclient.Dial(address)
	if err != nil {
		return nil, err
	}

	return &Client{client: client}, nil
}

type Client struct {
	client *ethclient.Client
}

func (c *Client) BlockByNumber(ctx context.Context, blockNo uint64) (blk interface{}, err error) {
	return c.client.BlockByNumber(ctx, big.NewInt(int64(blockNo)))
}

func (c *Client) NetworkID(ctx context.Context) (*big.Int, error) {
	return c.client.NetworkID(ctx)
}

func (c *Client) LatestBlockNumber(ctx context.Context) (blockNo uint64, err error) {
	hdr, err := c.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, err
	}

	return hdr.Number.Uint64(), nil
}

func (c *Client) GetLatest(ctx context.Context) (blk interface{}, err error) {
	return c.client.BlockByNumber(ctx, nil)
}
