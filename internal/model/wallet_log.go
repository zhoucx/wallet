package model

import (
	"fmt"
	"time"
)

// TransferLog Wallet transaction log information,can query transaction records based on the wallet ID
type TransferLog struct {
	WalletId         string // 钱包ID, 通过这个ID可以查询到所有的交易记录
	TransferWalletId string // 参与交易的另一个钱包ID
	Amount           int64  // 交易金额
	Message          string
	Balance          int64 // 余额
	TimeStr          string
}

// 交易日志记录, 正常情况下应该记录数据库, 这里直接打日志好了
func addTransferLog(req TransferReq, srcWallet Wallet, destWallet Wallet) {
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	srcLog := TransferLog{
		WalletId:         req.SrcWalletId,
		TransferWalletId: req.DestWalletId,
		Amount:           -req.Amount,
		Message:          "transfer to other",
		Balance:          srcWallet.Balance,
		TimeStr:          timeStr,
	}
	fmt.Printf("%+v\n", srcLog)
	destLog := TransferLog{
		WalletId:         req.DestWalletId,
		TransferWalletId: req.SrcWalletId,
		Amount:           req.Amount,
		Message:          "transfer from other",
		Balance:          destWallet.Balance,
		TimeStr:          timeStr,
	}
	fmt.Printf("%+v\n", destLog)
}
