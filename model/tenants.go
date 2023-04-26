package model

import (
	"time"

	"github.com/hebecoding/digital-dash-commons/utils"
)

type Tenant struct {
	ID             utils.XID                  `json:"id"`
	Company        *TenantCompanyDetails      `json:"company"`
	PaymentDetails *TenantPaymentDetails      `json:"payment_details"`
	Subscription   *TenantSubscriptionDetails `json:"subscription"`
	TenantMetadata *TenantMetadata            `json:"tenant_metadata"`
	CreatedAt      time.Time                  `json:"created_at"`
	UpdatedAt      time.Time                  `json:"updated_at"`
	IsActive       bool                       `json:"is_active"`
	DeletedAt      *time.Time                 `json:"deleted_at,omitempty"`
}

type TenantPaymentDetails struct {
	PaymentMethod string `json:"payment_method"`
	CardNumber    string `json:"card_number"`
	ExpMonth      int    `json:"exp_month"`
	ExpYear       int    `json:"exp_year"`
	SecurityCode  string `json:"security_code"`
}

type TenantCompanyDetails struct {
	CompanyName     string                  `json:"company_name"`
	Address         string                  `json:"address"`
	City            string                  `json:"city"`
	State           string                  `json:"state"`
	ZipCode         string                  `json:"zip_code"`
	Country         string                  `json:"country"`
	PrimaryContacts []*TenantContactDetails `json:"primary_contacts"`
}

type TenantContactDetails struct {
	FirstName   string        `json:"first_name"`
	LastName    string        `json:"last_name"`
	Email       string        `json:"email"`
	PhoneNumber string        `json:"phone_number"`
	TenantRoles []*TenantRBAC `json:"tenant_roles"`
}

type TenantSubscriptionDetails struct {
	Plan      string    `json:"plan"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Active    bool      `json:"active"`
}

type TenantRBAC struct {
	RoleID   string `json:"role_id"`
	RoleName string `json:"role_name"`
}

type TenantMetadata struct {
	ID           string    `json:"id"`
	DatabaseName string    `json:"database_name"`
	DatabaseType string    `json:"database_type"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
