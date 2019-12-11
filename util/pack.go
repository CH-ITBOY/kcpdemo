package util

import (
    "encoding/gob"
    "bytes"
    "github.com/xtaci/kcp-go"
    "log"
    "io"
)

type File struct {
    Name string
    Path string
    Size int64
}

type Header struct {
    Async bool
    Method string
}

type Body struct {
    Header Header
    Data []byte
    File []File
}

func ReadKpc (conn *kcp.UDPSession, size int64) []byte {
    buf := make([]byte, size)
    if _, err := io.ReadFull(conn, buf); err == nil {
        return buf
    } else {
        log.Println(err)
    }   
    return buf
    /*buf := make([]byte, size)
    rec := make([]byte, 0)
    var reclength int64 = 0
    for {
        n, err := conn.Read(buf)
        if err != nil {
            log.Println(err)
            return rec
        }
        rec = append(rec, buf[:n]...)
        reclength = reclength + int64(n)
        if (reclength == size) {
            break
        }
    }
    return rec*/
}

func Unpack (conn *kcp.UDPSession) Body {
    bodySizeBuf := ReadKpc(conn, 8)
    bodySize := BytesToInt64(bodySizeBuf)
    body := ReadKpc(conn, bodySize)
    bodyPtr := bytes.NewBuffer(body)
    decoder := gob.NewDecoder(bodyPtr)
    var b Body
    if err := decoder.Decode(&b); err != nil {
        log.Println(err)
        return Body{}
    }
    for _, f := range b.File {
       fname := f.Name
       fsize := f.Size
       fcontent := ReadKpc(conn, fsize)
       UploadFile(fname, fcontent) 
    }
    return b
}

// 封包
func Pack (data []byte, files ...File) []byte{
    return PackWithHeader(Header{}, data, files...)
}

func PackWithHeader (header Header, data []byte, files ...File) []byte{
    filesByte := make([]byte, 0)
    for index, f := range files {
        fileContent, _ := ReadFile(f.Path)
        filesByte = append(filesByte, fileContent...)
        files[index].Size, _ = FileSize(f.Path)
    }
    body := Body{
        Header: header,
        Data: data,
        File: files,
    }
    var buf bytes.Buffer
    encoder := gob.NewEncoder(&buf)
    if err := encoder.Encode(&body); err != nil {
        log.Println(err)
        return []byte{}
    }
    buflen := len(buf.Bytes())
    d := append(Int64ToBytes(int64(buflen)), (buf.Bytes())...)
    return append(d, filesByte...)
}
