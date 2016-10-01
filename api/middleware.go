package api

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine"
	"net"
	"net/http"
	"net/url"
	"strings"
)

func ValidateMiddleware(nnext echo.HandlerFunc) echo.HandlerFun {
	return func(c echo.Context) error {
		req := c.Request()

		if err := ValidateCert(req); err != nil {
			return c.JSON(
				http.StatusPreconditionFailed,
				struct {
					Error string `json:"error"`
				}{err.Error()},
			)
		}

		return next(c)
	}
}

func ValidateCert(req engine.Request) error {
	if certURL := req.Header().Get("SignatureCertChainUrl"); certURL == "" {
		return fmt.Errorf("Missing cert chain url")
	}

	if parsedURL, err := url.Parse(certURL); err != nil {
		return err
	}

	if !(strings.ToLower(parsedURL.Scheme) == "https") {
		return fmt.Errorf("cert url must be https")
	}

	if parsedURL.Host != "s3.amazonaws.com" {
		return fmt.Errorf("cert url must be from s3")
	}

	if strings.Index(parsedURL.Path, "/echo.api/") != 0 {
		return fmt.Errorf("path much start with /echo.api/")
	}

	_, p, err := net.SplitHostPort(parsedURL.Host)
	if err != nil {
		return err
	}

	if p != "443" {
		return fmt.Errorf("port must be 443")
	}

	return nil
}
