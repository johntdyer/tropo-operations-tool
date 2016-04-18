package main

// ApplicationConfig global config
type ApplicationConfig struct {
	API struct {
		Protocol           string
		URL                string
		InsecureSkipVerify bool
	}
	Credentials struct {
		Username string
		Password string
	}
}

// Application structs
type Application struct {
	ID                     string
	Name                   string
	Platform               string
	VoiceEnvironmentID     string
	MessagingEnvironmentID string
	MessagingURL           string
	VoiceURL               string
	Environment            string
	Partition              string
	UserID                 int
	UserData               User
}

// User structs
type User struct {
	ID                     string
	Address                string
	Address2               string
	Username               string
	State                  string
	Email                  string
	FirstName              string
	LastName               string
	Notes                  string
	Status                 string
	JoinDate               string
	PasswordFailedAttempts int
}

// Address structs
type Address struct {
	Type                string
	Prefix              string
	Number              string
	DisplayNumber       string
	ServiceID           string
	City                string
	State               string
	Country             string
	ProviderName        string
	SmsEnabled          bool
	ExcludeFromBilling  bool
	SmsRateLimit        int
	ExchangeID          int
	ApplicationID       int
	RequireVerification bool
	OwnerID             int    `json:"ownerId"`
	Owner               string `json:"owner"`
}

// PapiFeaturesResponse structs
type PapiFeaturesResponse []struct {
	Href        string `json:"href"`
	Feature     string `json:"feature"`
	FeatureName string `json:"featureName"`
	FeatureFlag string `json:"featureFlag"`
}

// Feature struct
type Feature []struct {
	ID                       int
	FeatureFlag, FeatureName string
}

// ApplicationAddresses struct
type ApplicationAddresses []struct {
	Type                string
	Prefix              string
	Number              string
	Channel             string
	Address             string
	Token               string
	DisplayNumber       string
	ServiceID           string
	City                string
	State               string
	Country             string
	ProviderName        string
	SmsEnabled          bool
	ExcludeFromBilling  bool
	SmsRateLimit        int
	ExchangeID          int
	ApplicationID       int
	RequireVerification bool
}

// Applications struct
type Applications []struct {
	Href                      string
	ID                        string
	Name                      string
	Platform                  string
	VoiceEnvironmentID        string
	VoiceURL                  string
	MessagingEnvironmentID    string
	MessagingURL              string
	Environment               string
	EventNotificationEnabled  bool
	ResultNotificationEnabled bool
	UserID                    int
	User                      string
	Partition                 string
}
