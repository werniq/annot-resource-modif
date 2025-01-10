package controller

import (
	"context"
	annotresourcemodifv1 "ericsson.com/resource-modif-annotations/api/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"slices"
	"time"
)

const (
	// successRemovingFinalizers
	successRemovingFinalizers = "Successfully removed finalizers"
)

// executeRemoveAnyFinalizerAnnotation
// This function removes any finalizers from the resource, if there were one.
func (r *ResourceModifierReconciler) executeRemoveAnyFinalizerAnnotation(resource client.Object,
	rm annotresourcemodifv1.ResourceModifier) error {
	if resource.GetFinalizers() == nil {
		return nil
	}

	resource.SetFinalizers(nil)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err := r.Client.Update(ctx, resource)
	if err != nil {
		return err
	}

	err = r.updateStatusSuccess(rm, successRemovingFinalizers)
	if err != nil {
		updateErr := r.updateErrorStatus(rm, err.Error())
		if updateErr != nil {
			return updateErr
		}
		return err
	}

	return nil
}

// executeAddFinalizer adds provided finalizer to the target resource.
func (r *ResourceModifierReconciler) executeAddFinalizer(resource client.Object,
	rm annotresourcemodifv1.ResourceModifier, finalizer string) error {
	existentFinalizers := resource.GetFinalizers()
	if slices.Contains(existentFinalizers, finalizer) {
		return nil
	}
	existentFinalizers = append(existentFinalizers, finalizer)

	resource.SetFinalizers(existentFinalizers)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err := r.Client.Update(ctx, resource)
	if err != nil {
		return err
	}

	err = r.updateStatusSuccess(rm, successRemovingFinalizers)
	if err != nil {
		updateErr := r.updateErrorStatus(rm, err.Error())
		if updateErr != nil {
			return updateErr
		}
		return err
	}

	return nil
}
