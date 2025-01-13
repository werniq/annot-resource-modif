package controller

import (
	"context"
	annotresourcemodifv1 "ericsson.com/resource-modif-annotations/api/v1"
	errs "errors"
	v1 "k8s.io/api/apps/v1"
	v2 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/remotecommand"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"slices"
	"strings"
	"time"
)

const (
	// successRemovingFinalizers
	successRemovingFinalizers = "Successfully removed finalizers"

	// successAddFinalizers
	successAddFinalizers = "Successfully added finalizers"

	// successAddLabel
	successAddLabel = "Successfully added label"

	// successDeploymentScale
	successDeploymentScale = "Successfully scaled deployment"

	// errTypeNotScalable
	errTypeNotScalable = "given Resource Type is not scalable"
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

	err = r.updateStatusSuccess(rm, successAddFinalizers)
	if err != nil {
		updateErr := r.updateErrorStatus(rm, err.Error())
		if updateErr != nil {
			return updateErr
		}
		return err
	}

	return nil
}

// executeAddLabel adds new label to the resource.
func (r *ResourceModifierReconciler) executeAddLabel(resource client.Object,
	rm annotresourcemodifv1.ResourceModifier, label string) error {
	labels := resource.GetLabels()
	s := strings.Split(label, ":")
	key, value := s[0], s[1]

	if _, exists := labels[key]; exists {
		return nil
	}
	labels[key] = value

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err := r.Client.Update(ctx, resource)
	if err != nil {
		return err
	}

	err = r.updateStatusSuccess(rm, successAddLabel)
	if err != nil {
		updateErr := r.updateErrorStatus(rm, err.Error())
		if updateErr != nil {
			return updateErr
		}
		return err
	}

	return nil
}

// executeAddLabel removes label from the resource.
func (r *ResourceModifierReconciler) executeRemoveLabel(resource client.Object,
	rm annotresourcemodifv1.ResourceModifier, label string) error {
	labels := resource.GetLabels()

	if _, exists := labels[label]; !exists {
		return nil
	}
	delete(labels, label)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err := r.Client.Update(ctx, resource)
	if err != nil {
		return err
	}

	err = r.updateStatusSuccess(rm, successAddLabel)
	if err != nil {
		updateErr := r.updateErrorStatus(rm, err.Error())
		if updateErr != nil {
			return updateErr
		}
		return err
	}

	return nil
}

// executeScale
func (r *ResourceModifierReconciler) executeScale(resource client.Object,
	rm annotresourcemodifv1.ResourceModifier, desiredReplicas int) error {

	var deployment *v1.Deployment
	var ok bool
	deployment, ok = resource.(*v1.Deployment)
	if !ok {
		return errs.New(errTypeNotScalable)
	}

	err := r.scaleDeployment(deployment, desiredReplicas)
	if err != nil {
		return err
	}

	err = r.updateStatusSuccess(rm, successDeploymentScale)
	if err != nil {
		updateErr := r.updateErrorStatus(rm, err.Error())
		if updateErr != nil {
			return updateErr
		}
		return err
	}

	return nil
}

// scaleDeployment updates Spec.Replicas in the specified deployment.
func (r *ResourceModifierReconciler) scaleDeployment(d *v1.Deployment, desiredReplicas int) error {
	var replicas *int32
	replicas = new(int32)
	*replicas = int32(desiredReplicas)

	d.Spec.Replicas = replicas

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err := r.Client.Update(ctx, d)
	if err != nil {
		return err
	}

	return nil
}
