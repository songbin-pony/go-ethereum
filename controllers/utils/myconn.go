package utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
)

//本地以太坊私链的某一个账户，这个账户信息在私链数据目录下的keystore目录里面
const key = `{"address":"9c1e470eb98e857064a7029e3f3cd951344335ab","crypto":{"cipher":"aes-128-ctr","ciphertext":"2057b5b63e2d6d3222632bb11f7e1387551ff713d9471337dfee1dc740c69806","cipherparams":{"iv":"64f02e08dcc7d534e5e1a4ce9d2e22e1"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":4096,"p":6,"r":8,"salt":"2e6af5cebbd092939d852295e9bf2325bc2b4efa17a4327293c366b212bcef7c"},"mac":"b5cca716238d3579f88bac95faf7cf676fc8a115bd40b1a128185306ae485bc5"},"id":"2bb0daa4-0b77-49bb-9b80-eda0a095fbe0","version":3}`
const addr = "0x0D99527c81Aa312ac5Af8BE244d9D8c01AF2fd49"
func GetContract() *TestStructSession{
	client, err1 := ethclient.Dial("http://192.168.209.128:8545")
	if err1 != nil {
		fmt.Printf("不能连接到本地节点: %v", err1)
	}
	// 将字符串地址转为common.Address
	commonAddr := common.HexToAddress(addr)
	// 创建合约对象
	contractObj, err2 := NewTestStruct(commonAddr, client)
	if err2 != nil {
		fmt.Printf("创建合约对象失败: %v", err2)
	}
	auth, err2 := bind.NewTransactorWithChainID(strings.NewReader(key), "123",big.NewInt(1337))
	if err2 != nil {
		fmt.Printf("解锁账户失败: %v", err2)
	}
	temp:=TestStructSession{
		Contract:     contractObj,
		CallOpts:     bind.CallOpts{Pending: true},
		TransactOpts: bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			Value:    nil,
		},
	}
	return &temp
}
