/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	errs "errors"
	v1 "k8s.io/api/apps/v1"
	v1beta3 "k8s.io/api/apps/v1beta2"
	v3 "k8s.io/api/batch/v1"
	v2 "k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
	v1beta2 "k8s.io/api/networking/v1beta1"
	rbac "k8s.io/api/rbac/v1"
	v1beta4 "k8s.io/api/rbac/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"strings"

	annotresourcemodifv1 "ericsson.com/resource-modif-annotations/api/v1"
)

const (
	// resourceNotFound is an error message indicating that specified resource was not found
	resourceNotFound = "No matches found for specified resource: "
)

// ResourceModifierReconciler reconciles a ResourceModifier object
type ResourceModifierReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=annot-resource-modif.ericsson.com,resources=resourcemodifiers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=annot-resource-modif.ericsson.com,resources=resourcemodifiers/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=annot-resource-modif.ericsson.com,resources=resourcemodifiers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *ResourceModifierReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	var resourceModifier annotresourcemodifv1.ResourceModifier
	if err := r.Get(ctx, req.NamespacedName, &resourceModifier); err != nil {
		log.Error(err, "unable to fetch resourceModifier")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	resource, err := r.determineResourceType(resourceModifier.Spec.ResourceData)
	if err != nil {
		log.Error(err, "Error determining resource type. Wrong resource type specified")
		return ctrl.Result{}, err
	}

	objectKey, err := r.determineResourceSelector(resourceModifier.Spec.ResourceData)
	if err != nil {
		log.Error(err, "Error determining selector. Make sure that either Name or Labels are specified")
		return ctrl.Result{}, err
	}

	err = r.Client.Get(ctx, objectKey, resource)
	if err != nil {
		log.Error(err, "Error while trying to get an object")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ResourceModifierReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&annotresourcemodifv1.ResourceModifier{}).
		Named("resourcemodifier").
		Complete(r)
}

// determineResourceType analyzes resourceData from the arguments, and returns the object which was specified
// in resourceData.ResourceType field.
// If no matches were found, it returns an empty object, and an resourceNotFound error.
func (r *ResourceModifierReconciler) determineResourceType(resourceData annotresourcemodifv1.TargetResourceData) (client.Object, error) {
	switch strings.ToLower(resourceData.ResourceType) {
	case "pod":
		return &v2.Pod{}, nil
	case "deployment":
		return &v1.Deployment{}, nil
	case "cronjob":
		return &v3.CronJob{}, nil
	case "pv":
		return &v2.PersistentVolume{}, nil
	case "pvc":
		return &v2.PersistentVolumeClaim{}, nil
	case "service":
		return &v2.Service{}, nil
	case "ingress":
		return &v1beta2.Ingress{}, nil
	case "role":
		return &rbac.Role{}, nil
	case "rb":
		return &rbac.RoleBinding{}, nil
	case "clusterrole":
		return &rbac.ClusterRole{}, nil
	case "crb":
		return &rbac.ClusterRoleBinding{}, nil
	}

	return nil, errs.New(resourceNotFound + resourceData.ResourceType)
}

func (r *ResourceModifierReconciler) determineResourceSelector(resourceData annotresourcemodifv1.TargetResourceData) (client.ObjectKey, error) {
	if resourceData.Labels != nil {
		// TODO: implement listing by labels
	}

	objectKey := client.ObjectKey{}

	if resourceData.Name != "" {
		objectKey.Name = resourceData.Name
	}

	if resourceData.Namespace != "" {
		objectKey.Namespace = resourceData.Namespace
	}

	return objectKey, nil
}
