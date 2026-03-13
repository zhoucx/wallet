package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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
	Wallet  *model.Wallet  `json:",omitempty"`
	ErrCode *pkg.ErrorCode `json:",omitempty"`
}

func (r *CreateWalletResp) MarshalToString() string {
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
		http.Error(w, resp.MarshalToString(), http.StatusMethodNotAllowed)
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
	resp := CreateWalletResp{}
	if r.Method != http.MethodGet {
		resp.ErrCode = pkg.NewErrCode(http.StatusBadRequest, "/wallets/:id only support GET method")
		http.Error(w, resp.MarshalToString(), http.StatusMethodNotAllowed)
		return
	}
	items := strings.Split(r.URL.Path, "/")
	if len(items) != 3 && len(items[2]) == 0 {
		resp.ErrCode = pkg.NewErrCode(http.StatusBadRequest, "url.Path invalid")
		http.Error(w, resp.MarshalToString(), http.StatusMethodNotAllowed)
		return
	}
	walletId := items[2]
	resp.Wallet, resp.ErrCode = model.GetWallet(walletId)
	json.NewEncoder(w).Encode(resp)
	pkg.Info("GetWallet resp wallet: %+v, errCode: %+v", resp.Wallet, resp.ErrCode)
}

// TransferWallet POST /wallets/transfer Transfers an amount from one wallet to another
func TransferWallet(w http.ResponseWriter, r *http.Request) {
	pkg.Debug("TransferWallet entry, method: %s, url: %s", r.Method, r.URL.Path)

	resp := CreateWalletResp{}
	if r.Method != http.MethodPost {
		resp.ErrCode = pkg.NewErrCode(http.StatusBadRequest, "/walletstransfer only support POST method")
		http.Error(w, resp.MarshalToString(), http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resp.ErrCode = pkg.NewErrCode(pkg.LibErrCode, "read http.Request body failed")
		http.Error(w, resp.MarshalToString(), http.StatusBadRequest)
		return
	}
	fmt.Printf("TransferWallet request body: %s\n", string(body))
	transferReq := model.TransferReq{}
	err = json.Unmarshal(body, &transferReq)
	if err != nil {
		resp.ErrCode = pkg.NewErrCode(pkg.LibErrCode, "unmarshal http.Request body failed")
		http.Error(w, resp.MarshalToString(), http.StatusBadRequest)
		return
	}
	if transferReq.SrcWalletId == "" || transferReq.DestWalletId == "" || transferReq.Amount <= 0 || transferReq.SrcWalletId == transferReq.DestWalletId {
		resp.ErrCode = pkg.NewErrCode(pkg.LibErrCode, fmt.Sprintf("transferReq invalid: %+v", transferReq))
		http.Error(w, resp.MarshalToString(), http.StatusBadRequest)
		return
	}
	resp.Wallet, resp.ErrCode = model.TransferWallet(transferReq)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	pkg.Info("TransferWallet resp wallet: %+v, errCode: %+v", resp.Wallet, resp.ErrCode)
}
