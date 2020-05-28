package model

import (
	"time"
)

// Client model
type Client struct {
	ID           string    `json:"id"`
	PrivateKey   string    `json:"private_key"`
	PublicKey    string    `json:"public_key"`
	PresharedKey string	   `json:"preshared_key"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	AllocatedIPs []string  `json:"allocated_ips"`
	AllowedIPs   []string  `json:"allowed_ips"`
	Enabled      bool      `json:"enabled"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// ClientData includes the Client and extra data
type ClientData struct {
	Client *Client
	QRCode string
}
