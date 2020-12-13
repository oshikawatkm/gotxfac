package main

import (
	"fmt"

	"github.com/oshikawatkm/gotxfac/internal/gotxfac"
	"github.com/oshikawatkm/gotxfac/pkg/transaction"
)

var scan_flag bool = false

func main() {
	fmt.Println("[*] Start GOTXFAC!!!")

	for {
		prevTxId, txout, _ := gotxfac.ScanSettings()
		scan_flag, _ := gotxfac.CheckSettings(prevTxId, txout)
		if scan_flag == true {
			fmt.Printf("[*] try again.")
			break
		}
	}

	tx, _ := transaction.CreateTx("hoge", "hoge", 6000)
	fmt.Printf(tx)
}
