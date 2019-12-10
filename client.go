package main

import (
    "log"
    "./util"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    for i := 1; i < 2; i++ {
        wg.Add(1)
        go func (i int) {
            //rec := util.Client("m1", false, []byte("test"), util.File{Name: "upload.zip", Path: "/tmp/upload/1.zip"})
            rec := util.Client("m1", false, []byte("sssssssssssssss1111111111"))
            log.Printf("%v, client rec: %v", i, string(rec.Data))
            wg.Done()
        }(i)
    }
    wg.Wait()
}
