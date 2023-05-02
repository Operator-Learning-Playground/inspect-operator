package controller

import (
	"context"
	inspectv1alpha1 "github.com/myoperator/inspectoperator/pkg/apis/inspect/v1alpha1"
	"github.com/myoperator/inspectoperator/pkg/sysconfig"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
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
		if client.IgnoreNotFound(err) != nil {
			klog.Error("get inspect error: ", err)
			return reconcile.Result{}, err
		}
		// 如果未找到的错误，不再进入调协
		return reconcile.Result{}, nil
	}
	klog.Info(inspect)

	// update config
	err = sysconfig.AppConfig(inspect)
	if err != nil {
		return reconcile.Result{}, err
	}

	// 使用CreateOrUpdate处理业务逻辑
	mutateInspectRes, err := controllerutil.CreateOrUpdate(ctx, r.Client, inspect, func() error {
		// FIXME: 如何解决重复进入的问题
		klog.Info("is in...?")
		// 业务逻辑
		err = handleImage(&inspect.Spec)
		if err != nil {
			klog.Error("handle image error: ", err)
			return err
		}
		err = handleScript(&inspect.Spec)
		if err != nil {
			klog.Error("handle script error: ", err)
			return err
		}
		return nil
	})
	if err != nil {
		klog.Error("CreateOrUpdate error ", " Inspect ", err)
		return reconcile.Result{}, err
	}
	klog.Info("CreateOrUpdate ", "Inspect ", mutateInspectRes)

	return reconcile.Result{}, nil
}

// 使用controller-runtime 需要注入的client
func (r *InspectController) InjectClient(c client.Client) error {
	r.Client = c
	return nil
}

// TODO: 删除逻辑并未处理
