package auth

import (
	"bytes"
	"encoding/gob"
	"log"
	"rpc"
	"sync"
)

// AuthClient rpc client
type AuthClient struct {
	t       *rpc.TcpClient
	timeout int
}

var t *rpc.TcpClient
var once sync.Once

// NewAuthClient : return auth client
func NewAuthClient(address string, timeout int) *AuthClient {
	once.Do(func() {
		t = rpc.NewTcpClient(address)
	})

	return &AuthClient{t, timeout}
}

// LoginByPass : login by pass, shoud be generated by code
func (c *AuthClient) LoginByPass(req *LoginByPassReq, rsp *LoginByPassRsp) int {
	// encode req
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	enc.Encode(req)

	// call
	ret, recvBuf := t.Invoke("AuthSvr", "LoginByPass", buf.Bytes(), c.timeout)
	if ret != 0 {
		return int(ret)
	}

	// decode rsp
	buf = new(bytes.Buffer)
	buf.Write(recvBuf)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(rsp)
	if err != nil {
		log.Println(err)
		return rpc.PkgDecodeErr
	}

	return 0
}

// LoginByToken : login by token, shoud be generated by code
func (c *AuthClient) LoginByToken(req *LoginByTokenReq, rsp *LoginByTokenRsp) int {
	// encode req
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	enc.Encode(req)

	// call
	ret, recvBuf := t.Invoke("AuthSvr", "LoginByToken", buf.Bytes(), c.timeout)
	if ret != 0 {
		return int(ret)
	}

	// decode rsp
	buf = new(bytes.Buffer)
	buf.Write(recvBuf)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(rsp)
	if err != nil {
		log.Println(err)
		return rpc.PkgDecodeErr
	}

	return 0
}
