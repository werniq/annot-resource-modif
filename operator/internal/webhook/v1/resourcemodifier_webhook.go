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

package v1

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	annotresourcemodifv1 "ericsson.com/resource-modif-annotations/api/v1"
)

// nolint:unused
// log is for logging in this package.
var resourcemodifierlog = logf.Log.WithName("resourcemodifier-resource")

// SetupResourceModifierWebhookWithManager registers the webhook for ResourceModifier in the manager.
func SetupResourceModifierWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&annotresourcemodifv1.ResourceModifier{}).
		WithValidator(&ResourceModifierCustomValidator{}).
		WithDefaulter(&ResourceModifierCustomDefaulter{}).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-annot-resource-modif-ericsson-com-v1-resourcemodifier,mutating=true,failurePolicy=fail,sideEffects=None,groups=annot-resource-modif.ericsson.com,resources=resourcemodifiers,verbs=create;update,versions=v1,name=mresourcemodifier-v1.kb.io,admissionReviewVersions=v1

// ResourceModifierCustomDefaulter struct is responsible for setting default values on the custom resource of the
// Kind ResourceModifier when those are created or updated.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as it is used only for temporary operations and does not need to be deeply copied.
type ResourceModifierCustomDefaulter struct {
	// TODO(user): Add more fields as needed for defaulting
}

var _ webhook.CustomDefaulter = &ResourceModifierCustomDefaulter{}

// Default implements webhook.CustomDefaulter so a webhook will be registered for the Kind ResourceModifier.
func (d *ResourceModifierCustomDefaulter) Default(ctx context.Context, obj runtime.Object) error {
	resourcemodifier, ok := obj.(*annotresourcemodifv1.ResourceModifier)

	if !ok {
		return fmt.Errorf("expected an ResourceModifier object but got %T", obj)
	}
	resourcemodifierlog.Info("Defaulting for ResourceModifier", "name", resourcemodifier.GetName())

	// TODO(user): fill in your defaulting logic.

	return nil
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// NOTE: The 'path' attribute must follow a specific pattern and should not be modified directly here.
// Modifying the path for an invalid path can cause API server errors; failing to locate the webhook.
// +kubebuilder:webhook:path=/validate-annot-resource-modif-ericsson-com-v1-resourcemodifier,mutating=false,failurePolicy=fail,sideEffects=None,groups=annot-resource-modif.ericsson.com,resources=resourcemodifiers,verbs=create;update,versions=v1,name=vresourcemodifier-v1.kb.io,admissionReviewVersions=v1

// ResourceModifierCustomValidator struct is responsible for validating the ResourceModifier resource
// when it is created, updated, or deleted.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as this struct is used only for temporary operations and does not need to be deeply copied.
type ResourceModifierCustomValidator struct {
	//TODO(user): Add more fields as needed for validation
}

var _ webhook.CustomValidator = &ResourceModifierCustomValidator{}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type ResourceModifier.
func (v *ResourceModifierCustomValidator) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	resourcemodifier, ok := obj.(*annotresourcemodifv1.ResourceModifier)
	if !ok {
		return nil, fmt.Errorf("expected a ResourceModifier object but got %T", obj)
	}
	resourcemodifierlog.Info("Validation for ResourceModifier upon creation", "name", resourcemodifier.GetName())

	// TODO(user): fill in your validation logic upon object creation.

	return nil, nil
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type ResourceModifier.
func (v *ResourceModifierCustomValidator) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	resourcemodifier, ok := newObj.(*annotresourcemodifv1.ResourceModifier)
	if !ok {
		return nil, fmt.Errorf("expected a ResourceModifier object for the newObj but got %T", newObj)
	}
	resourcemodifierlog.Info("Validation for ResourceModifier upon update", "name", resourcemodifier.GetName())

	// TODO(user): fill in your validation logic upon object update.

	return nil, nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type ResourceModifier.
func (v *ResourceModifierCustomValidator) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	resourcemodifier, ok := obj.(*annotresourcemodifv1.ResourceModifier)
	if !ok {
		return nil, fmt.Errorf("expected a ResourceModifier object but got %T", obj)
	}
	resourcemodifierlog.Info("Validation for ResourceModifier upon deletion", "name", resourcemodifier.GetName())

	// TODO(user): fill in your validation logic upon object deletion.

	return nil, nil
}
