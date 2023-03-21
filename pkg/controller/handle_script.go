package controller

import (
	inspectv1alpha1 "github.com/myoperator/inspectoperator/pkg/apis/inspect/v1alpha1"
	"github.com/myoperator/inspectoperator/pkg/workload/script"
	"k8s.io/klog/v2"
)

func handleScript(spec *inspectv1alpha1.InspectSpec) error {
	for _, task := range spec.Tasks {
		if task.Task.Type == "script" {
			err := script.RunLocalNode(task.Task.Source)
			if err != nil {
				klog.Error("create job err: ", err)
				return err
			}
		}
	}

	return nil
}
