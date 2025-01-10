package controller

import (
	v1 "ericsson.com/resource-modif-annotations/api/v1"
	"github.com/stretchr/testify/assert"
	v2 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"testing"
)

func TestResourceModifierReconciler_executeRemoveAnyFinalizerAnnotation(t *testing.T) {
	scheme := runtime.NewScheme()
	assert.Nil(t, v1.AddToScheme(scheme))
	assert.Nil(t, v2.AddToScheme(scheme))

	pod := &v2.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:       "test-pod",
			Namespace:  "test-ns",
			Finalizers: []string{"test-finalizer"},
		},
	}
	rm := v1.ResourceModifier{
		Spec: v1.ResourceModifierSpec{
			ResourceData: v1.TargetResourceData{
				Name:         "test-pod",
				Namespace:    "test-ns",
				ResourceType: "pod",
			},
			Annotations: []string{
				"removeAnyFinalizers",
			},
		},
	}

	k8sClient = fake.NewClientBuilder().
		WithScheme(scheme).
		WithObjects(pod).
		Build()
	type fields struct {
		Client client.Client
		Scheme *runtime.Scheme
	}
	type args struct {
		resource client.Object
		rm       v1.ResourceModifier
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Successful finalizer remove",
			fields: fields{
				Client: nil,
				Scheme: nil,
			},
			args: args{
				resource: pod,
				rm:       rm,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ResourceModifierReconciler{
				Client: tt.fields.Client,
				Scheme: tt.fields.Scheme,
			}
			if err := r.executeRemoveAnyFinalizerAnnotation(tt.args.resource, tt.args.rm); (err != nil) != tt.wantErr {
				t.Errorf("executeRemoveAnyFinalizerAnnotation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
