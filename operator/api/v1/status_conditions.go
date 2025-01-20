package v1

const (
	// StatusSuccess is a key to Conditions map, which indicates that there were no errors during Reconciliation
	StatusSuccess = "Success"

	// StatusError is a key to Conditions map, which indicates that there were errors during Reconciliation
	StatusError = "Error"
)

// ResourceModifierStatus defines the observed state of ResourceModifier.
type ResourceModifierStatus struct {
	// Conditions are used to describe current state of ResourceModifier.
	// In case of errors, this field is updated, indicating that error had occurred.
	// If Reconciliation was successful - this fields will also be updated, with
	// successful condition type and appropriate message.
	Conditions map[string]string `json:"conditions"`
}

// ErrorStatus initializes/updates the Conditions field with key StatusError and reason as value
func (r *ResourceModifierStatus) ErrorStatus(reason string) {
	r.Conditions[StatusError] = reason
}

// SuccessfulStatus initializes/updates the Conditions field with key StatusSuccess and reason as value
func (r *ResourceModifierStatus) SuccessfulStatus(reason string) {
	if _, exists := r.Conditions[StatusError]; exists {
		delete(r.Conditions, StatusError)
	}
	r.Conditions[StatusSuccess] = reason
}
