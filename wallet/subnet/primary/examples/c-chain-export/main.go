// Copyright (C) 2019-2023, Lux Partners Limited. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"context"
	"log"
	"time"

	"github.com/luxfi/node/genesis"
	"github.com/luxfi/node/ids"
	"github.com/luxfi/node/utils/constants"
	"github.com/luxfi/node/utils/units"
	"github.com/luxfi/node/vms/secp256k1fx"
	"github.com/luxfi/node/wallet/subnet/primary"
)

func main() {
	key := genesis.EWOQKey
	uri := primary.LocalAPIURI
	kc := secp256k1fx.NewKeychain(key)
	luxAddr := key.Address()

	ctx := context.Background()

	// MakeWallet fetches the available UTXOs owned by [kc] on the network that
	// [uri] is hosting.
	walletSyncStartTime := time.Now()
	wallet, err := primary.MakeWallet(ctx, &primary.WalletConfig{
		URI:          uri,
		LUXKeychain: kc,
		EthKeychain:  kc,
	})
	if err != nil {
		log.Fatalf("failed to initialize wallet: %s\n", err)
	}
	log.Printf("synced wallet in %s\n", time.Since(walletSyncStartTime))

	// Get the P-chain wallet
	pWallet := wallet.P()
	cWallet := wallet.C()

	// Pull out useful constants to use when issuing transactions.
	cChainID := cWallet.BlockchainID()
	owner := secp256k1fx.OutputOwners{
		Threshold: 1,
		Addrs: []ids.ShortID{
			luxAddr,
		},
	}

	exportStartTime := time.Now()
	exportTx, err := cWallet.IssueExportTx(
		constants.PlatformChainID,
		[]*secp256k1fx.TransferOutput{{
			Amt:          units.Lux,
			OutputOwners: owner,
		}},
	)
	if err != nil {
		log.Fatalf("failed to issue export transaction: %s\n", err)
	}
	log.Printf("issued export %s in %s\n", exportTx.ID(), time.Since(exportStartTime))

	importStartTime := time.Now()
	importTx, err := pWallet.IssueImportTx(cChainID, &owner)
	if err != nil {
		log.Fatalf("failed to issue import transaction: %s\n", err)
	}
	log.Printf("issued import %s in %s\n", importTx.ID(), time.Since(importStartTime))
}
