package controller

import (
	v1 "ericsson.com/resource-modif-annotations/api/v1"
	ctrl "sigs.k8s.io/controller-runtime"
)

func (r *ResourceModifierReconciler) initResourceModifierStatus(resource v1.ResourceModifier) {
	resource.Status.Conditions = make(map[string]string)
}

func (r *ResourceModifierReconciler) updateStatusSuccess(resource v1.ResourceModifier, reason string) (ctrl.Result, error) {
	return ctrl.Result{}, nil
}
