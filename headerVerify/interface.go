package headerVerify

import (
	"context"

	"github.com/ElrondNetwork/elrond-go/state"
	"github.com/ElrondNetwork/elrond-sdk-erdgo/data"
)

type Proxy interface {
	GetNetworkConfig(ctx context.Context) (*data.NetworkConfig, error)
	GetRatingsConfig(ctx context.Context) (*data.RatingsConfig, error)
	GetRawBlockByHash(ctx context.Context, shardId uint32, hash string) ([]byte, error)
	GetRawBlockByNonce(ctx context.Context, shardId uint32, nonce uint64) ([]byte, error)
	GetRawMiniBlockByHash(ctx context.Context, shardId uint32, hash string) ([]byte, error)
}

type EpochsConfigCacheHandler interface {
	Add(epoch uint32, validatorsInfo []*state.ShardValidatorInfo) error
	IsInCache(epoch uint32) bool
	IsInterfaceNil() bool
}
