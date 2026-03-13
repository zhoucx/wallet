package test

import (
	"sync"
	"testing"
	"time"
)

var (
	gBenchmarkResult = &BenchmarkResult{
		pool: make(map[string]*BenchMarkCount),
	}
)

func TestABenchmark(t *testing.T) {
	host := "http://localhost:9199/"

	go startServer(9199)
	time.Sleep(1 * time.Second)

	concurrency := 10     // 并发数
	maxRequestNunm := 100 // 每个并发调用多少次
	var wg sync.WaitGroup

	start := time.Now()
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			doCoroutineRequest(host, maxRequestNunm)
		}()
	}
	wg.Wait()
	uiseTime := time.Now().Sub(start)
	t.Logf("concurrency: %d, perRounteRun %d, totalUse: %d ms", concurrency, maxRequestNunm, uiseTime.Milliseconds())
	for k, v := range gBenchmarkResult.pool {
		t.Logf("%s: %+v", k, v)
	}
}

func doCoroutineRequest(host string, n int) {
	for i := 0; i < n; i++ {
		createResp, err := createWallet(host + "wallets")
		if err != nil || createResp.ErrCode != nil {
			gBenchmarkResult.AddCount("createWallet", 0, 1)
			continue
		}
		gBenchmarkResult.AddCount("createWallet", 1, 0)
		getResp, err := getWallet(host + "wallets/" + createResp.Wallet.Id)
		if err != nil || getResp.ErrCode != nil {
			gBenchmarkResult.AddCount("getWallet", 0, 1)
			continue
		}
		gBenchmarkResult.AddCount("getWallet", 1, 0)
	}
}

// 压力测试统计
type BenchMarkCount struct {
	SuccessNum int64
	FailNum    int64
}

type BenchmarkResult struct {
	pool map[string]*BenchMarkCount
	mtx  sync.Mutex
}

func (b *BenchmarkResult) AddCount(name string, successNum int64, failNum int64) {
	b.mtx.Lock()
	defer b.mtx.Unlock()
	count, ok := b.pool[name]
	if !ok {
		count = &BenchMarkCount{}
		b.pool[name] = count
	}
	count.SuccessNum += successNum
	count.FailNum += failNum
}
