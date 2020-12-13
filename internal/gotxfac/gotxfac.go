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

func ScanSettings() (string, int, error) {
	fmt.Printf("[+] Set prevTxId: ")
	prevTxId := StrStdin()

	fmt.Printf("[+] Set txout: ")
	txoutstr := StrStdin()
	txout, _ := strconv.Atoi(strings.TrimSpace(txoutstr))

	return prevTxId, txout, nil
}

func CheckSettings(prevTxId string, txout int) (bool, error) {
	fmt.Printf("[+] Are you sure about this setting? [y/n]: ")

	checked := StrStdin()
	if checked == "y" {
		return true, nil
	} else if checked == "f" {
		return false, nil
	} else {
		return false, nil
	}
}

func New(prevTxId string, txout int) *Settings {
	return &Settings{
		PrevTxId: prevTxId,
		TxOut:    txout,
	}
}
