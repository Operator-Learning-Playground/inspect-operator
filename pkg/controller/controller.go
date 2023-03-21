package controller

import (
	"context"
	inspectv1alpha1 "github.com/myoperator/inspectoperator/pkg/apis/inspect/v1alpha1"
	"github.com/myoperator/inspectoperator/pkg/sysconfig"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type InspectController struct {
	client.Client

}

func NewInspectController() *InspectController {
	return &InspectController{}
}

// Reconcile 调协loop
func (r *InspectController) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {

	inspect := &inspectv1alpha1.Inspect{}
	err := r.Get(ctx, req.NamespacedName, inspect)
	if err != nil {
		return reconcile.Result{}, err
	}
	klog.Info(inspect)

	err = sysconfig.AppConfig(inspect)
	if err != nil {
		return reconcile.Result{}, nil
	}

	// 业务逻辑
	err = handleImage(&inspect.Spec)
	if err != nil {
		return reconcile.Result{}, nil
	}
	err = handleScript(&inspect.Spec)
	if err != nil {
		return reconcile.Result{}, nil
	}

	return reconcile.Result{}, nil
}

// 使用controller-runtime 需要注入的client
func(r *InspectController) InjectClient(c client.Client) error {
	r.Client = c
	return nil
}

// TODO: 删除逻辑并未处理

