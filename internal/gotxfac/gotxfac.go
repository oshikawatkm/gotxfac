package gotxfac

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Settings struct {
	PrevTxId string
	TxOut    int
}

func StrStdin() (stringInput string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	stringInput = scanner.Text()

	stringInput = strings.TrimSpace(stringInput)
	return stringInput
}

func ScanSettings() (string, int) {
	stdin := bufio.NewScanner(os.Stdin)
	fmt.Printf("[+] Set prevTxId: ")
	stdin.Scan()
	prevTxId := StrStdin()

	fmt.Printf("[+] Set txout: ")
	stdin.Scan()
	txoutstr := StrStdin()
	txout, _ := strconv.Atoi(strings.TrimSpace(txoutstr))

	return prevTxId, txout
}

func New(prevTxId string, txout int) *Settings {
	return &Settings{
		PrevTxId: prevTxId,
		TxOut:    txout,
	}
}
