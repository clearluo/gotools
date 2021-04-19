package pool

import (
	"bytes"
	"crypto/md5"
	"net/http"
	"sync"
	"time"
)

var (
	Md5ObjPool      *sync.Pool
	httpClientPool  *sync.Pool
	bytesBufferPool *sync.Pool
)

func init() {
	Md5ObjPool = &sync.Pool{
		New: func() interface{} {
			return md5.New()
		},
	}
	httpClientPool = &sync.Pool{
		New: func() interface{} {
			t := http.DefaultTransport.(*http.Transport).Clone()
			t.MaxIdleConns = 100
			t.MaxConnsPerHost = 100
			t.MaxIdleConnsPerHost = 100
			return &HttpClient{
				Client: &http.Client{Timeout: time.Second * 10, Transport: t},
			}
		},
	}
	bytesBufferPool = &sync.Pool{
		New: func() interface{} {
			return &BytesBuff{
				Buffer: bytes.NewBuffer(make([]byte, 0, 1024)),
			}
		},
	}
}

type BytesBuff struct {
	*bytes.Buffer
}

func (b *BytesBuff) Free() {
	b.Reset()
	bytesBufferPool.Put(b)
}

func NewMyBytesBuff() *BytesBuff {
	return bytesBufferPool.Get().(*BytesBuff)
}

type HttpClient struct {
	*http.Client
}

func (b *HttpClient) Free() {
	httpClientPool.Put(b)
}

func NewHttpClient() *HttpClient {
	return httpClientPool.Get().(*HttpClient)
}
