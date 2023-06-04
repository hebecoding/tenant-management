package apperrors

import (
	"github.com/pkg/errors"
)

var (
	ErrCreatingTenantDocument      = errors.New("error creating tenant document in database")
	ErrRetrievingTenantDocument    = errors.New("error retrieving tenant document from database")
	ErrRetrievingTenantDocuments   = errors.New("error retrieving tenant documents from database")
	ErrNoTenantDocumentsFound      = errors.New("no tenant documents found")
	ErrInsertingTenantDocuments    = errors.New("error inserting tenant document(s) into database")
	ErrUpdatingTenantDocuments     = errors.New("error updating tenant documents in database")
	ErrDeletingTenantDocuments     = errors.New("error deleting tenant documents from database")
	ErrUnmarshallingTenantDocument = errors.New("error unmarshalling tenant document")
)

const (
	ErrRollingBackTransaction = "rolling back transaction"
	ErrCreatingTenant         = "error creating tenant - %v"
	ErrRetrievingTenant       = "error retrieving tenant - %v"
	ErrRetrievingTenants      = "error retrieving tenants"
	ErrUnmarshallingTenant    = "error unmarshalling tenants"
	ErrNoTenantFound          = "no tenant found"
)
