package models

import "shortlink/api"

// ShortcutStore is an interface for interacting with shortcuts
type ShortcutStore interface {
	// Create creates a new shortcut to a targetUrl
	Create(targetURL string) (*api.ShortcutDetail, error)
	// Delete deletes a shortcut described by the urlToken
	Delete(urlToken string) error
	// Get retrieves the targetUrl of a shortcut described by a urlToken
	Get(urlToken string) (string, error)
}
