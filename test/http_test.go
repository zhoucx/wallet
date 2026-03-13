package test

import (
	"testing"
	"net/http"
	"io/ioutil"
	"wallet/internal/app"
	"encoding/json"
	"time"
)

func TestCreateWallet(t *testing.T) {
	go startServer(9091)
	time.Sleep(1*time.Second)

	httpResp, err := http.Post("http://localhost:9091/wallets", "application/json", nil)
	if err != nil {
		t.Errorf("http.Get failed: %v", err)
		return 
	}
	defer httpResp.Body.Close()

    body, err := ioutil.ReadAll(httpResp.Body)
    if err != nil {
        t.Errorf("read httpResp.body failed: %v", err)
		return 
    }
	t.Logf("body: %s", string(body))
    createResp := app.CreateWalletResp{}
	err = json.Unmarshal(body, &createResp)
	if err != nil {
		t.Errorf("json.Unmarshal createResp failed: %v", err)
		return 
	}
	if createResp.ErrCode != nil {
		t.Errorf("createResp have errCode")
		return 
	}
	if createResp.Wallet == nil || createResp.Wallet.Id == "" || createResp.Wallet.Balance != 0 {
		t.Errorf("createResp wallet invalid")
		return 
	}
}




func startServer(port int) {
	app.StartServer(port)
}