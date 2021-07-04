package model

type domainError struct {
	err string
}

type infraError struct {
	err string
}

//DomainError interface representing the domainError struct
type DomainError interface {
	Error() string
}

//InfraError interface representing the infraError struct
type InfraError interface {
	Error() string
}

//NewDomainError create a new domainError struct
func NewDomainError(message string) DomainError {
	return &domainError{err: message}
}

//NewInfraError create a new infraError struct
func NewInfraError(message string) InfraError {
	return &infraError{err: message}
}

//Error returns the error message
func (domainError *domainError) Error() string {
	return domainError.err
}

//Error returns the error message
func (infraError *infraError) Error() string {
	return infraError.err
}
