package main

import (
	"WebServerExample/ETHRpc"
	"ethlog"
	"github.com/regcostajr/go-web3/dto"
	"math/big"
)

func main()  {
	//初始化rpc
	rpc :=  ETHRpc.RPCClient{};
	rpc.InitClient();

	//设置合约的ABI
	abi := "[ { \"constant\": true, \"inputs\": [], \"name\": \"name\", \"outputs\": [ { \"name\": \"\", \"type\": \"string\", \"value\": \"WECTest\" } ], \"payable\": false, \"stateMutability\": \"view\", \"type\": \"function\" }, { \"constant\": false, \"inputs\": [ { \"name\": \"_spender\", \"type\": \"address\" }, { \"name\": \"_value\", \"type\": \"uint256\" } ], \"name\": \"approve\", \"outputs\": [ { \"name\": \"\", \"type\": \"bool\" } ], \"payable\": false, \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"constant\": false, \"inputs\": [ { \"name\": \"delAdmin\", \"type\": \"address\" } ], \"name\": \"removeAdmin\", \"outputs\": [], \"payable\": false, \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"constant\": true, \"inputs\": [], \"name\": \"totalSupply\", \"outputs\": [ { \"name\": \"\", \"type\": \"uint256\", \"value\": \"1000\" } ], \"payable\": false, \"stateMutability\": \"view\", \"type\": \"function\" }, { \"constant\": false, \"inputs\": [ { \"name\": \"_from\", \"type\": \"address\" }, { \"name\": \"_to\", \"type\": \"address\" }, { \"name\": \"_value\", \"type\": \"uint256\" } ], \"name\": \"transferFrom\", \"outputs\": [ { \"name\": \"\", \"type\": \"bool\" } ], \"payable\": false, \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"constant\": true, \"inputs\": [], \"name\": \"decimals\", \"outputs\": [ { \"name\": \"\", \"type\": \"uint8\", \"value\": \"1\" } ], \"payable\": false, \"stateMutability\": \"view\", \"type\": \"function\" }, { \"constant\": true, \"inputs\": [ { \"name\": \"\", \"type\": \"address\" } ], \"name\": \"admins\", \"outputs\": [ { \"name\": \"\", \"type\": \"bool\", \"value\": false } ], \"payable\": false, \"stateMutability\": \"view\", \"type\": \"function\" }, { \"constant\": false, \"inputs\": [ { \"name\": \"_spender\", \"type\": \"address\" }, { \"name\": \"_subtractedValue\", \"type\": \"uint256\" } ], \"name\": \"decreaseApproval\", \"outputs\": [ { \"name\": \"\", \"type\": \"bool\" } ], \"payable\": false, \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"constant\": false, \"inputs\": [ { \"name\": \"newAdmin\", \"type\": \"address\" } ], \"name\": \"addAdmin\", \"outputs\": [], \"payable\": false, \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"constant\": true, \"inputs\": [ { \"name\": \"_owner\", \"type\": \"address\" } ], \"name\": \"balanceOf\", \"outputs\": [ { \"name\": \"balance\", \"type\": \"uint256\", \"value\": \"0\" } ], \"payable\": false, \"stateMutability\": \"view\", \"type\": \"function\" }, { \"constant\": false, \"inputs\": [ { \"name\": \"target\", \"type\": \"address\" }, { \"name\": \"mintedAmount\", \"type\": \"uint256\" } ], \"name\": \"mintToken\", \"outputs\": [], \"payable\": false, \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"constant\": true, \"inputs\": [], \"name\": \"owner\", \"outputs\": [ { \"name\": \"\", \"type\": \"address\", \"value\": \"0x6b1a125ab9895ecb7d7d5eb550ffc1dfdf618a27\" } ], \"payable\": false, \"stateMutability\": \"view\", \"type\": \"function\" }, { \"constant\": true, \"inputs\": [], \"name\": \"symbol\", \"outputs\": [ { \"name\": \"\", \"type\": \"string\", \"value\": \"WECT\" } ], \"payable\": false, \"stateMutability\": \"view\", \"type\": \"function\" }, { \"constant\": false, \"inputs\": [ { \"name\": \"_to\", \"type\": \"address\" }, { \"name\": \"_value\", \"type\": \"uint256\" } ], \"name\": \"transfer\", \"outputs\": [ { \"name\": \"\", \"type\": \"bool\" } ], \"payable\": false, \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"constant\": true, \"inputs\": [ { \"name\": \"\", \"type\": \"address\" } ], \"name\": \"frozenAccount\", \"outputs\": [ { \"name\": \"\", \"type\": \"bool\", \"value\": false } ], \"payable\": false, \"stateMutability\": \"view\", \"type\": \"function\" }, { \"constant\": false, \"inputs\": [ { \"name\": \"_from\", \"type\": \"address\" }, { \"name\": \"_to\", \"type\": \"address\" }, { \"name\": \"_value\", \"type\": \"uint256\" } ], \"name\": \"transferFromByAdmin\", \"outputs\": [ { \"name\": \"\", \"type\": \"bool\" } ], \"payable\": false, \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"constant\": false, \"inputs\": [ { \"name\": \"_spender\", \"type\": \"address\" }, { \"name\": \"_addedValue\", \"type\": \"uint256\" } ], \"name\": \"increaseApproval\", \"outputs\": [ { \"name\": \"\", \"type\": \"bool\" } ], \"payable\": false, \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"constant\": true, \"inputs\": [ { \"name\": \"_owner\", \"type\": \"address\" }, { \"name\": \"_spender\", \"type\": \"address\" } ], \"name\": \"allowance\", \"outputs\": [ { \"name\": \"\", \"type\": \"uint256\", \"value\": \"0\" } ], \"payable\": false, \"stateMutability\": \"view\", \"type\": \"function\" }, { \"constant\": false, \"inputs\": [ { \"name\": \"target\", \"type\": \"address\" }, { \"name\": \"freeze\", \"type\": \"bool\" } ], \"name\": \"freezeAccount\", \"outputs\": [], \"payable\": false, \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"constant\": false, \"inputs\": [ { \"name\": \"_u\", \"type\": \"address\" } ], \"name\": \"setNewOwner\", \"outputs\": [], \"payable\": false, \"stateMutability\": \"nonpayable\", \"type\": \"function\" }, { \"inputs\": [ { \"name\": \"initialSupply\", \"type\": \"uint256\", \"index\": 0, \"typeShort\": \"uint\", \"bits\": \"256\", \"displayName\": \"initial Supply\", \"template\": \"elements_input_uint\", \"value\": \"1000\" }, { \"name\": \"_name\", \"type\": \"string\", \"index\": 1, \"typeShort\": \"string\", \"bits\": \"\", \"displayName\": \"name\", \"template\": \"elements_input_string\", \"value\": \"WECTest\" }, { \"name\": \"_symbol\", \"type\": \"string\", \"index\": 2, \"typeShort\": \"string\", \"bits\": \"\", \"displayName\": \"symbol\", \"template\": \"elements_input_string\", \"value\": \"WECT\" }, { \"name\": \"_decimals\", \"type\": \"uint8\", \"index\": 3, \"typeShort\": \"uint\", \"bits\": \"8\", \"displayName\": \"decimals\", \"template\": \"elements_input_uint\", \"value\": \"1\" } ], \"payable\": true, \"stateMutability\": \"payable\", \"type\": \"constructor\" }, { \"anonymous\": false, \"inputs\": [ { \"indexed\": false, \"name\": \"target\", \"type\": \"address\" }, { \"indexed\": false, \"name\": \"frozen\", \"type\": \"bool\" } ], \"name\": \"FrozenFunds\", \"type\": \"event\" }, { \"anonymous\": false, \"inputs\": [ { \"indexed\": true, \"name\": \"owner\", \"type\": \"address\" }, { \"indexed\": true, \"name\": \"spender\", \"type\": \"address\" }, { \"indexed\": false, \"name\": \"value\", \"type\": \"uint256\" } ], \"name\": \"Approval\", \"type\": \"event\" }, { \"anonymous\": false, \"inputs\": [ { \"indexed\": true, \"name\": \"from\", \"type\": \"address\" }, { \"indexed\": true, \"name\": \"to\", \"type\": \"address\" }, { \"indexed\": false, \"name\": \"value\", \"type\": \"uint256\" } ], \"name\": \"Transfer\", \"type\": \"event\" } ]"
    //获取合约实例
	contract, err := rpc.Web3.Eth.Contract(abi);//rpc.Web3.Eth.NewContractWithInterface("", abi)
	if err != nil {
		ethlog.Error(err.Error())
		return
	}

    //调用合约前需要解锁账户
	result, error := rpc.Web3.Personal.UnlockAccount("0x6b1a125ab9895EcB7D7D5Eb550fFC1DFdf618a27", "123456789",120)
	if error != nil {
		ethlog.Error("unlock account error: " + error.Error())
		return ;
	}
	if !result {
		ethlog.Error("unlock account error")
		return ;
	}


	//调用合约方法 : transfer(address, uint256), 给地址:0xb0e6500e4303a11941bF07fA4dE41d05D1F621A2 发送 10个代币
	gprice , error := rpc.Web3.Eth.GetGasPrice();
	if error != nil {
		ethlog.Error("get gas price error: " + error.Error())
		return ;
	}
	//rpc.EthRpc.EthSendRawTransaction()

	transaction := new(dto.TransactionParameters)
	transaction.From = "0x6b1a125ab9895EcB7D7D5Eb550fFC1DFdf618a27" //调用合约账户
	transaction.To = "0xf78A0C798038ac0768043b1f18cA08f9e6507043"    //合约地址
	transaction.Gas = new(big.Int).SetUint64(60000)
    transaction.GasPrice = new(big.Int).SetUint64(gprice.ToUInt64());
	var amount uint32 = 10
   //调用合约
	txID, err :=contract.Send(transaction, "transfer","0xb0e6500e4303a11941bF07fA4dE41d05D1F621A2", amount)
	if err != nil {
		ethlog.Error(err.Error())
		return
	}

	ethlog.Debug(txID)
}
