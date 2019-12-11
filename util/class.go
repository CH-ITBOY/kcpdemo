package util

var mclass = map[string]interface{}{}

func M (methodName string, mfunc interface{}) {
    mclass[methodName] = mfunc
}
