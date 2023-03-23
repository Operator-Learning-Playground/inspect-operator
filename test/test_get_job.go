package main

import (
	"context"
	"fmt"
	. "github.com/myoperator/inspectoperator/pkg/workload"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

func main() {
	r, err := ClientSet.BatchV1().Jobs("default").Get(context.Background(), "res", v1.GetOptions{})

	if err != nil && errors.IsNotFound(err) {
		klog.Error("not found error ", err)
		return
	} else if err != nil {
		klog.Error("get job error ", err)
		return
	}
	fmt.Println(r)
}
