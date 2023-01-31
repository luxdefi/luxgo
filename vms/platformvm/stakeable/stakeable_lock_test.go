// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package stakeable

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/require"

	"github.com/ava-labs/avalanchego/vms/components/avax"
)

var errTest = errors.New("hi mom")

func TestLockOutVerify(t *testing.T) {
	tests := []struct {
		name             string
		locktime         uint64
		transferableOutF func(*gomock.Controller) avax.TransferableOut
		expectedErr      error
	}{
		{
			name:     "happy path",
			locktime: 1,
			transferableOutF: func(ctrl *gomock.Controller) avax.TransferableOut {
				o := avax.NewMockTransferableOut(ctrl)
				o.EXPECT().Verify().Return(nil)
				return o
			},
			expectedErr: nil,
		},
		{
			name:     "invalid locktime",
			locktime: 0,
			transferableOutF: func(ctrl *gomock.Controller) avax.TransferableOut {
				return nil
			},
			expectedErr: errInvalidLocktime,
		},
		{
			name:     "nested",
			locktime: 1,
			transferableOutF: func(ctrl *gomock.Controller) avax.TransferableOut {
				return &LockOut{}
			},
			expectedErr: errNestedStakeableLocks,
		},
		{
			name:     "inner output fails verification",
			locktime: 1,
			transferableOutF: func(ctrl *gomock.Controller) avax.TransferableOut {
				o := avax.NewMockTransferableOut(ctrl)
				o.EXPECT().Verify().Return(errTest)
				return o
			},
			expectedErr: errTest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
<<<<<<< HEAD
<<<<<<< HEAD
=======
			require := require.New(t)
>>>>>>> a8631aa5c (Add Fx tests (#1838))
=======
>>>>>>> 7c09e7074 (Standardize `require` usage and remove `t.Fatal` from platformvm (#2297))
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			lockOut := &LockOut{
				Locktime:        tt.locktime,
				TransferableOut: tt.transferableOutF(ctrl),
			}
<<<<<<< HEAD
<<<<<<< HEAD
			require.Equal(t, tt.expectedErr, lockOut.Verify())
=======
			require.Equal(tt.expectedErr, lockOut.Verify())
>>>>>>> a8631aa5c (Add Fx tests (#1838))
=======
			require.Equal(t, tt.expectedErr, lockOut.Verify())
>>>>>>> 7c09e7074 (Standardize `require` usage and remove `t.Fatal` from platformvm (#2297))
		})
	}
}

func TestLockInVerify(t *testing.T) {
	tests := []struct {
		name            string
		locktime        uint64
		transferableInF func(*gomock.Controller) avax.TransferableIn
		expectedErr     error
	}{
		{
			name:     "happy path",
			locktime: 1,
			transferableInF: func(ctrl *gomock.Controller) avax.TransferableIn {
				o := avax.NewMockTransferableIn(ctrl)
				o.EXPECT().Verify().Return(nil)
				return o
			},
			expectedErr: nil,
		},
		{
			name:     "invalid locktime",
			locktime: 0,
			transferableInF: func(ctrl *gomock.Controller) avax.TransferableIn {
				return nil
			},
			expectedErr: errInvalidLocktime,
		},
		{
			name:     "nested",
			locktime: 1,
			transferableInF: func(ctrl *gomock.Controller) avax.TransferableIn {
				return &LockIn{}
			},
			expectedErr: errNestedStakeableLocks,
		},
		{
			name:     "inner input fails verification",
			locktime: 1,
			transferableInF: func(ctrl *gomock.Controller) avax.TransferableIn {
				o := avax.NewMockTransferableIn(ctrl)
				o.EXPECT().Verify().Return(errTest)
				return o
			},
			expectedErr: errTest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
<<<<<<< HEAD
<<<<<<< HEAD
=======
			require := require.New(t)
>>>>>>> a8631aa5c (Add Fx tests (#1838))
=======
>>>>>>> 7c09e7074 (Standardize `require` usage and remove `t.Fatal` from platformvm (#2297))
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			lockOut := &LockIn{
				Locktime:       tt.locktime,
				TransferableIn: tt.transferableInF(ctrl),
			}
<<<<<<< HEAD
<<<<<<< HEAD
			require.Equal(t, tt.expectedErr, lockOut.Verify())
=======
			require.Equal(tt.expectedErr, lockOut.Verify())
>>>>>>> a8631aa5c (Add Fx tests (#1838))
=======
			require.Equal(t, tt.expectedErr, lockOut.Verify())
>>>>>>> 7c09e7074 (Standardize `require` usage and remove `t.Fatal` from platformvm (#2297))
		})
	}
}
