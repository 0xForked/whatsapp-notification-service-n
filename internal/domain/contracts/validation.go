package contracts

// ValidationErrors is a contract for validation errors.
type ValidationErrors interface {
	// All returns the errors.
	All(model interface{}, err error) map[string]string
}
