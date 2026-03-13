package app

import (
	"encoding/json"
	"net/http"
	"fmt"
	"wallet/internal/model"
	"wallet/internal/pkg"
)

func initRouter() {
	http.HandleFunc("/wallets", CreateWallet)
	http.HandleFunc("/wallets/transfer", TransferWallet)
	http.HandleFunc("/wallets/", GetWallet)
	http.HandleFunc("/", defaultRouter)
}

func defaultRouter(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "default router: "+r.URL.Path, http.StatusMethodNotAllowed)
}

// CreateWalletResp create wallet response data
type CreateWalletResp struct {
	Wallet  *model.Wallet `json:",omitempty"`
	ErrCode *pkg.ErrorCode `json:",omitempty"`
}

func (r *CreateWalletResp) Marshal() string {
	buf, _ := json.Marshal(r)
	return string(buf)
}


// CreateWallet POST /wallets create a new wallet, Initial balance is zero
func CreateWallet(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("path: %+v\n", r.URL.Path)
	resp := CreateWalletResp{}
	pkg.Debug("CreateWallet entry, method: %s, url: %v", r.Method, r.URL.Path)
	if r.Method != http.MethodPost {
		resp.ErrCode = pkg.NewErrCode(http.StatusBadRequest, "/wallets only support POST method")
		http.Error(w, resp.Marshal(), http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	resp.Wallet, resp.ErrCode = model.CreateWallet()
	json.NewEncoder(w).Encode(resp)
	pkg.Info("CreateWallet resp wallet: %+v, errCode: %+v", resp.Wallet, resp.ErrCode)
}

// GetWallet GET /wallets/:id get wallet by id, Returns wallet ID and current balance
func GetWallet(w http.ResponseWriter, r *http.Request) {
	pkg.Debug("GetWallet entry, method: %s, url: %s", r.Method, r.URL.Path)
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	//pkg.Info("CreateWallet resp wallet: %+v, errCode: %+v", resp.Wallet, resp.Encode)
}

// TransferWallet POST /wallets/transfer Transfers an amount from one wallet to another
func TransferWallet(w http.ResponseWriter, r *http.Request) {
	pkg.Debug("GetWallet entry, method: %s, url: %s", r.Method, r.URL.Path)

	if r.Method != http.MethodPost {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	// pkg.Info("CreateWallet resp wallet: %+v, errCode: %+v", resp.Wallet, resp.Encode)
}
