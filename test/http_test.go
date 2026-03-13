package test

import (
	"testing"
	"net/http"
	"io/ioutil"
	"wallet/internal/app"
	"encoding/json"
	"time"
	"fmt"
	"wallet/internal/pkg"
)

func TestCreateWallet(t *testing.T) {
	go startServer(9091)
	time.Sleep(1*time.Second)

	resp, err := createWallet("http://localhost:9091/wallets")
	if err != nil || resp == nil{
		t.Errorf("createWallet failed: %v", err)
		return 
	}
	if resp.ErrCode != nil {
		t.Errorf("createResp have errCode: %+v", resp.ErrCode)
		return 
	}
	if resp.Wallet == nil || resp.Wallet.Id == "" || resp.Wallet.Balance != 0 {
		t.Errorf("createResp wallet invalid")
		return 
	}
}

func TestGetWallet(t *testing.T) {
	go startServer(9092)
	time.Sleep(1*time.Second)
	// 获取不存在的钱包
	resp, err := getWallet("http://localhost:9092/wallets/001")
	if err != nil {
		t.Errorf("get wallet failed: %v", err)
		return
	}
	if resp.ErrCode == nil || resp.ErrCode.Code != pkg.WallentNotExistCode {
		t.Errorf("get not exist wallet return success")
		return 
	}
	// 创建一个钱包再获取
	createResp, err := createWallet("http://localhost:9092/wallets")
	if err != nil {
		t.Errorf("create wallet failed: %v", err)
		return
	}
	if createResp.ErrCode != nil || createResp.Wallet.Id == "" {
		t.Errorf("create wallet resp invalid： %s", createResp.MarshalToString())
		return 
	}
	getResp, err := getWallet("http://localhost:9092/wallets/"+createResp.Wallet.Id)
	if err != nil {
		t.Errorf("get wallet failed: %v", err)
		return
	}
	if getResp.Wallet == nil || getResp.ErrCode != nil {
		t.Errorf("get wallet resp invalid: %s", getResp.MarshalToString())
		return 
	}
	if getResp.Wallet.Id != createResp.Wallet.Id && getResp.Wallet.Balance != 0 {
		t.Errorf("get wallet resp id or balance invalid: %s", getResp.MarshalToString())
		return
	}
	t.Logf("get wallet: %s", getResp.MarshalToString())
}



func startServer(port int) {
	app.StartServer(port)
}

func createWallet(url string) (*app.CreateWalletResp, error) {
	httpResp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return nil, fmt.Errorf("http.Get failed: %v", err)
	}
	defer httpResp.Body.Close()

    body, err := ioutil.ReadAll(httpResp.Body)
    if err != nil {
		return nil, fmt.Errorf("read httpResp.body failed: %v", err)
    }
	fmt.Printf("createWallet resp body: %s\n", string(body))
    createResp := &app.CreateWalletResp{}
	err = json.Unmarshal(body, createResp)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal createResp failed: %v", err)
	}
	return createResp, nil
}

func getWallet(url string) (*app.CreateWalletResp, error) {
	httpResp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http.Get failed: %v", err)
	}
	defer httpResp.Body.Close()

    body, err := ioutil.ReadAll(httpResp.Body)
    if err != nil {
		return nil, fmt.Errorf("read httpResp.body failed: %v", err)
    }
	fmt.Printf("get wallet resp.body: %s\n", string(body))
    getResp := &app.CreateWalletResp{}
	err = json.Unmarshal(body, getResp)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal getResp failed: %v", err)
	}
	return getResp, nil
}