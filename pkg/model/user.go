package model

type Registration struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type Login struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginDetails struct {
	Token string `json:"token,omitempty"`
	Name  string `json:"name,omitempty"`
}

type SucessMessage struct {
	Message string `json:"message,omitempty"`
}

type Error struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
