package v1

import "testing"

func TestResourceModifierStatus_ErrorStatus(t *testing.T) {
	type fields struct {
		Conditions map[string]string
	}
	type args struct {
		reason string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Successfully update condition",
			fields: fields{
				Conditions: make(map[string]string),
			},
			args: args{reason: "error while updating status"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ResourceModifierStatus{
				Conditions: tt.fields.Conditions,
			}
			r.ErrorStatus(tt.args.reason)
		})
	}
}

func TestResourceModifierStatus_SuccessfulStatus(t *testing.T) {
	type fields struct {
		Conditions map[string]string
	}
	type args struct {
		reason string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Updating non-existent conditions",
			fields: fields{
				Conditions: map[string]string{
					StatusError: "error while updating status",
				},
			},
			args: args{reason: "error while updating status 2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ResourceModifierStatus{
				Conditions: tt.fields.Conditions,
			}
			r.SuccessfulStatus(tt.args.reason)
		})
	}
}
