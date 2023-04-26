package tenants

import (
	"time"

	"github.com/hebecoding/digital-dash-commons/utils"
)

type Tenant struct {
	ID              utils.XID                  `json:"id"`
	Company         *TenantCompanyDetails      `json:"company"`
	PaymentDetails  *TenantPaymentDetails      `json:"payment_details"`
	Subscription    *TenantSubscriptionDetails `json:"subscription"`
	TenantMetadata  *TenantMetadata            `json:"tenant_metadata"`
	PrimaryContacts []*TenantContactDetails    `json:"primary_contacts"`
	Subdomain       string                     `json:"subdomain"`
	CreatedAt       time.Time                  `json:"created_at"`
	UpdatedAt       time.Time                  `json:"updated_at"`
	IsActive        bool                       `json:"is_active"`
	DeletedAt       *time.Time                 `json:"deleted_at,omitempty"`
}

type TenantPaymentDetails struct {
	PaymentMethod  string `json:"payment_method"`
	CardNumber     string `json:"card_number"`
	ExpMonth       int    `json:"exp_month"`
	ExpYear        int    `json:"exp_year"`
	SecurityCode   string `json:"security_code"`
	BillingAddress string `json:"billing_address"`
	BillingEmail   string `json:"billing_email"`
}

type TenantCompanyDetails struct {
	CompanyName        string `json:"company_name"`
	Address            string `json:"address"`
	City               string `json:"city"`
	State              string `json:"state"`
	ZipCode            string `json:"zip_code"`
	Country            string `json:"country"`
	WebsiteURL         string `json:"website_url"`
	LogoURL            string `json:"logo_url"`
	Industry           string `json:"industry"`
	RegistrationNumber string `json:"registration_number"`
	VATNumber          string `json:"vat_number"`
	OperatingHours     string `json:"operating_hours"`
}

type TenantContactDetails struct {
	FirstName         string        `json:"first_name"`
	LastName          string        `json:"last_name"`
	Email             string        `json:"email"`
	PhoneNumber       string        `json:"phone_number"`
	AvatarURL         string        `json:"avatar_url"`
	JobTitle          string        `json:"job_title"`
	PreferredLanguage string        `json:"preferred_language"`
	Timezone          string        `json:"timezone"`
	TenantRoles       []*TenantRBAC `json:"tenant_roles"`
}

type TenantSubscriptionDetails struct {
	Plan            string    `json:"plan"`
	StartDate       time.Time `json:"start_date"`
	EndDate         time.Time `json:"end_date"`
	Active          bool      `json:"active"`
	BillingCycle    string    `json:"billing_cycle"`
	TrialEndDate    time.Time `json:"trial_end_date"`
	NextBillingDate time.Time `json:"next_billing_date"`
	PaymentStatus   string    `json:"payment_status"`
	Discount        bool      `json:"discount"`
	DiscountRate    float64   `json:"discount_rate"`
	AutoRenew       bool      `json:"auto_renew"`
	LastPaymentDate time.Time `json:"last_payment_date"`
	PaymentGateway  string    `json:"payment_gateway"`
}

type TenantRBAC struct {
	RoleID   string `json:"role_id"`
	RoleName string `json:"role_name"`
}

type TenantMetadata struct {
	ID           string    `json:"id"`
	DatabaseName string    `json:"database_name"`
	DatabaseType string    `json:"database_type"`
	StorageQuota int64     `json:"storage_quota"`
	StorageUsed  int64     `json:"storage_used"`
	APIKey       string    `json:"api_key"`
	TimeZone     string    `json:"time_zone"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
