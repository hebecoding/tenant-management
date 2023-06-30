package entities

import (
	"time"
)

type Tenant struct {
	ID              string                  `json:"_id,omitempty" bson:"_id"`
	Name            string                  `json:"name,omitempty" bson:"name"`
	Subdomain       string                  `json:"subdomain,omitempty" bson:"subdomain"`
	UpdatedBy       string                  `json:"updated_by,omitempty" bson:"updated_by"`
	IsActive        bool                    `json:"is_active,omitempty" bson:"is_active"`
	Companies       []*TenantCompanyDetails `json:"companies,omitempty" bson:"companies"`
	PaymentDetails  []*TenantPaymentDetails `json:"payment_details,omitempty" bson:"payment_details"`
	TenantMetadata  *TenantMetadata         `json:"tenant_metadata,omitempty" bson:"tenant_metadata"`
	PrimaryContacts []*TenantContactDetails `json:"primary_contacts,omitempty" bson:"primary_contacts"`
	CreatedAt       time.Time               `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt       time.Time               `json:"updated_at,omitempty" bson:"updated_at"`
	DeletedAt       time.Time               `json:"deleted_at,omitempty" bson:"deleted_at"`
}

type TenantPaymentDetails struct {
	ID           string   `json:"_id,omitempty" bson:"_id,omitempty"`
	Address      *Address `json:"billing_address,omitempty" bson:"billing_address"`
	CardType     string   `json:"card_type,omitempty" bson:"card_type"`
	CardNumber   string   `json:"card_number,omitempty" bson:"card_number"`
	SecurityCode string   `json:"security_code,omitempty" bson:"security_code"`
	ExpMonth     int      `json:"exp_month,omitempty" bson:"exp_month"`
	ExpYear      int      `json:"exp_year,omitempty" bson:"exp_year"`
	IsActive     bool     `json:"is_active,omitempty" bson:"is_active"`
}

type TenantCompanyDetails struct {
	ID                 string                       `json:"_id,omitempty" bson:"_id,omitempty"`
	Name               string                       `json:"name,omitempty" bson:"company_name"`
	WebsiteURL         string                       `json:"website_url,omitempty" bson:"website_url"`
	LogoURL            string                       `json:"logo_url,omitempty" bson:"logo_url"`
	Industry           string                       `json:"industry,omitempty" bson:"industry"`
	RegistrationNumber string                       `json:"registration_number,omitempty" bson:"registration_number"`
	IsActive           bool                         `json:"is_active,omitempty" bson:"is_active"`
	Subscriptions      []*TenantSubscriptionDetails `json:"subscriptions,omitempty" bson:"subscriptions"`
	Address            *Address                     `json:"address,omitempty" bson:"address"`
}

type TenantContactDetails struct {
	ID                string  `json:"_id,omitempty" bson:"_id"`
	FirstName         string  `json:"first_name,omitempty" bson:"first_name"`
	LastName          string  `json:"last_name,omitempty" bson:"last_name"`
	Email             string  `json:"email,omitempty" bson:"email"`
	PhoneNumber       string  `json:"phone_number,omitempty" bson:"phone_number"`
	AvatarURL         string  `json:"avatar_url,omitempty" bson:"avatar_url"`
	JobTitle          string  `json:"job_title,omitempty" bson:"job_title"`
	PreferredLanguage string  `json:"preferred_language,omitempty" bson:"preferred_language"`
	Timezone          string  `json:"timezone,omitempty" bson:"timezone"`
	IsActive          bool    `json:"is_active,omitempty" bson:"is_active"`
	Roles             []*Role `json:"roles,omitempty" bson:"roles"`
}

type TenantSubscriptionDetails struct {
	ID              string    `json:"_id,omitempty" bson:"_id"`
	Plan            string    `json:"plan,omitempty" bson:"plan"`
	BillingCycle    string    `json:"billing_cycle,omitempty" bson:"billing_cycle"`
	PaymentStatus   string    `json:"payment_status,omitempty" bson:"payment_status"`
	PaymentGateway  string    `json:"payment_gateway,omitempty" bson:"payment_gateway"`
	DiscountRate    float64   `json:"discount_rate,omitempty" bson:"discount_rate"`
	Discount        bool      `json:"discount,omitempty" bson:"discount"`
	Active          bool      `json:"active,omitempty" bson:"active"`
	AutoRenew       bool      `json:"auto_renew,omitempty" bson:"auto_renew"`
	StartDate       time.Time `json:"start_date,omitempty" bson:"start_date"`
	EndDate         time.Time `json:"end_date,omitempty" bson:"end_date"`
	NextBillingDate time.Time `json:"next_billing_date,omitempty" bson:"next_billing_date"`
	LastPaymentDate time.Time `json:"last_payment_date,omitempty" bson:"last_payment_date"`
}

type TenantMetadata struct {
	ID           string    `json:"_id,omitempty" bson:"_id"`
	DatabaseName string    `json:"database_name,omitempty" bson:"database_name"`
	TimeZone     string    `json:"time_zone,omitempty" bson:"time_zone"`
	StorageQuota int64     `json:"storage_quota,omitempty" bson:"storage_quota"`
	StorageUsed  int64     `json:"storage_used,omitempty" bson:"storage_used"`
	CreatedAt    time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}
