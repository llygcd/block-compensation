//init client from clientPool.
//client is httpClient of tendermint

package pool

import (
	"context"
	"errors"
	commonPool "github.com/jolestar/go-commons-pool"
	"github.com/llygcd/block-compensation/config"
	"github.com/sirupsen/logrus"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"strings"
	"sync"
	"time"
)

var (
	poolObject        *commonPool.ObjectPool
	poolFactory       PoolFactory
	ctx               = context.Background()
	isJsonRpcProtocol bool
)

func Init(conf *config.Config) {
	var (
		syncMap sync.Map
	)
	isJsonRpcProtocol = conf.Server.IsJsonRpcProtocol
	nodeUrls := strings.Split(conf.Server.NodeUrls, ",")
	for _, url := range nodeUrls {
		key := generateId(url)
		endPoint := EndPoint{
			Address:   url,
			Available: true,
		}

		syncMap.Store(key, endPoint)
	}

	poolFactory = PoolFactory{
		peersMap: syncMap,
	}

	config := commonPool.NewDefaultPoolConfig()
	config.MaxTotal = conf.Server.MaxConnectionNum
	config.MaxIdle = conf.Server.InitConnectionNum
	config.MinIdle = conf.Server.InitConnectionNum
	config.TestOnBorrow = true
	config.TestOnCreate = true
	config.TestWhileIdle = true

	poolObject = commonPool.NewObjectPool(ctx, &poolFactory, config)
	poolObject.PreparePool(ctx)
}

// get client from pool
func GetClient() *Client {
	c, err := poolObject.BorrowObject(ctx)
	for err != nil {
		logrus.Errorf("GetClient failed,will try again after 3 seconds,%s", err.Error())
		time.Sleep(3 * time.Second)
		c, err = poolObject.BorrowObject(ctx)
	}

	return c.(*Client)
}

// release client
func (c *Client) Release() {
	err := poolObject.ReturnObject(ctx, c)
	if err != nil {
		logrus.Error(err.Error())
	}
}

func (c *Client) HeartBeat() error {
	if isJsonRpcProtocol {
		return c.Jrpc.HeartBeat(context.Background())
	}
	http := c.HTTP
	_, err := http.Health(context.Background())
	return err
}

func (c *Client) Status(ctx context.Context) (*ctypes.ResultStatus, error) {
	if isJsonRpcProtocol {
		return c.Jrpc.Status(ctx)
	}
	http := c.HTTP
	return http.Status(ctx)

}

func ClosePool() {
	poolObject.Close(ctx)
}

func GetClientWithTimeout(timeout time.Duration) (*Client, error) {
	c := make(chan interface{})
	errCh := make(chan error)
	go func() {
		client, err := poolObject.BorrowObject(ctx)
		if err != nil {
			errCh <- err
		} else {
			c <- client
		}
	}()
	select {
	case res := <-c:
		return res.(*Client), nil
	case res := <-errCh:
		return nil, res
	case <-time.After(timeout):
		return nil, errors.New("rpc node timeout")
	}
}
