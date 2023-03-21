package job

import (
	"context"
	"encoding/json"
	"fmt"
	inspectv1alpha1 "github.com/myoperator/inspectoperator/pkg/apis/inspect/v1alpha1"
	"github.com/myoperator/inspectoperator/pkg/common"
	. "github.com/myoperator/inspectoperator/pkg/workload"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
	"time"
)

// CreateJob
// @see https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/job-v1/#create-create-a-job
func CreateJob(task *inspectv1alpha1.Task, image string) error {

	job := jobSpec(task, image)

	j, _ := json.Marshal(job)
	fmt.Printf("Create jobs with params: %s\n", j)

	// create cronjob
	jobResult, err := ClientSet.BatchV1().Jobs("default").Create(context.TODO(), job, metav1.CreateOptions{})
	fmt.Printf("Create job complete, job result: %v\n", jobResult)
	return err
}

func jobSpec(task *inspectv1alpha1.Task, image string) *batchv1.Job {
	fmt.Printf("Create k8s job params: taskName=%v\n", task.TaskName)
	// init
	taskName := fmt.Sprintf("inspect-manual-task-%v-%d", task.TaskName, time.Now().Unix())
	labels := map[string]string{"taskType": "inspect", "app": taskName}
	return &batchv1.Job{
		// metadata
		ObjectMeta: metav1.ObjectMeta{
			Name:   taskName,
			Labels: labels,
		}, Spec: batchv1.JobSpec{
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: v1.PodSpec{
					// containers
					Containers: []v1.Container{{
						Name:            "default",
						Image:           image,
						ImagePullPolicy: v1.PullIfNotPresent,
					}},
					// restart policy
					RestartPolicy:      v1.RestartPolicyNever,
					ServiceAccountName: common.DefaultServiceAccount,
				},
			},
		},
	}
}

// DeleteJob
// @see https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/job-v1/#delete-delete-a-job
func DeleteJob(tasks []*inspectv1alpha1.Task) error {

	ids := make([]string, len(tasks))
	for i, task := range tasks {
		ids[i] = task.TaskName
	}
	fmt.Printf("Delete k8s job, task ids: %#v\n", ids)
	listOptions := metav1.ListOptions{
		LabelSelector: fmt.Sprintf("taskId in (%s)", strings.Join(ids, ",")),
	}
	foreground := metav1.DeletePropagationForeground
	deleteOptions := metav1.DeleteOptions{PropagationPolicy: &foreground}
	return ClientSet.BatchV1().Jobs("default").DeleteCollection(context.TODO(), deleteOptions, listOptions)
}