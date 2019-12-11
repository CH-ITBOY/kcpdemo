package util

import (
    "crypto/sha1"
    "github.com/xtaci/kcp-go"
    "golang.org/x/crypto/pbkdf2"
)

var Laddr = "127.0.0.1:12345"
var KeyBlock, _ = kcp.NewAESBlockCrypt(pbkdf2.Key([]byte("connect pass"), []byte("connect salt"), 1024, 32, sha1.New))

func SetKcp (conn *kcp.UDPSession) {
    conn.SetNoDelay(1, 10, 2, 1)
}
