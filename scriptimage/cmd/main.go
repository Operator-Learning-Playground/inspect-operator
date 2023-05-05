package main

import (
	"k8s.io/klog/v2"
	"os"
	"scriptimage/pkg/decode"
	"scriptimage/pkg/execute"
)

func main() {
	// 1. 从环境变量中取
	scriptEncryptedString := os.Getenv("script")

	err := decode.WriteStringToFile(scriptEncryptedString)
	if err != nil {
		klog.Error("write err:", err)
		return
	}

	err = decode.GenEncodeFile()
	if err != nil {
		klog.Error("decode err:", err)
		return
	}
	if os.Getenv("script_location") == "remote" {
		err = execute.RunRemoteNode(os.Getenv("user"), os.Getenv("password"), os.Getenv("ip"))
		if err != nil {
			klog.Error("execute err:", err)
			return
		}
	} else {
		klog.Info("run inspect task in local...")
		err = execute.RunLocalNode()
		if err != nil {
			klog.Error("execute err:", err)
			return
		}
	}


	klog.Info("finished script inspect task...")
}
