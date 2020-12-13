package transaction

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
)

func NewTx() (*wire.MsgTx, error) {
	return wire.NewMsgTx(wire.TxVersion), nil
}

func GetUTXO(address string) (string, int64, string, error) {
	var previousTxid string = "16688d2946c3e029ca91ce730109994c2bcafb859d580a6f7c820fb2bb5b6afc"
	var balance int64 = 62000
	var pubKeyScript string = "76a91455d5e92958a8b06b4ff15cd2dd3d254f375e98db88ac"
	return previousTxid, balance, pubKeyScript, nil
}

func CreateTx(privKey string, destination string, amount int64) (string, error) {

	wif, err := btcutil.DecodeWIF(privKey)
	if err != nil {
		log.Fatal("[-] DecodeWIF error: ", err)
		return "", nil
	}

	addrPubKey, err := btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeUncompressed(), &chaincfg.TestNet3Params)
	if err != nil {
		log.Fatal("[-] NewAddressPubKey error: ", err)
		return "", err
	}

	txid, balance, pkScript, err := GetUTXO(addrPubKey.EncodeAddress())
	if err != nil {
		log.Fatal("[-] GetUTXO error: ", err)
		return "", err
	}

	if balance < amount {
		return "", fmt.Errorf("the balance of the account is not sufficient")
	}

	destinationAddr, err := btcutil.DecodeAddress(destination, &chaincfg.TestNet3Params)
	if err != nil {
		log.Fatal("[-] DecodeAddress error: ", err)
		return "", err
	}

	destinationAddrByte, err := txscript.PayToAddrScript(destinationAddr)
	if err != nil {
		log.Fatal("[-] PayToAddrScript error: ", err)
		return "", err
	}

	redeemTx, err := NewTx()
	if err != nil {
		log.Fatal("[-] NewTx error: ", err)
		return "", err
	}

	utxoHash, err := chainhash.NewHashFromStr(txid)
	if err != nil {
		log.Fatal("[-] NewHashFromStr error: ", err)
		return "", err
	}

	outPoint := wire.NewOutPoint(utxoHash, 0)

	txIn := wire.NewTxIn(outPoint, nil, nil)
	redeemTx.AddTxIn(txIn)

	redeemTxOut := wire.NewTxOut(amount, destinationAddrByte)
	redeemTx.AddTxOut(redeemTxOut)

	finalRawTx, err := SignTx(privKey, pkScript, redeemTx)
	return finalRawTx, nil
}

func SignTx(privKey string, pkScript string, redeemTx *wire.MsgTx) (string, error) {

	wif, err := btcutil.DecodeWIF(privKey)
	if err != nil {
		return "", err
	}

	sourcePKScript, err := hex.DecodeString(pkScript)
	if err != nil {
		return "", nil
	}

	// since there is only one input in our transaction
	// we use 0 as second argument, if the transaction
	// has more args, should pass related index
	signature, err := txscript.SignatureScript(redeemTx, 0, sourcePKScript, txscript.SigHashAll, wif.PrivKey, false)
	if err != nil {
		return "", nil
	}

	// since there is only one input, and want to add
	// signature to it use 0 as index
	redeemTx.TxIn[0].SignatureScript = signature

	var signedTx bytes.Buffer
	redeemTx.Serialize(&signedTx)

	hexSignedTx := hex.EncodeToString(signedTx.Bytes())

	return hexSignedTx, nil
}
