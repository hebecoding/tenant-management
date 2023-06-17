package apperrors

import (
	"github.com/pkg/errors"
)

var (
	ErrCreatingTenantDocument      = errors.New("error creating tenant document in database")
	ErrRetrievingTenantDocument    = errors.New("error retrieving tenant document(s) from database")
	ErrNoTenantDocumentsFound      = errors.New("no tenant documents found")
	ErrUpdatingTenantDocument      = errors.New("error updating tenant document(s) in database")
	ErrDeletingTenantDocument      = errors.New("error deleting tenant document(s) from database")
	ErrUnmarshallingTenantDocument = errors.New("error unmarshalling tenant document")
	ErrInvalidTenantSubscription   = errors.New("invalid tenant subscription")
)

const (
	ErrRollingBackTransaction = "rolling back transaction"
	ErrCreatingTenant         = "error creating tenant - %v"
	ErrDeletingTenant         = "error deleting tenant - %v"
	ErrRetrievingTenant       = "error retrieving tenant - %v"
	ErrRetrievingTenants      = "error retrieving tenants"
	ErrUnmarshallingTenant    = "error unmarshalling tenants"
	ErrNoTenantFound          = "no tenant found - %v"
	ErrUpdatingTenant         = "error updating tenant - %v"
)
