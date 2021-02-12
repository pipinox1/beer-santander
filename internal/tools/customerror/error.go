package customerror

type BusinessError struct {
	Message string
}

type UnauthorizedError struct {
	Message string
}

type ExternalServiceError struct {
	Message         string
	ExternalService string
	Status          int
}

func (e BusinessError) Error() string {
	return e.Message
}
func (e ExternalServiceError) Error() string {
	return e.Message
}

func (e UnauthorizedError) Error() string {
	return e.Message
}

func NewBusinessError(message string) error {
	return BusinessError{
		Message: message,
	}
}

func NewUnauthorizedError(message string) error {
	return UnauthorizedError{
		Message: message,
	}
}

func NewExternalServiceError(message string, serviceName string, status int) error {
	return ExternalServiceError{
		Message:         message,
		Status:          status,
		ExternalService: serviceName,
	}
}
