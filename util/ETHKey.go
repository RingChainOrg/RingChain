package util

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	//"ethlog"
	"errors"
	"github.com/regcostajr/go-web3/utils"
	"github.com/regcostajr/go-web3/complex/types"
	"strings"
	"fmt"
)

//func GetKey(fromstr string, fullpath string) (accounts.Wallet,*accounts.Account ){
//	from := common.HexToAddress(fromstr)
//
//	//var defaultKeyStores cli.StringSlice
//	defaultKeyStores := []string{
//		fullpath,
//		//os.Getenv("APPDATA") + "/Ethereum/keystore",
//		//os.Getenv("APPDATA") + "/Parity/Ethereum/keys",
//	}
//
//	backends := []accounts.Backend{ }
//
//	var paths []string
//	paths = defaultKeyStores
//
//	for _, x := range(paths) {
//		ks := keystore.NewKeyStore(
//			x, keystore.StandardScryptN, keystore.StandardScryptP)
//		backends = append(backends, ks)
//	}
//
//	manager := accounts.NewManager(backends...)
//	wallets := manager.Wallets()
//	var wallet accounts.Wallet
//	var acct *accounts.Account
//
//	needPassphrase := true
//
//	//找到keystore
//Scan:
//	for _, x := range(wallets) {
//		if x.URL().Scheme == "keystore" {
//			for _, y := range(x.Accounts()) {
//				if (y.Address == from) {
//					wallet = x
//					acct = &y
//					break Scan
//				}
//			}
//		}
//	}
//
//	if acct == nil {
//		ethlog.Fatal("error")
//		return nil,nil
//	}
//
//	if needPassphrase {
//
//	}
//
//	return wallet,acct
//}

func GetPrivateKey(fromstr string,password string , fullpath string) (string, error ){
	from := common.HexToAddress(fromstr)

	//var defaultKeyStores cli.StringSlice
	defaultKeyStores := []string{
		fullpath,
		//os.Getenv("APPDATA") + "/Ethereum/keystore",
		//os.Getenv("APPDATA") + "/Parity/Ethereum/keys",
	}

	//backends := []accounts.Backend{ }

	var paths []string
	paths = defaultKeyStores

	for _, x := range(paths) {
		ks := keystore.NewKeyStore(
			x, keystore.StandardScryptN, keystore.StandardScryptP)

		for _, y := range(ks.Accounts()) {
			if (y.Address == from) {
				account := y;
				privatekey, error := ks.ExportPrivateKey(account, password)
				if error != nil {
					return "",error
				}
				return privatekey, nil;
			}
		}

	}

	return "", errors.New("not found the account :" + fromstr)
}


// prepareTransaction ...
func GetTransferData(utils *utils.Utils, address string, amount uint32) (string, error) {

	fullFunction := "transfer(address,uint256)"
	sha3Function, err := utils.Sha3(types.ComplexString(fullFunction))

	if err != nil {
		return "", err
	}

	var data string
	data += getHexValue("address",address)
	data += getHexValue("uint256",amount)

	return string(types.ComplexString(sha3Function[0:10] + data)), nil
}

func getHexValue(inputType string, value interface{}) string {

	var data string

	if strings.HasPrefix(inputType, "uint256") {
		data += fmt.Sprintf("%064s", fmt.Sprintf("%x", value.(uint32)))
		return data
	}

	if strings.HasPrefix(inputType, "int") ||
		strings.HasPrefix(inputType, "uint") ||
		strings.HasPrefix(inputType, "fixed") ||
		strings.HasPrefix(inputType, "ufixed") {
		data += fmt.Sprintf("%064s", fmt.Sprintf("%x", value.(int)))
	}

	if strings.Compare("address", inputType) == 0 {
		data += fmt.Sprintf("%064s", value.(string)[2:])
	}

	if strings.Compare("string", inputType) == 0 {
		data += fmt.Sprintf("%064s", fmt.Sprintf("%x", value.(string)))
	}

	return data

}