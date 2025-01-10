package controller

import (
	"context"
	v1 "ericsson.com/resource-modif-annotations/api/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"time"
)

func (r *ResourceModifierReconciler) initResourceModifierStatus(resource v1.ResourceModifier) {
	resource.Status.Conditions = make(map[string]string)
}

func (r *ResourceModifierReconciler) updateStatusSuccess(resource v1.ResourceModifier, reason string) (ctrl.Result, error) {
	resource.Status.SuccessfulStatus(reason)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err := r.Client.Update(ctx, &resource)
	if err != nil {

	}

	return ctrl.Result{}, nil
}
