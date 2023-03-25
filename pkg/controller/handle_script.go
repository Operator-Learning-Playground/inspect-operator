package controller

import (
	inspectv1alpha1 "github.com/myoperator/inspectoperator/pkg/apis/inspect/v1alpha1"
	"github.com/myoperator/inspectoperator/pkg/workload/script"
	"k8s.io/klog/v2"
)

func handleScript(spec *inspectv1alpha1.InspectSpec) error {
	for _, task := range spec.Tasks {
		if task.Task.Type == "script" {
			// FIXME: 目前就是异步执行，后面controller就不管。
			go func() {
				err := script.RunLocalNode(task.Task.Source)
				if err != nil {
					klog.Error("create script err: ", err)
					return
				}

			}()

		}
	}

	return nil
}
