package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/decred/dcrd/chaincfg"
	"github.com/decred/dcrd/dcrec"
	"github.com/decred/dcrd/dcrec/secp256k1"
	"github.com/decred/dcrd/dcrutil"
)

func generateKeys(params *chaincfg.Params) error {
	key, err := secp256k1.GeneratePrivateKey()
	if err != nil {
		return err
	}

	// Key
	wif, err := dcrutil.NewWIF(key, params, dcrec.STEcdsaSecp256k1)
	if err != nil {
		return err
	}

	fmt.Printf("Private key: %x\n", key.Serialize())
	fmt.Printf("Public  key: %x\n", key.PubKey().Serialize())
	fmt.Printf("WIF        : %s\n", wif)

	return nil
}

func main() {
	mainnet := flag.Bool("mainnet", false, "use testnet parameters")
	simnet := flag.Bool("simnet", false, "use testnet parameters")
	testnet := flag.Bool("testnet", false, "use testnet parameters")
	flag.Parse()

	var net *chaincfg.Params
	flags := 0
	if *mainnet {
		flags++
		net = &chaincfg.MainNetParams
	}
	if *testnet {
		flags++
		net = &chaincfg.TestNet3Params
	}
	if *simnet {
		flags++
		net = &chaincfg.SimNetParams
	}
	if flags != 1 {
		fmt.Println("One and only one flag must be selected")
		flag.Usage()
		os.Exit(1)
	}

	if err := generateKeys(net); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
