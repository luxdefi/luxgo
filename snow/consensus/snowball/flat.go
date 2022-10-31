// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowball

import (
	"github.com/luxdefi/luxd/ids"
)

var (
	_ Factory   = (*FlatFactory)(nil)
	_ Consensus = (*Flat)(nil)
)

// FlatFactory implements Factory by returning a flat struct
type FlatFactory struct{}

func (FlatFactory) New() Consensus { return &Flat{} }

// Flat is a naive implementation of a multi-choice snowball instance
type Flat struct {
	// wraps the n-nary snowball logic
	nnarySnowball

	// params contains all the configurations of a snowball instance
	params Parameters
}

func (f *Flat) Initialize(params Parameters, choice ids.ID) {
	f.nnarySnowball.Initialize(params.BetaVirtuous, params.BetaRogue, choice)
	f.params = params
}

func (f *Flat) Parameters() Parameters { return f.params }

func (f *Flat) RecordPoll(votes ids.Bag) bool {
	if pollMode, numVotes := votes.Mode(); numVotes >= f.params.Alpha {
		f.RecordSuccessfulPoll(pollMode)
		return true
	}

	f.RecordUnsuccessfulPoll()
	return false
}
