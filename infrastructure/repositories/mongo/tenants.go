package mongo

import (
	"context"
	"time"

	"github.com/hebecoding/digital-dash-commons/utils"
	"github.com/hebecoding/tenant-management/infrastructure/apperrors"
	"github.com/hebecoding/tenant-management/internal/domain/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TenantRepository struct {
	db     *mongo.Collection
	logger *utils.Logger
}

func NewTenantRepository(db *mongo.Collection, logger *utils.Logger) *TenantRepository {
	return &TenantRepository{
		db:     db,
		logger: logger,
	}
}

// Create creates a new tenant in the database.
// Ctx is used to cancel the operation if the context is cancelled.
// Tenants is the tenant to be created.
func (r *TenantRepository) Create(ctx context.Context, tenant *entities.Tenant) error {
	r.logger.Infof("inserting tenant into database: %v", tenant.ID)
	_, err := r.db.InsertOne(ctx, tenant)
	if err != nil {
		r.logger.Errorf(apperrors.ErrCreatingTenant, tenant.ID)
		r.logger.Error(err)
		r.logger.Error(apperrors.ErrRollingBackTransaction)

		return apperrors.ErrCreatingTenantDocument
	}
	return nil
}

// DeleteTenant deletes a tenant from the database.
// Ctx is used to cancel the operation if the context is cancelled.
// ID is the id of the tenant to be deleted.
// This is a soft delete, isActive is set to false.
func (r *TenantRepository) DeleteTenant(ctx context.Context, id string) error {
	var tenant *entities.Tenant

	r.logger.Infof("deleting tenant from database: %v", id)
	tenant, err := r.GetTenantByID(ctx, id)
	if err != nil {
		return err
	}

	// set isActive to false
	tenant.IsActive = false
	tenant.DeletedAt = time.Now()

	// update tenant in database
	if err := r.UpdateTenant(ctx, tenant); err != nil {
		return err
	}

	return nil
}

// GetTenantByID returns a tenant from the database.
// Ctx is used to cancel the operation if the context is cancelled.
// ID is the id of the tenant to be retrieved.
func (r *TenantRepository) GetTenantByID(ctx context.Context, id string) (*entities.Tenant, error) {
	var tenant *entities.Tenant

	// get tenant from database
	r.logger.Infof("retrieving tenant from database: %v", id)
	if err := r.db.FindOne(
		ctx, bson.M{"_id.id": id},
	).
		Decode(&tenant); err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			r.logger.Errorf(apperrors.ErrNoTenantFound, id)
			return nil, apperrors.ErrNoTenantDocumentsFound
		default:
			r.logger.Errorf(apperrors.ErrRetrievingTenant, id)
			r.logger.Error(err)
			return nil, apperrors.ErrRetrievingTenantDocument
		}
	}

	return tenant, nil
}

// GetTenants returns a list of tenants from the database.
// Ctx is used to cancel the operation if the context is cancelled.
func (r *TenantRepository) GetTenants(ctx context.Context) ([]*entities.Tenant, error) {
	var tenants []*entities.Tenant

	// get all tenants from database
	r.logger.Infoln("retrieving tenants from database")
	cursor, err := r.db.Find(ctx, bson.D{{"is_active", false}})
	if err != nil {
		r.logger.Error(apperrors.ErrRetrievingTenants)
		r.logger.Error(err)
		return nil, apperrors.ErrRetrievingTenantDocument
	}

	// unmarshal all tenants into a slice
	if err := cursor.All(ctx, &tenants); err != nil {
		r.logger.Error(apperrors.ErrUnmarshallingTenant)
		r.logger.Error(err)
		return nil, apperrors.ErrUnmarshallingTenantDocument
	}

	r.logger.Infof("found %d tenants", len(tenants))

	return tenants, nil
}

// UpdateTenant updates a tenant in the database.
// Ctx is used to cancel the operation if the context is cancelled.
// Tenant is the tenant to be updated.
// Only included fields will be updated.
func (r *TenantRepository) UpdateTenant(ctx context.Context, tenant *entities.Tenant) error {
	r.logger.Infof("updating tenant in database: %v", tenant.ID)
	result, err := r.db.UpdateOne(
		ctx,
		bson.M{"_id": tenant.ID},
		bson.M{"$set": tenant},
	)
	if err != nil {
		r.logger.Errorf(apperrors.ErrUpdatingTenant, tenant.ID)
		r.logger.Error(err)
		return apperrors.ErrUpdatingTenantDocument
	}

	r.logger.Infof("updated %v documents", result.ModifiedCount)

	return nil
}
