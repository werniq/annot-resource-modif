package controller

import "sigs.k8s.io/controller-runtime/pkg/client"

// executeRemoveAnyFinalizerAnnotation
// This function removes any finalizers from the resource, if there were one.
func (r *ResourceModifierReconciler) executeRemoveAnyFinalizerAnnotation(resource client.Object) error {
	if resource.GetFinalizers() == nil {
		return nil
	}

	resource.SetFinalizers(nil)

	return nil
}
