package script

import (
	"bytes"
	"github.com/myoperator/inspectoperator/pkg/common"
	"log"
	"os/exec"
)

func RunLocalNode(script string) error {
	path := common.GetWd()
	cmd := exec.Command("bash", path + "/script/" + script)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	err := cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	log.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	if err != nil {
		log.Printf("cmd.Run() failed with %s\n", err)
		return err
	}

	return nil
}