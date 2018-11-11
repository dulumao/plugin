package executor

import (
	"gitlab.33.cn/chain33/chain33/types"
	pty "gitlab.33.cn/chain33/plugin/plugin/dapp/valnode/types"
)

func (val *ValNode) Exec_Node(node *pty.ValNode, tx *types.Transaction, index int) (*types.Receipt, error) {
	receipt := &types.Receipt{types.ExecOk, nil, nil}
	return receipt, nil
}

func (val *ValNode) Exec_BlockInfo(blockInfo *pty.TendermintBlockInfo, tx *types.Transaction, index int) (*types.Receipt, error) {
	receipt := &types.Receipt{types.ExecOk, nil, nil}
	return receipt, nil
}