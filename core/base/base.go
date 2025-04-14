package base

type DecodedUser struct {
	Name               string `json:"name"  mapstructure:"name"`  //  i need to do anonymization here for this name
	ID                 string `json:"id"  mapstructure:"id"`      // i need to do anonymization here for this id
	OrgID              string `json:"orgId" mapstructure:"orgId"` // i need to do anonymization here for this id
	Role               string `json:"role"  mapstructure:"role"`
	SubscriptionStatus string `json:"subscription_status" mapstructure:"subscription_status"`
}

type TokenAbilities struct {
	Role string      `json:"role"  mapstructure:"role"`
	User DecodedUser `json:"user" mapstructure:"user"`
}

type PaginatedResponse[T any] struct {
	Data        *[]T  `json:"data"`
	Total       int64 `json:"total"`
	TotalPage   int64 `json:"total_page"`
	CurrentPage int64 `json:"current_page"`
	PerPage     int   `json:"per_page"`
}
