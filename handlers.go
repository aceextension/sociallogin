package sociallogin

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)

// BeginAuth starts the authentication process for a provider
func BeginAuth(c echo.Context) error {
	provider := c.Param("provider")
	
	// Gothic expects the provider to be in the query or a custom function
	// We'll temporarily add it to the query for the handler to pick up
	req := c.Request()
	q := req.URL.Query()
	q.Add("provider", provider)
	req.URL.RawQuery = q.Encode()

	gothic.BeginAuthHandler(c.Response().Writer, req)
	return nil
}

// CompleteAuth completes the authentication process
func CompleteAuth(c echo.Context) (map[string]string, error) {
	provider := c.Param("provider")
	
	req := c.Request()
	q := req.URL.Query()
	q.Add("provider", provider)
	req.URL.RawQuery = q.Encode()

	user, err := gothic.CompleteUserAuth(c.Response().Writer, req)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"email":    user.Email,
		"name":     user.Name,
		"provider": user.Provider,
		"id":       user.UserID,
	}, nil
}

// Logout completes the logout process
func Logout(c echo.Context) error {
	gothic.Logout(c.Response().Writer, c.Request())
	c.Response().Header().Set("Location", "/")
	c.Response().WriteHeader(http.StatusTemporaryRedirect)
	return nil
}
