package entities

import (
	"time"
)

type Tenant struct {
	ID              string                     `json:"_id" bson:"_id"`
	Company         []*TenantCompanyDetails    `json:"company" bson:"company"`
	PaymentDetails  []*TenantPaymentDetails    `json:"payment_details" bson:"payment_details"`
	Subscription    *TenantSubscriptionDetails `json:"subscription_details" bson:"subscription_details"`
	TenantMetadata  *TenantMetadata            `json:"tenant_metadata" bson:"tenant_metadata"`
	PrimaryContacts []*TenantContactDetails    `json:"primary_contacts" bson:"primary_contacts"`
	Subdomain       string                     `json:"subdomain" bson:"subdomain"`
	CreatedAt       time.Time                  `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time                  `json:"updated_at" bson:"updated_at"`
	IsActive        bool                       `json:"is_active" bson:"is_active"`
	DeletedAt       time.Time                  `json:"deleted_at" bson:"deleted_at"`
	UpdatedBy       string                     `json:"updated_by,omitempty" bson:"updated_by"`
}

type TenantPaymentDetails struct {
	ID             string `json:"_id,omitempty" bson:"_id,omitempty"`
	PaymentMethod  string `json:"payment_method" bson:"payment_method"`
	CardNumber     string `json:"card_number" bson:"card_number"`
	ExpMonth       int    `json:"exp_month" bson:"exp_month"`
	ExpYear        int    `json:"exp_year" bson:"exp_year"`
	SecurityCode   string `json:"security_code" bson:"security_code"`
	BillingAddress string `json:"billing_address" bson:"billing_address"`
	BillingEmail   string `json:"billing_email" bson:"billing_email"`
	IsActive       bool   `json:"is_active" bson:"is_active"`
}

type TenantCompanyDetails struct {
	ID                 string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name               string `json:"name" bson:"company_name"`
	Address            string `json:"address" bson:"address"`
	City               string `json:"city" bson:"city"`
	State              string `json:"state" bson:"state"`
	ZipCode            string `json:"zip_code" bson:"zip_code"`
	Country            string `json:"country" bson:"country"`
	WebsiteURL         string `json:"website_url" bson:"website_url"`
	LogoURL            string `json:"logo_url" bson:"logo_url"`
	Industry           string `json:"industry" bson:"industry"`
	RegistrationNumber string `json:"registration_number" bson:"registration_number"`
	VATNumber          string `json:"vat_number" bson:"vat_number"`
	OperatingHours     string `json:"operating_hours" bson:"operating_hours"`
	IsActive           bool   `json:"is_active" bson:"is_active"`
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
	Roles             []*Role `json:"roles" bson:"roles"`
	IsActive          bool    `json:"is_active" bson:"is_active"`
}

type TenantSubscriptionDetails struct {
	ID              string    `json:"_id" bson:"_id"`
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
	IsActive        bool      `json:"is_active" bson:"is_active"`
}

type TenantMetadata struct {
	ID           string    `json:"_id" bson:"_id"`
	DatabaseName string    `json:"database_name" bson:"database_name"`
	DatabaseType string    `json:"database_type" bson:"database_type"`
	StorageQuota int64     `json:"storage_quota" bson:"storage_quota"`
	StorageUsed  int64     `json:"storage_used" bson:"storage_used"`
	APIKey       string    `json:"api_key" bson:"api_key"`
	TimeZone     string    `json:"time_zone" bson:"time_zone"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at"`
}
