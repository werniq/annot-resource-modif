package controller

import (
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"testing"
)

func TestResourceModifierReconciler_executeRemoveAnyFinalizerAnnotation(t *testing.T) {
	type fields struct {
		Client client.Client
		Scheme *runtime.Scheme
	}
	type args struct {
		resource client.Object
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ResourceModifierReconciler{
				Client: tt.fields.Client,
				Scheme: tt.fields.Scheme,
			}
			if err := r.executeRemoveAnyFinalizerAnnotation(tt.args.resource); (err != nil) != tt.wantErr {
				t.Errorf("executeRemoveAnyFinalizerAnnotation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
