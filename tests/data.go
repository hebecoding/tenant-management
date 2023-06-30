package tests

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/hebecoding/tenant-management/internal/domain/entities"
)

var generator = gofakeit.NewCrypto()

func CreateTenant() *entities.Tenant {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	var mockTenant = entities.Tenant{}
	mockTenant.ID = generator.UUID()
	mockTenant.Name = generator.Company()
	mockTenant.Subdomain = strings.ToLower(strings.Replace(mockTenant.Name, " ", "", -1))
	mockTenant.IsActive = true

	// generate companies
	for i := 0; i < rand.Intn(3); i++ {
		mockTenant.Companies = append(mockTenant.Companies, GenerateCompany())
	}

	if len(mockTenant.Companies) == 0 {
		mockTenant.Companies = append(mockTenant.Companies, GenerateCompany())
		mockTenant.Companies[0].IsActive = true
	}

	// generate payment details
	for i := 0; i < rand.Intn(3); i++ {
		mockTenant.PaymentDetails = append(mockTenant.PaymentDetails, GeneratePaymentDetails())
	}

	if len(mockTenant.PaymentDetails) == 0 {
		mockTenant.PaymentDetails = append(mockTenant.PaymentDetails, GeneratePaymentDetails())
		mockTenant.PaymentDetails[0].IsActive = true
	}

	// generate contact details
	for i := 0; i < rand.Intn(3); i++ {
		mockTenant.PrimaryContacts = append(mockTenant.PrimaryContacts, GenerateContactDetails())
	}

	if len(mockTenant.PrimaryContacts) == 0 {
		mockTenant.PrimaryContacts = append(mockTenant.PrimaryContacts, GenerateContactDetails())
		mockTenant.PrimaryContacts[0].IsActive = true
	}

	// generate tenant metadata
	mockTenant.TenantMetadata = GenerateTenantMetadata(mockTenant)

	mockTenant.CreatedAt = time.Now().UTC().Truncate(time.Millisecond)

	return &mockTenant
}

func CreateTenantList(amount int) []*entities.Tenant {
	var tenants []*entities.Tenant

	for i := 0; i < amount; i++ {
		tenants = append(tenants, CreateTenant())
	}

	return tenants
}

func GeneratePaymentDetails() *entities.TenantPaymentDetails {
	cc := generator.CreditCard()
	addr := generator.Address()
	addressInfo := entities.Address{
		ID:      generator.UUID(),
		Address: addr.Street,
		City:    addr.City,
		State:   addr.State,
		ZipCode: addr.Zip,
		Country: addr.Country,
	}

	expiration := strings.Split(cc.Exp, "/")
	month, _ := strconv.Atoi(expiration[0])
	year, _ := strconv.Atoi(expiration[1])

	return &entities.TenantPaymentDetails{
		ID:           generator.UUID(),
		Address:      &addressInfo,
		CardType:     cc.Type,
		CardNumber:   cc.Number,
		SecurityCode: cc.Cvv,
		ExpMonth:     month,
		ExpYear:      year,
		IsActive:     false,
	}
}

func GenerateSubscriptionDetails() *entities.TenantSubscriptionDetails {

	plans := []string{"starter", "basic", "premium", "enterprise"}
	startDate := gofakeit.DateRange(time.Now(), time.Now()).UTC().Truncate(time.Millisecond)
	endDate := startDate.Add(time.Hour * 24 * 30).UTC().UTC().Truncate(time.Millisecond)
	billingDate := startDate.Add(time.Hour * 24 * 29).UTC().Truncate(time.Millisecond)

	return &entities.TenantSubscriptionDetails{
		ID:              generator.UUID(),
		Plan:            gofakeit.RandomString(plans),
		BillingCycle:    gofakeit.RandomString([]string{"weekly", "monthly", "yearly"}),
		PaymentStatus:   gofakeit.RandomString([]string{"active", "inactive", "suspended"}),
		PaymentGateway:  gofakeit.RandomString([]string{"stripe", "paypal", "braintree"}),
		DiscountRate:    float64(gofakeit.RandomInt([]int{0, 5, 10, 15, 20, 25})),
		Discount:        gofakeit.Bool(),
		Active:          gofakeit.Bool(),
		AutoRenew:       gofakeit.Bool(),
		StartDate:       startDate,
		EndDate:         endDate,
		NextBillingDate: billingDate,
		LastPaymentDate: billingDate,
	}
}

func GenerateCompany() *entities.TenantCompanyDetails {

	addr := generator.Address()
	addressInfo := entities.Address{
		ID:      generator.UUID(),
		Address: addr.Street,
		City:    addr.City,
		State:   addr.State,
		ZipCode: addr.Zip,
		Country: addr.Country,
	}

	industry := generator.RandomString(
		[]string{
			"IT", "Finance", "Healthcare", "Education", "Retail", "Manufacturing", "Transportation", "Hospitality",
			"Real Estate", "Construction", "Agriculture", "Mining", "Utilities", "Telecommunications", "Media",
			"Entertainment", "Government", "Non-profit",
		},
	)

	subscriptions := []*entities.TenantSubscriptionDetails{}

	for i := 0; i <= rand.Intn(2); i++ {
		subscriptions = append(subscriptions, GenerateSubscriptionDetails())
	}

	return &entities.TenantCompanyDetails{
		ID:                 generator.UUID(),
		Name:               generator.Company(),
		WebsiteURL:         generator.URL(),
		LogoURL:            generator.ImageURL(4, 6),
		Industry:           industry,
		RegistrationNumber: strconv.FormatInt(rand.Int63(), 10),
		IsActive:           generator.Bool(),
		Subscriptions:      subscriptions,
		Address:            &addressInfo,
	}
}

func GenerateContactDetails() *entities.TenantContactDetails {
	languages := []string{"en", "fr", "es", "de", "it", "pt", "ru", "zh", "ja", "ko"}

	roles := []*entities.Role{
		{
			ID:          generator.UUID(),
			Name:        "Admin",
			Description: "The admin role has full access to the tenant",
			Permissions: []entities.Permission{
				entities.ReadPermission, entities.WritePermission, entities.DeletePermission, entities.EditPermission,
			},
		}, {
			ID:          generator.UUID(),
			Name:        "Manager",
			Description: "The admin role has full access to the tenant",
			Permissions: []entities.Permission{
				entities.ReadPermission, entities.WritePermission, entities.EditPermission,
			},
		}, {
			ID:          generator.UUID(),
			Name:        "Staff",
			Description: "The admin role has full access to the tenant",
			Permissions: []entities.Permission{entities.ReadPermission},
		}, {
			ID:          generator.UUID(),
			Name:        "Back Office",
			Description: "The admin role has full access to the tenant",
			Permissions: []entities.Permission{
				entities.ReadPermission, entities.WritePermission, entities.DeletePermission, entities.EditPermission,
			},
		},
	}

	// generate random roles from the list above
	randomRoles := []*entities.Role{}
	for i := 0; i < rand.Intn(2); i++ {
		randomRoles = append(randomRoles, roles[rand.Intn(len(roles))])
	}

	return &entities.TenantContactDetails{
		ID:                generator.UUID(),
		FirstName:         generator.FirstName(),
		LastName:          generator.LastName(),
		Email:             generator.Email(),
		PhoneNumber:       generator.PhoneFormatted(),
		AvatarURL:         generator.ImageURL(4, 6),
		JobTitle:          generator.JobTitle(),
		PreferredLanguage: generator.RandomString(languages),
		Timezone:          generator.TimeZone(),
		IsActive:          generator.Bool(),
		Roles:             randomRoles,
	}
}

func GenerateTenantMetadata(mockTenant entities.Tenant) *entities.TenantMetadata {
	startDate := gofakeit.DateRange(time.Now(), time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC))

	updated := startDate.Add(time.Hour * 24 * 30)

	return &entities.TenantMetadata{
		ID:           generator.UUID(),
		DatabaseName: fmt.Sprintf("%s-%s", mockTenant.Name, generator.UUID()),
		TimeZone:     generator.TimeZone(),
		StorageQuota: int64(generator.IntRange(100, 1000000000)),
		StorageUsed:  int64(generator.IntRange(0, 100000)),
		CreatedAt:    startDate.Truncate(time.Millisecond),
		UpdatedAt:    updated.Truncate(time.Millisecond),
	}
}
