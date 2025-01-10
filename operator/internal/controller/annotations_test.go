package controller

import (
	"context"
	v1 "ericsson.com/resource-modif-annotations/api/v1"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	v2 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/client/interceptor"
	"testing"
)

var k8sClient client.Client

func TestResourceModifierReconciler_executeRemoveAnyFinalizerAnnotation(t *testing.T) {
	scheme := runtime.NewScheme()
	assert.Nil(t, v1.AddToScheme(scheme))
	assert.Nil(t, v2.AddToScheme(scheme))

	podWithFinalizers := &v2.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:       "test-podWithFinalizers",
			Namespace:  "test-ns",
			Finalizers: []string{"test-finalizer"},
		},
	}
	podWithoutFinalizers := &v2.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-podWithFinalizers",
			Namespace: "test-ns",
		},
	}
	rm := &v1.ResourceModifier{
		ObjectMeta: metav1.ObjectMeta{
			Name: "rm-test",
		},
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
	rm.Status.Conditions = make(map[string]string)
	k8sClient = fake.NewClientBuilder().
		WithScheme(scheme).
		WithObjects(podWithFinalizers, rm).
		Build()

	updateErrK8sClient := fake.NewClientBuilder().
		WithScheme(scheme).
		WithInterceptorFuncs(interceptor.Funcs{
			Update: func(ctx context.Context, client client.WithWatch, obj client.Object, opts ...client.UpdateOption) error {
				return errors.New("error during update")
			}}).
		WithObjects(podWithFinalizers, rm).
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
			name: "Error while removing finalizer [UPDATE ERROR]",
			fields: fields{
				Client: updateErrK8sClient,
				Scheme: scheme,
			},
			args: args{
				resource: podWithFinalizers,
				rm:       *rm,
			},
			wantErr: true,
		},
		{
			name: "Successful finalizer remove",
			fields: fields{
				Client: k8sClient,
				Scheme: scheme,
			},
			args: args{
				resource: podWithFinalizers,
				rm:       *rm,
			},
			wantErr: false,
		},
		{
			name: "Pod with no finalizers - should fast fail",
			fields: fields{
				Client: k8sClient,
				Scheme: scheme,
			},
			args: args{
				resource: podWithoutFinalizers,
				rm:       *rm,
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

func TestResourceModifierReconciler_executeAddFinalizer(t *testing.T) {
	scheme := runtime.NewScheme()
	assert.Nil(t, v1.AddToScheme(scheme))
	assert.Nil(t, v2.AddToScheme(scheme))

	podWithoutFinalizers := &v2.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:       "test-podWithoutFinalizers",
			Namespace:  "test-ns",
			Finalizers: []string{},
		},
	}

	desiredFinalizer := "finalizer.ericsson.com"

	podWithDesiredFinalizer := &v2.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:       "test-podWithFinalizers",
			Namespace:  "test-ns",
			Finalizers: []string{desiredFinalizer},
		},
	}

	rm := &v1.ResourceModifier{
		ObjectMeta: metav1.ObjectMeta{
			Name: "rm-test",
		},
		Spec: v1.ResourceModifierSpec{
			ResourceData: v1.TargetResourceData{
				Name:         "test-pod",
				Namespace:    "test-ns",
				ResourceType: "pod",
			},
			Annotations: []string{
				"addFinalizer:" + desiredFinalizer,
			},
		},
	}

	rm.Status.Conditions = make(map[string]string)
	k8sClient = fake.NewClientBuilder().
		WithScheme(scheme).
		WithObjects(podWithoutFinalizers, rm).
		Build()

	updateErrK8sClient := fake.NewClientBuilder().
		WithScheme(scheme).
		WithInterceptorFuncs(interceptor.Funcs{
			Update: func(ctx context.Context, client client.WithWatch, obj client.Object, opts ...client.UpdateOption) error {
				return errors.New("error during update")
			}}).
		WithObjects(podWithoutFinalizers, rm).
		Build()

	type fields struct {
		Client client.Client
		Scheme *runtime.Scheme
	}
	type args struct {
		resource  client.Object
		rm        v1.ResourceModifier
		finalizer string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Successful finalizer append",
			fields: fields{
				Client: k8sClient,
				Scheme: scheme,
			},
			args: args{
				resource:  podWithoutFinalizers,
				rm:        *rm,
				finalizer: desiredFinalizer,
			},
			wantErr: false,
		},
		{
			name: "Unsuccessful finalizer append - Finalizer already exists",
			fields: fields{
				Client: k8sClient,
				Scheme: scheme,
			},
			args: args{
				resource:  podWithDesiredFinalizer,
				rm:        *rm,
				finalizer: desiredFinalizer,
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
			err := r.executeAddFinalizer(tt.args.resource, tt.args.rm, tt.args.finalizer)
			// TODO: add additional test for error message
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
