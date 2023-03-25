package controller

import (
	inspectv1alpha1 "github.com/myoperator/inspectoperator/pkg/apis/inspect/v1alpha1"
	"github.com/myoperator/inspectoperator/pkg/workload/job"
	"k8s.io/klog/v2"
)

func handleImage(spec *inspectv1alpha1.InspectSpec) error {

	for _, task := range spec.Tasks {
		if task.Task.Type == "image" {
			err := job.CreateJob(&task.Task, task.Task.Source)
			if err != nil {
				klog.Error("create job err: ", err)
				return err
			}
			// 启动定时任务
			ticker := job.NewTickerTask(30, job.GetJobStatus, task.Task.TaskName)
			klog.Info("启动定时任务")
			go ticker.Start()
		}
	}

	return nil

}
