package api

import (
	"github.com/ava-labs/avalanchego/api/admin"
	"github.com/ava-labs/avalanchego/api/health"
	"github.com/ava-labs/avalanchego/api/info"
	"github.com/ava-labs/avalanchego/indexer"
	"github.com/ava-labs/avalanchego/vms/avm"
	"github.com/ava-labs/avalanchego/vms/platformvm"
	evmclient "github.com/ava-labs/coreth/plugin/evm/client"
)

// Issues API calls to a node
// TODO: byzantine api. check if appropriate. improve implementation.
type Client interface {
	PChainAPI() platformvm.Client
	XChainAPI() avm.Client
	XChainWalletAPI() avm.WalletClient
	CChainAPI() evmclient.Client
	CChainEthAPI() EthClient // ethclient websocket wrapper that adds mutexed calls, and lazy conn init (on first call)
	InfoAPI() info.Client
	HealthAPI() health.Client
	AdminAPI() admin.Client
	PChainIndexAPI() indexer.Client
	CChainIndexAPI() indexer.Client
	// TODO add methods
}
