package v1

const (
	StatusSuccess = "Success"

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

func (r *ResourceModifierStatus) errorStatus() {

}

func (r *ResourceModifierStatus) successfulStatus() {

}
