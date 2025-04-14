package mastodon

type Profile struct {
	Username    string  `json:"username"`
	DisplayName string  `json:"display_name"`
	Bio         string  `json:"bio"`
	Avatar      string  `json:"avatar"`
	Header      string  `json:"header"`
	Followers   int     `json:"followers_count"`
	Following   int     `json:"following_count"`
	Posts       int     `json:"statuses_count"`
	Fields      []Field `json:"fields"` // Extra fields
}

type Field struct {
	Name  string `json:"name"`  // The name of the field
	Value string `json:"value"` // The value of the field
}
