package jsons

// JSON structure to validate User and Token pair
type JsonUserToken struct {
	Username string `json:"username"`
	Token string `json:"token"`
}

// JSON structure to validate User and Passwd pair
type JsonUserPassword struct {
	Username string `json:"username"`
	Password string `json:"password"`
}