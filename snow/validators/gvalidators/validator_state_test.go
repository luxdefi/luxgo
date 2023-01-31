// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gvalidators

import (
	"context"
	"errors"
	"net"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/require"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/vms/rpcchainvm/grpcutils"

	pb "github.com/ava-labs/avalanchego/proto/pb/validatorstate"
)

const bufSize = 1024 * 1024

var errCustom = errors.New("custom")

type testState struct {
	client  *Client
	server  *validators.MockState
	closeFn func()
}

func setupState(t testing.TB, ctrl *gomock.Controller) *testState {
	t.Helper()

	state := &testState{
		server: validators.NewMockState(ctrl),
	}

	listener := bufconn.Listen(bufSize)
	serverCloser := grpcutils.ServerCloser{}

	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		server := grpcutils.NewDefaultServer(opts)
		pb.RegisterValidatorStateServer(server, NewServer(state.server))
		serverCloser.Add(server)
		return server
	}

	go grpcutils.Serve(listener, serverFunc)

	dialer := grpc.WithContextDialer(
		func(context.Context, string) (net.Conn, error) {
			return listener.Dial()
		},
	)

	dopts := grpcutils.DefaultDialOptions
	dopts = append(dopts, dialer)
	conn, err := grpcutils.Dial("", dopts...)
	if err != nil {
		t.Fatalf("Failed to dial: %s", err)
	}

	state.client = NewClient(pb.NewValidatorStateClient(conn))
	state.closeFn = func() {
		serverCloser.Stop()
		_ = conn.Close()
		_ = listener.Close()
	}
	return state
}

func TestGetMinimumHeight(t *testing.T) {
	require := require.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	state := setupState(t, ctrl)
	defer state.closeFn()

	// Happy path
	expectedHeight := uint64(1337)
	state.server.EXPECT().GetMinimumHeight(gomock.Any()).Return(expectedHeight, nil)

	height, err := state.client.GetMinimumHeight(context.Background())
	require.NoError(err)
	require.Equal(expectedHeight, height)

	// Error path
	state.server.EXPECT().GetMinimumHeight(gomock.Any()).Return(expectedHeight, errCustom)

	_, err = state.client.GetMinimumHeight(context.Background())
	require.Error(err)
}

func TestGetCurrentHeight(t *testing.T) {
	require := require.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	state := setupState(t, ctrl)
	defer state.closeFn()

	// Happy path
	expectedHeight := uint64(1337)
	state.server.EXPECT().GetCurrentHeight(gomock.Any()).Return(expectedHeight, nil)

	height, err := state.client.GetCurrentHeight(context.Background())
	require.NoError(err)
	require.Equal(expectedHeight, height)

	// Error path
	state.server.EXPECT().GetCurrentHeight(gomock.Any()).Return(expectedHeight, errCustom)

	_, err = state.client.GetCurrentHeight(context.Background())
<<<<<<< HEAD
	require.Error(err)
}

func TestGetSubnetID(t *testing.T) {
	require := require.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	state := setupState(t, ctrl)
	defer state.closeFn()

	// Happy path
	chainID := ids.GenerateTestID()
	expectedSubnetID := ids.GenerateTestID()
	state.server.EXPECT().GetSubnetID(gomock.Any(), chainID).Return(expectedSubnetID, nil)

	subnetID, err := state.client.GetSubnetID(context.Background(), chainID)
	require.NoError(err)
	require.Equal(expectedSubnetID, subnetID)

	// Error path
	state.server.EXPECT().GetSubnetID(gomock.Any(), chainID).Return(expectedSubnetID, errCustom)

	_, err = state.client.GetSubnetID(context.Background(), chainID)
=======
>>>>>>> f94b52cf8 ( Pass message context through the validators.State interface (#2242))
	require.Error(err)
}

func TestGetValidatorSet(t *testing.T) {
	require := require.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	state := setupState(t, ctrl)
	defer state.closeFn()

	// Happy path
	sk0, err := bls.NewSecretKey()
	require.NoError(err)
<<<<<<< HEAD
<<<<<<< HEAD
	vdr0 := &validators.GetValidatorOutput{
=======
	vdr0 := &validators.Validator{
>>>>>>> 117ff9a78 (Add BLS keys to `GetValidatorSet` (#2111))
=======
	vdr0 := &validators.GetValidatorOutput{
>>>>>>> 62b728221 (Add txID to `validators.Set#Add` (#2312))
		NodeID:    ids.GenerateTestNodeID(),
		PublicKey: bls.PublicFromSecretKey(sk0),
		Weight:    1,
	}

	sk1, err := bls.NewSecretKey()
	require.NoError(err)
<<<<<<< HEAD
<<<<<<< HEAD
	vdr1 := &validators.GetValidatorOutput{
=======
	vdr1 := &validators.Validator{
>>>>>>> 117ff9a78 (Add BLS keys to `GetValidatorSet` (#2111))
=======
	vdr1 := &validators.GetValidatorOutput{
>>>>>>> 62b728221 (Add txID to `validators.Set#Add` (#2312))
		NodeID:    ids.GenerateTestNodeID(),
		PublicKey: bls.PublicFromSecretKey(sk1),
		Weight:    2,
	}

<<<<<<< HEAD
<<<<<<< HEAD
	vdr2 := &validators.GetValidatorOutput{
=======
	vdr2 := &validators.Validator{
>>>>>>> 117ff9a78 (Add BLS keys to `GetValidatorSet` (#2111))
=======
	vdr2 := &validators.GetValidatorOutput{
>>>>>>> 62b728221 (Add txID to `validators.Set#Add` (#2312))
		NodeID:    ids.GenerateTestNodeID(),
		PublicKey: nil,
		Weight:    3,
	}

<<<<<<< HEAD
<<<<<<< HEAD
	expectedVdrs := map[ids.NodeID]*validators.GetValidatorOutput{
=======
	expectedVdrs := map[ids.NodeID]*validators.Validator{
>>>>>>> 117ff9a78 (Add BLS keys to `GetValidatorSet` (#2111))
=======
	expectedVdrs := map[ids.NodeID]*validators.GetValidatorOutput{
>>>>>>> 62b728221 (Add txID to `validators.Set#Add` (#2312))
		vdr0.NodeID: vdr0,
		vdr1.NodeID: vdr1,
		vdr2.NodeID: vdr2,
	}
	height := uint64(1337)
	subnetID := ids.GenerateTestID()
	state.server.EXPECT().GetValidatorSet(gomock.Any(), height, subnetID).Return(expectedVdrs, nil)

	vdrs, err := state.client.GetValidatorSet(context.Background(), height, subnetID)
	require.NoError(err)
	require.Equal(expectedVdrs, vdrs)

	// Error path
	state.server.EXPECT().GetValidatorSet(gomock.Any(), height, subnetID).Return(expectedVdrs, errCustom)

	_, err = state.client.GetValidatorSet(context.Background(), height, subnetID)
	require.Error(err)
}
