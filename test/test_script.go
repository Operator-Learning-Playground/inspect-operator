package main

import (
	"bytes"
	"k8s.io/klog/v2"

	"log"
	"os/exec"
)

func main() {
	//path := common.GetWd()
	cmd := exec.Command("sh", "./script/test.sh")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	klog.Info("out:\n%s\nerr:\n%s\n", outStr, errStr)
	if err != nil {
		klog.Info("cmd.Run() failed with %s\n", err)
		return
	}

	return

}
