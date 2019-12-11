package util

import (
    "log"
    "github.com/xtaci/kcp-go"
    "reflect"
    "sync"
)

func call(m map[string]interface{}, methodName string, args ...interface{}) ([]byte) {
    inputs := make([]reflect.Value, len(args))
    for i, _ := range args {
        inputs[i] = reflect.ValueOf(args[i])
    }
    res := reflect.ValueOf(m[methodName]).Call(inputs)
    return res[0].Bytes()
    //return reflect.TypeOf(res[0]).Bits()
}

func Server() {
    if listener, err := kcp.ListenWithOptions(Laddr, KeyBlock, 10, 3); err == nil {
        var mu = sync.Mutex{}
        for {
            s, err := listener.AcceptKCP()
            if err != nil {
                log.Println(err)
            }
            SetKcp(s)
            go handle(s, &mu)
        }
    } else {
        log.Println(err)
    }
}
func handle(conn *kcp.UDPSession, mu *sync.Mutex) {
    body := Unpack(conn)
    if body.Header.Async == true {
        // 并发执行 
        res := call(mclass, body.Header.Method, body)
        conn.Write(Pack([]byte(res)))
    } else {
        // 顺序执行
        mu.Lock()
        defer mu.Unlock()
        res := call(mclass, body.Header.Method, body)
        conn.Write(Pack([]byte(res)))
    }
}
