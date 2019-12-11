package util 

import (
    "os"
    "path"
    "io/ioutil"
)

var UploadPath = "/tmp/upload"

// 路径是否存在
func PathExists(path string) (bool, error) {
    fi, err := os.Stat(path)
    if err != nil {
        return false, nil
    }
    return fi.IsDir(), err
}

func FileSize (fileRealPath string) (int64, error) {
    f, err := os.Stat(fileRealPath)
    if err != nil {
        return 0, err
    }
    return f.Size(), nil
}

func UploadFile (filename string, data []byte) error {
    exist, _ := PathExists(UploadPath)
    if !exist {
        // 创建文件夹
        os.MkdirAll(UploadPath, os.ModePerm)
    }
    return ioutil.WriteFile(path.Join(UploadPath, filename), data, 0644)
}

func ReadFile (fileRealPath string) ([]byte, error) {
   return ioutil.ReadFile(fileRealPath)
}
