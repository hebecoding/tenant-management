package entities

import (
	"time"

	"github.com/hebecoding/digital-dash-commons/utils"
)

type Tenant struct {
	ID              utils.XID                  `bson:"_id,omitempty" json:"_id,omitempty"`
	Company         *TenantCompanyDetails      `bson:"company" json:"company"`
	PaymentDetails  *TenantPaymentDetails      `bson:"payment_details" json:"payment_details"`
	Subscription    *TenantSubscriptionDetails `bson:"subscription" json:"subscription"`
	TenantMetadata  *TenantMetadata            `bson:"tenant_metadata" json:"tenant_metadata"`
	PrimaryContacts []*TenantContactDetails    `bson:"primary_contacts" json:"primary_contacts"`
	Subdomain       string                     `bson:"subdomain" json:"subdomain"`
	CreatedAt       time.Time                  `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time                  `bson:"updated_at" json:"updated_at"`
	IsActive        bool                       `bson:"is_active" json:"is_active"`
	DeletedAt       *time.Time                 `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}

type TenantPaymentDetails struct {
	ID             utils.XID `bson:"_id,omitempty" json:"_id,omitempty"`
	PaymentMethod  string    `bson:"payment_method" json:"payment_method"`
	CardNumber     string    `bson:"card_number" json:"card_number"`
	ExpMonth       int       `bson:"exp_month" json:"exp_month"`
	ExpYear        int       `bson:"exp_year" json:"exp_year"`
	SecurityCode   string    `bson:"security_code" json:"security_code"`
	BillingAddress string    `bson:"billing_address" json:"billing_address"`
	BillingEmail   string    `bson:"billing_email" json:"billing_email"`
}

type TenantCompanyDetails struct {
	ID                 utils.XID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name               string    `bson:"company_name" json:"name"`
	Address            string    `bson:"address" json:"address"`
	City               string    `bson:"city" json:"city"`
	State              string    `bson:"state" json:"state"`
	ZipCode            string    `bson:"zip_code" json:"zip_code"`
	Country            string    `bson:"country" json:"country"`
	WebsiteURL         string    `bson:"website_url" json:"website_url"`
	LogoURL            string    `bson:"logo_url" json:"logo_url"`
	Industry           string    `bson:"industry" json:"industry"`
	RegistrationNumber string    `bson:"registration_number" json:"registration_number"`
	VATNumber          string    `bson:"vat_number" json:"vat_number"`
	OperatingHours     string    `bson:"operating_hours" json:"operating_hours"`
}

type TenantContactDetails struct {
	ID                utils.XID     `json:"_id" bson:"_id"`
	FirstName         string        `json:"first_name" bson:"first_name"`
	LastName          string        `json:"last_name" bson:"last_name"`
	Email             string        `json:"email" bson:"email"`
	UserID            utils.XID     `json:"user_id" bson:"user_id"`
	PhoneNumber       string        `json:"phone_number" bson:"phone_number"`
	AvatarURL         string        `json:"avatar_url" bson:"avatar_url"`
	JobTitle          string        `json:"job_title" bson:"job_title"`
	PreferredLanguage string        `json:"preferred_language" bson:"preferred_language"`
	Timezone          string        `json:"timezone" bson:"timezone"`
	Roles             []*TenantRBAC `json:"roles" bson:"roles"`
	Permissions       []Permissions `json:"permissions" bson:"permissions"`
}

type TenantSubscriptionDetails struct {
	ID              utils.XID `json:"_id" bson:"_id"`
	Plan            string    `json:"plan" bson:"plan"`
	StartDate       time.Time `json:"start_date" bson:"start_date"`
	EndDate         time.Time `json:"end_date" bson:"end_date"`
	Active          bool      `json:"active" bson:"active"`
	BillingCycle    string    `json:"billing_cycle" bson:"billing_cycle"`
	TrialEndDate    time.Time `json:"trial_end_date" bson:"trial_end_date"`
	NextBillingDate time.Time `json:"next_billing_date" bson:"next_billing_date"`
	PaymentStatus   string    `json:"payment_status" bson:"payment_status"`
	Discount        bool      `json:"discount" bson:"discount"`
	DiscountRate    float64   `json:"discount_rate" bson:"discount_rate"`
	AutoRenew       bool      `json:"auto_renew" bson:"auto_renew"`
	LastPaymentDate time.Time `json:"last_payment_date" bson:"last_payment_date"`
	PaymentGateway  string    `json:"payment_gateway" bson:"payment_gateway"`
}

type TenantRBAC struct {
	ID       utils.XID `json:"_id" bson:"_id"`
	RoleName string    `json:"role_name" bson:"role_name"`
}

type TenantMetadata struct {
	ID           utils.XID `json:"_id" bson:"_id"`
	DatabaseName string    `json:"database_name" bson:"database_name"`
	DatabaseType string    `json:"database_type" bson:"database_type"`
	StorageQuota int64     `json:"storage_quota" bson:"storage_quota"`
	StorageUsed  int64     `json:"storage_used" bson:"storage_used"`
	APIKey       string    `json:"api_key" bson:"api_key"`
	TimeZone     string    `json:"time_zone" bson:"time_zone"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at"`
}
