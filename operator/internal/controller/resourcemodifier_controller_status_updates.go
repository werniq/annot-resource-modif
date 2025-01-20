package controller

import (
	"context"
	v1 "ericsson.com/resource-modif-annotations/api/v1"
	"time"
)

func (r *ResourceModifierReconciler) initResourceModifierStatus(resource v1.ResourceModifier) {
	resource.Status.Conditions = make(map[string]string)
}

// updateErrorStatus updates resource's Conditions with appropriate message. If an error were returned, returns it.
func (r *ResourceModifierReconciler) updateErrorStatus(resource v1.ResourceModifier, reason string) error {
	resource.Status.ErrorStatus(reason)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err := r.Client.Update(ctx, &resource)
	if err != nil {
		return err
	}

	return nil
}

// updateStatusSuccess updates resource's Conditions by adding new Successful status, and removing any previously added
// error statuses (if applicable).
func (r *ResourceModifierReconciler) updateStatusSuccess(resource v1.ResourceModifier, reason string) error {
	resource.Status.SuccessfulStatus(reason)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err := r.Client.Update(ctx, &resource)
	if err != nil {
		return r.updateErrorStatus(resource, err.Error())
	}

	return nil
}
