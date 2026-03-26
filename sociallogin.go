package sociallogin

import (
	"log"

	"net/http"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
)

// Config holds the configuration for social login providers
type Config struct {
	SessionSecret string
	IsProduction  bool
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
	if cfg.SessionSecret == "" {
		log.Fatal("SocialLogin ERROR: SessionSecret (SESSION_SECRET) is missing. Goth requires a valid session secret for secure cookies.")
	}

	// Initialize Session Store
	store := sessions.NewCookieStore([]byte(cfg.SessionSecret))
	store.Options.HttpOnly = true
	store.Options.Secure = cfg.IsProduction
	store.Options.SameSite = http.SameSiteLaxMode
	store.Options.Path = "/"
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
