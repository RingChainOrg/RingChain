package util

import (
	"math/big"
	"strconv"
	"errors"
	"crypto/elliptic"
	"io"
	"crypto/ecdsa"
)

func ConverEthToWei(amountString string) (*big.Int, error){
	//转账的数量, 单位从ETHER ->Wei
	amount,error:= strconv.ParseFloat(amountString, 64);
	if error != nil {
		return  nil, errors.New("转账数字错误: " + error.Error());
	}

	//浮点数=>整型大数; 大数溢出预防
	intNum := uint64(amount); //整数位
	decimalNum := amount - float64(intNum);//小数位

	bigInt :=new(big.Int).SetUint64(intNum);
	bigInt = bigInt.Mul(bigInt, new(big.Int).SetUint64(1000000000000000000))
	bigdecimalInt  :=new(big.Int).SetUint64(uint64(decimalNum * 1000000000000000000));
	bigNumAmount := bigInt.Add(bigInt, bigdecimalInt)

	return bigNumAmount, nil;
}

func ConverEthFloatToWei(amount float64) (*big.Int, error){
	//转账的数量, 单位从ETHER ->Wei
	//浮点数=>整型大数; 大数溢出预防
	intNum := uint64(amount); //整数位
	decimalNum := amount - float64(intNum);//小数位

	bigInt :=new(big.Int).SetUint64(intNum);
	bigInt = bigInt.Mul(bigInt, new(big.Int).SetUint64(1000000000000000000))
	bigdecimalInt  :=new(big.Int).SetUint64(uint64(decimalNum * 1000000000000000000));
	bigNumAmount := bigInt.Add(bigInt, bigdecimalInt)

	return bigNumAmount, nil;
}

// 产生一对公钥/私钥, GenerateKey generates a public and private key pair.
func GenerateKey(key *big.Int, c elliptic.Curve, rand io.Reader) (*ecdsa.PrivateKey, error) {
	priv := new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = c
	priv.D = key
	priv.PublicKey.X, priv.PublicKey.Y = c.ScalarBaseMult(key.Bytes())
	return priv, nil
}