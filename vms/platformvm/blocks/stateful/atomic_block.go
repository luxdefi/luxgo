// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package stateful

import (
	"errors"

	"github.com/ava-labs/avalanchego/chains/atomic"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/choices"
	"github.com/ava-labs/avalanchego/vms/platformvm/blocks/stateless"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs/executor"
)

var (
	ErrConflictingParentTxs = errors.New("block contains a transaction that conflicts with a transaction in a parent block")

	_ Block    = &AtomicBlock{}
	_ Decision = &AtomicBlock{}
)

// AtomicBlock being accepted results in the atomic transaction contained in the
// block to be accepted and committed to the chain.
type AtomicBlock struct {
	// TODO set this field
	acceptor Acceptor
	// TODO set this field
	rejector Rejector
	*stateless.AtomicBlock
	*decisionBlock

	// inputs are the atomic inputs that are consumed by this block's atomic
	// transaction
	inputs ids.Set

	atomicRequests map[ids.ID]*atomic.Requests
}

// NewAtomicBlock returns a new *AtomicBlock where the block's parent, a
// decision block, has ID [parentID].
func NewAtomicBlock(
	verifier Verifier2,
	txExecutorBackend executor.Backend,
	parentID ids.ID,
	height uint64,
	tx *txs.Tx,
) (*AtomicBlock, error) {
	statelessBlk, err := stateless.NewAtomicBlock(parentID, height, tx)
	if err != nil {
		return nil, err
	}
	return toStatefulAtomicBlock(statelessBlk, verifier, txExecutorBackend, choices.Processing)
}

func toStatefulAtomicBlock(
	statelessBlk *stateless.AtomicBlock,
	verifier Verifier2,
	txExecutorBackend executor.Backend,
	status choices.Status,
) (*AtomicBlock, error) {
	ab := &AtomicBlock{
		AtomicBlock: statelessBlk,
		decisionBlock: &decisionBlock{
			commonBlock: &commonBlock{
				baseBlk:           &statelessBlk.CommonBlock,
				status:            status,
				verifier:          verifier,
				txExecutorBackend: txExecutorBackend,
			},
		},
	}

	ab.Tx.Unsigned.InitCtx(ab.txExecutorBackend.Ctx)
	return ab, nil
}

// conflicts checks to see if the provided input set contains any conflicts with
// any of this block's non-accepted ancestors or itself.
func (ab *AtomicBlock) conflicts(s ids.Set) (bool, error) {
	return ab.conflictsAtomicBlock(ab, s)
	/* TODO remove
	if ab.Status() == choices.Accepted {
		return false, nil
	}
	if ab.inputs.Overlaps(s) {
		return true, nil
	}
	parent, err := ab.parentBlock()
	if err != nil {
		return false, err
	}
	return parent.conflicts(s)
	*/
}

// Verify this block performs a valid state transition.
//
// The parent block must be a decision block
//
// This function also sets onAcceptDB database if the verification passes.
func (ab *AtomicBlock) Verify() error {
	return ab.verifier.verifyAtomicBlock(ab)
}

func (ab *AtomicBlock) Accept() error {
	return ab.acceptor.acceptAtomicBlock(ab)
}

func (ab *AtomicBlock) Reject() error {
	return ab.rejector.rejectAtomicBlock(ab)
}
