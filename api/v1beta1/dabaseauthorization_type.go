package v1beta1

type DatabaseAuthorization struct {
	Enable   bool   `json:"enable,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
}
