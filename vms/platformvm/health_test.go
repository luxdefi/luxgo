// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package platformvm

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/version"
)

const defaultMinConnectedStake = 0.8

func TestHealthCheckPrimaryNetwork(t *testing.T) {
	require := require.New(t)

	vm, _, _ := defaultVM()
	vm.ctx.Lock.Lock()

	defer func() {
		require.NoError(vm.Shutdown(context.Background()))
		vm.ctx.Lock.Unlock()
	}()
	genesisState, _ := defaultGenesis()
	for index, validator := range genesisState.Validators {
		err := vm.Connected(context.Background(), validator.NodeID, version.CurrentApp)
		require.NoError(err)
		details, err := vm.HealthCheck(context.Background())
		if float64((index+1)*20) >= defaultMinConnectedStake*100 {
			require.NoError(err)
		} else {
			require.Contains(details, "primary-percentConnected")
			require.ErrorIs(err, errNotEnoughStake)
		}
	}
}

func TestHealthCheckSubnet(t *testing.T) {
	tests := map[string]struct {
		minStake   float64
		useDefault bool
	}{
		"default min stake": {
			useDefault: true,
			minStake:   0,
		},
		"custom min stake": {
			minStake: 0.40,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			require := require.New(t)

			vm, _, _ := defaultVM()
			vm.ctx.Lock.Lock()
			defer func() {
				require.NoError(vm.Shutdown(context.Background()))
				vm.ctx.Lock.Unlock()
			}()

			subnetID := ids.GenerateTestID()
			subnetVdrs := validators.NewSet()
<<<<<<< HEAD
			vm.TrackedSubnets.Add(subnetID)
			testVdrCount := 4
			for i := 0; i < testVdrCount; i++ {
				subnetVal := ids.GenerateTestNodeID()
				err := subnetVdrs.Add(subnetVal, nil, ids.Empty, 100)
				require.NoError(err)
			}
<<<<<<< HEAD
			ok := vm.Validators.Add(subnetID, subnetVdrs)
=======

			vals, ok := vm.Validators.Get(subnetID)
>>>>>>> f6ea8e56f (Rename validators.Manager#GetValidators to Get (#2279))
=======
			vm.WhitelistedSubnets.Add(subnetID)
			testVdrCount := 4
			for i := 0; i < testVdrCount; i++ {
				subnetVal := ids.GenerateTestNodeID()
				err := subnetVdrs.Add(subnetVal, nil, ids.Empty, 100)
				require.NoError(err)
			}
			ok := vm.Validators.Add(subnetID, subnetVdrs)
>>>>>>> f171d317d (Remove unnecessary functions from validators.Manager interface (#2277))
			require.True(ok)

			// connect to all primary network validators first
			genesisState, _ := defaultGenesis()
			for _, validator := range genesisState.Validators {
				err := vm.Connected(context.Background(), validator.NodeID, version.CurrentApp)
				require.NoError(err)
			}
			var expectedMinStake float64
			if test.useDefault {
				expectedMinStake = defaultMinConnectedStake
			} else {
				expectedMinStake = test.minStake
				vm.MinPercentConnectedStakeHealthy = map[ids.ID]float64{
					subnetID: expectedMinStake,
				}
			}
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
			for index, vdr := range subnetVdrs.List() {
				err := vm.ConnectedSubnet(context.Background(), vdr.NodeID, subnetID)
=======
			for index, validator := range vals.List() {
=======
			for index, validator := range subnetVdrs.List() {
>>>>>>> f171d317d (Remove unnecessary functions from validators.Manager interface (#2277))
				err := vm.Connected(context.Background(), validator.ID(), version.CurrentApp)
>>>>>>> 5be92660b (Pass message context through the VM interface (#2219))
=======
			for index, vdr := range subnetVdrs.List() {
<<<<<<< HEAD
				err := vm.Connected(context.Background(), vdr.NodeID, version.CurrentApp)
>>>>>>> 3e2b5865d (Convert validators.Validator into a struct (#2185))
=======
				err := vm.ConnectedSubnet(context.Background(), vdr.NodeID, subnetID)
>>>>>>> d6c7e2094 (Track subnet uptimes (#1427))
				require.NoError(err)
				details, err := vm.HealthCheck(context.Background())
				connectedPerc := float64((index + 1) * (100 / testVdrCount))
				if connectedPerc >= expectedMinStake*100 {
					require.NoError(err)
				} else {
					require.Contains(details, fmt.Sprintf("%s-percentConnected", subnetID))
					require.ErrorIs(err, errNotEnoughStake)
				}
			}
		})
	}
}
