package util

import (
    "log"
    "github.com/xtaci/kcp-go"
)

func Client(method string, async bool, data []byte, files ...File) Body {
    if sess, err := kcp.DialWithOptions(Laddr, KeyBlock, 10, 3); err == nil {
        SetKcp(sess)
        data := PackWithHeader(Header{Async:async, Method: method}, data, files...)
        if _, err := sess.Write(data); err == nil {
            rec := Unpack(sess)
            return rec
        }
    } else {
        log.Println(err)
    }
    return Body{}
}
