package model

type domainError struct {
	err string
}

type infraError struct {
	err string
}

type DomainError interface {
	Error() string
}

type InfraError interface {
	Error() string
}

func NewDomainError(message string) *domainError {
	return &domainError{err: message}
}

func NewInfraError(message string) *infraError {
	return &infraError{err: message}
}

func (domainError *domainError) Error() string {
	return domainError.err
}

func (infraError *infraError) Error() string {
	return infraError.err
}
