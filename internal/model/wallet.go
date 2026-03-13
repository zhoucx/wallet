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
func CreateWallet() (wallet Wallet, errCode *pkg.ErrorCode) {
	return gWalletPool.CreateWallet()
}

// GetWallet get wallet information
func GetWallet(id string) (wallet Wallet, errCode *pkg.ErrorCode) {
	return gWalletPool.GetWallet(id)
}

// TransferReq Transfers wallet amount request information
type TransferReq struct {
	SrcWalletId  string // 交易发起的钱包ID
	DestWalletId string // 交易的目的钱包ID
	Amount       int64  // 交易的金额
}

// TransferWallet Transfers an amount from one wallet to another
func TransferWallet(req TransferReq) *pkg.ErrorCode {
	return gWalletPool.Transfer(req)
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

func (wp *WalletPool) Transfer(req TransferReq) *pkg.ErrorCode {
	wp.mtx.Lock()
	defer wp.mtx.Unlock()

	srcWallet, ok := wp.pool[req.SrcWalletId]
	if !ok {
		return pkg.NewWallentNotExistErr()
	}
	destWallet, ok := wp.pool[req.DestWalletId]
	if !ok {
		return pkg.NewDestWallentNotExistErr()
	}
	if srcWallet.Balance < req.Amount {
		return pkg.NewAmountNotEnoughErr()
	}
	srcWallet.Balance -= req.Amount
	wp.pool[req.SrcWalletId] = srcWallet
	destWallet.Balance += req.Amount
	wp.pool[req.DestWalletId] = destWallet
	addTransferLog(req, srcWallet, destWallet)
	return nil
}
