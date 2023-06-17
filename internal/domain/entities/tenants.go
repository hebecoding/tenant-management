package entities

import (
	"time"
)

type Tenant struct {
	ID              string                  `json:"_id" bson:"_id"`
	Subdomain       string                  `json:"subdomain" bson:"subdomain"`
	UpdatedBy       string                  `json:"updated_by,omitempty" bson:"updated_by"`
	Companies       []*TenantCompanyDetails `json:"companies" bson:"companies"`
	PaymentDetails  []*TenantPaymentDetails `json:"payment_details" bson:"payment_details"`
	TenantMetadata  *TenantMetadata         `json:"tenant_metadata" bson:"tenant_metadata"`
	PrimaryContacts []*TenantContactDetails `json:"primary_contacts" bson:"primary_contacts"`
	CreatedAt       *time.Time              `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt       *time.Time              `json:"updated_at,omitempty" bson:"updated_at"`
	DeletedAt       *time.Time              `json:"deleted_at,omitempty" bson:"deleted_at"`
	IsActive        bool                    `json:"is_active" bson:"is_active"`
}

type TenantPaymentDetails struct {
	ID             string `json:"_id,omitempty" bson:"_id,omitempty"`
	BillingAddress string `json:"billing_address" bson:"billing_address"`
	BillingEmail   string `json:"billing_email" bson:"billing_email"`
	PaymentMethod  string `json:"payment_method" bson:"payment_method"`
	CardNumber     string `json:"card_number" bson:"card_number"`
	SecurityCode   string `json:"security_code" bson:"security_code"`
	ExpMonth       int    `json:"exp_month" bson:"exp_month"`
	ExpYear        int    `json:"exp_year" bson:"exp_year"`
	IsActive       bool   `json:"is_active" bson:"is_active"`
}

type TenantCompanyDetails struct {
	ID                 string                       `json:"_id,omitempty" bson:"_id,omitempty"`
	Name               string                       `json:"name" bson:"company_name"`
	Address            string                       `json:"address" bson:"address"`
	City               string                       `json:"city" bson:"city"`
	State              string                       `json:"state" bson:"state"`
	ZipCode            string                       `json:"zip_code" bson:"zip_code"`
	Country            string                       `json:"country" bson:"country"`
	WebsiteURL         string                       `json:"website_url" bson:"website_url"`
	LogoURL            string                       `json:"logo_url" bson:"logo_url"`
	Industry           string                       `json:"industry" bson:"industry"`
	RegistrationNumber string                       `json:"registration_number" bson:"registration_number"`
	VATNumber          string                       `json:"vat_number" bson:"vat_number"`
	OperatingHours     string                       `json:"operating_hours" bson:"operating_hours"`
	IsActive           bool                         `json:"is_active" bson:"is_active"`
	Subscriptions      []*TenantSubscriptionDetails `json:"subscriptions,omitempty" bson:"subscriptions"`
}

type TenantContactDetails struct {
	ID                string  `json:"_id" bson:"_id"`
	FirstName         string  `json:"first_name" bson:"first_name"`
	LastName          string  `json:"last_name" bson:"last_name"`
	Email             string  `json:"email" bson:"email"`
	UserID            string  `json:"user_id" bson:"user_id"`
	PhoneNumber       string  `json:"phone_number" bson:"phone_number"`
	AvatarURL         string  `json:"avatar_url" bson:"avatar_url"`
	JobTitle          string  `json:"job_title" bson:"job_title"`
	PreferredLanguage string  `json:"preferred_language" bson:"preferred_language"`
	Timezone          string  `json:"timezone" bson:"timezone"`
	IsActive          bool    `json:"is_active" bson:"is_active"`
	Roles             []*Role `json:"roles" bson:"roles"`
}

type TenantSubscriptionDetails struct {
	ID              string     `json:"_id" bson:"_id"`
	Plan            string     `json:"plan" bson:"plan"`
	BillingCycle    string     `json:"billing_cycle" bson:"billing_cycle"`
	PaymentStatus   string     `json:"payment_status" bson:"payment_status"`
	PaymentGateway  string     `json:"payment_gateway" bson:"payment_gateway"`
	StartDate       *time.Time `json:"start_date,omitempty" bson:"start_date"`
	EndDate         *time.Time `json:"end_date,omitempty" bson:"end_date"`
	TrialEndDate    *time.Time `json:"trial_end_date,omitempty" bson:"trial_end_date"`
	NextBillingDate *time.Time `json:"next_billing_date,omitempty" bson:"next_billing_date"`
	LastPaymentDate *time.Time `json:"last_payment_date,omitempty" bson:"last_payment_date"`
	DiscountRate    float64    `json:"discount_rate" bson:"discount_rate"`
	Discount        bool       `json:"discount" bson:"discount"`
	Active          bool       `json:"active" bson:"active"`
	AutoRenew       bool       `json:"auto_renew" bson:"auto_renew"`
}

type TenantMetadata struct {
	ID           string     `json:"_id" bson:"_id"`
	DatabaseName string     `json:"database_name" bson:"database_name"`
	DatabaseType string     `json:"database_type" bson:"database_type"`
	APIKey       string     `json:"api_key" bson:"api_key"`
	TimeZone     string     `json:"time_zone" bson:"time_zone"`
	StorageQuota int64      `json:"storage_quota" bson:"storage_quota"`
	StorageUsed  int64      `json:"storage_used" bson:"storage_used"`
	CreatedAt    *time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}
