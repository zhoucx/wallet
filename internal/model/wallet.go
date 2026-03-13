// package model
package model

import (
	"sync"
	"wallet/internal/pkg"
)

var (
	gWalletPool = &WalletPool{
		pool: make(map[string]Wallet),
	}
)

// Wallet user wallet
type Wallet struct {
	Id      string // 钱包ID
	UserId  string
	Balance int64 // 余额
}

type WalletPool struct {
	pool map[string]Wallet
	mtx  sync.Mutex
}

// CreateWallet create wallet, return walletId and balance is zero
func CreateWallet() (*Wallet, *pkg.ErrorCode) {
	wallet, errCode := gWalletPool.CreateWallet()
	if errCode != nil {
		return nil, errCode
	}
	return &wallet, nil
}

// GetWallet get wallet information
func GetWallet(id string) (*Wallet, *pkg.ErrorCode) {
	wallet, errCode := gWalletPool.GetWallet(id)
	if errCode != nil {
		return nil, errCode
	}
	return &wallet, nil
}

// TransferReq Transfers wallet amount request information
type TransferReq struct {
	SrcWalletId  string // 交易发起的钱包ID
	DestWalletId string // 交易的目的钱包ID
	Amount       int64  // 交易的金额
}

// TransferWallet Transfers an amount from one wallet to another
func TransferWallet(req TransferReq) (*Wallet, *pkg.ErrorCode) {
	wallet, errCode := gWalletPool.Transfer(req)
	if errCode != nil {
		return nil, errCode
	}
	return &wallet, nil
}

func AddWalletBalance(id string, amount int64) *pkg.ErrorCode {
	return gWalletPool.addBalance(id, amount)
}

func (wp *WalletPool) CreateWallet() (wallet Wallet, errCode *pkg.ErrorCode) {
	id, err := pkg.NewUid()
	if err != nil {
		return wallet, pkg.NewErrCode(pkg.LibErrCode, err.Error())
	}
	wallet.Id = id
	wallet.Balance = 0

	wp.mtx.Lock()
	defer wp.mtx.Unlock()
	wp.pool[wallet.Id] = wallet
	return wallet, nil
}

func (wp *WalletPool) GetWallet(id string) (wallet Wallet, errCode *pkg.ErrorCode) {
	wp.mtx.Lock()
	defer wp.mtx.Unlock()

	wallet, ok := wp.pool[id]
	if !ok {
		return wallet, pkg.NewWallentNotExistErr()
	}
	return wallet, nil
}

func (wp *WalletPool) Transfer(req TransferReq) (srcWallet Wallet, errCode *pkg.ErrorCode) {
	wp.mtx.Lock()
	defer wp.mtx.Unlock()

	srcWallet, ok := wp.pool[req.SrcWalletId]
	if !ok {
		return srcWallet, pkg.NewWallentNotExistErr()
	}
	destWallet, ok := wp.pool[req.DestWalletId]
	if !ok {
		return srcWallet, pkg.NewDestWallentNotExistErr()
	}
	if srcWallet.Balance < req.Amount {
		return srcWallet, pkg.NewAmountNotEnoughErr()
	}
	srcWallet.Balance -= req.Amount
	wp.pool[req.SrcWalletId] = srcWallet
	destWallet.Balance += req.Amount
	wp.pool[req.DestWalletId] = destWallet
	addTransferLog(req, srcWallet, destWallet)
	return srcWallet, nil
}

// 给账号加金额，用于测试或后台
func (wp *WalletPool) addBalance(id string, amount int64) *pkg.ErrorCode {
	wp.mtx.Lock()
	defer wp.mtx.Unlock()

	wallet, ok := wp.pool[id]
	if !ok {
		return pkg.NewWallentNotExistErr()
	}
	wallet.Balance += amount
	wp.pool[id] = wallet
	addBalanceLog(id, amount, wallet)
	return nil
}
