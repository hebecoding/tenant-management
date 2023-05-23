package mongo

import (
	"context"

	"github.com/hebecoding/digital-dash-commons/utils"
	"github.com/hebecoding/tenant-management/internal/domain/entities"
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
func (r *TenantRepository) Create(ctx context.Context, tenant *entities.Tenant) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	_, err := r.db.InsertOne(ctx, tenant)
	if err != nil {
		r.logger.Infof("error inserting tenant into database: %v", err)
		r.logger.Info("rolling back transaction")

		return err
	}
	return nil
}
