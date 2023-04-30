package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Inspect
type Inspect struct {
	metav1.TypeMeta `json:",inline"`

	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec InspectSpec `json:"spec,omitempty"`

	Status InspectStatus `json:"status,omitempty"`
}

type InspectSpec struct {
	Tasks []Tasks `json:"tasks"`
}

// InspectStatus 任务完成状态：TODO 目前还没用到
type InspectStatus struct {
	Results []TaskRes `json:"results"`
}

type TaskRes struct {
	TaskName string `json:"task_name"`
	Res      string `json:"res"`
}

type Tasks struct {
	Task Task `json:"task"`
}

type RemoteIp struct {
	Ip       string `json:"ip"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type RemoteIps struct {
	RemoteIp RemoteIp `json:"remote_ip"`
}

type Task struct {
	TaskName       string      `json:"task_name"`
	Type           string      `json:"type"`
	Source         string      `json:"source"`
	ScriptLocation string      `json:"script_location"`
	RemoteIps      []RemoteIps `json:"remote_ips"`
	Restart        bool        `json:"restart"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// InspectList
type InspectList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Inspect `json:"items"`
}
