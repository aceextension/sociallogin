package sociallogin

import (
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/github"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth/gothic"
)

// Config holds the configuration for social login providers
type Config struct {
	SessionSecret string
	Google        *ProviderConfig
	GitHub        *ProviderConfig
}

type ProviderConfig struct {
	ClientKey    string
	ClientSecret string
	CallbackURL  string
}

// Init initializes Goth with the provided configuration
func Init(cfg Config) {
	// Initialize Session Store
	store := sessions.NewCookieStore([]byte(cfg.SessionSecret))
	store.Options.HttpOnly = true
	store.Options.Secure = false // Set to true in production
	gothic.Store = store

	var providers []goth.Provider

	if cfg.Google != nil {
		providers = append(providers, google.New(cfg.Google.ClientKey, cfg.Google.ClientSecret, cfg.Google.CallbackURL))
	}

	if cfg.GitHub != nil {
		providers = append(providers, github.New(cfg.GitHub.ClientKey, cfg.GitHub.ClientSecret, cfg.GitHub.CallbackURL))
	}

	goth.UseProviders(providers...)
}
