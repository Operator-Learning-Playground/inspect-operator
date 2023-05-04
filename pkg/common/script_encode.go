package common

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"k8s.io/klog/v2"
)

// EncodeScript 加密脚本
func EncodeScript(str string) string {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)

	_, err := gz.Write([]byte(str))
	if err != nil {
		klog.Error(err)
		return ""
	}

	// 需要关掉
	err = gz.Close()
	if err != nil {
		klog.Error(err)
		return ""
	}

	return base64.StdEncoding.EncodeToString(buf.Bytes())
}
