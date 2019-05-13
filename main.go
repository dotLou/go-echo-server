package main

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// EchoResponse is used to response to requests with the provided request details
type EchoResponse struct {
	Headers  map[string][]string `json:"headers,omitempty"`
	Method   string              `json:"method"`
	Body     string              `json:"body,omitempty"`
	Path     string              `json:"body`
	Response string              `json:"response,omitempty"`
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	// Routes
	e.GET("/", hello)
	e.POST("/", helloPost)
	e.OPTIONS("/", helloOptions)
	e.GET("/fakeAuth", fakeAuth, middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == "valid-key", nil
	}))
	e.OPTIONS("/fakeAuth", fakeAuthOptions)

	// Start server
	go func() {
		if err := e.Start(":8080"); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func hello(c echo.Context) error {
	response := &EchoResponse{
		Headers:  c.Request().Header,
		Method:   "GET",
		Path:     "/",
		Response: "Hello, World!",
	}
	return c.JSONPretty(http.StatusOK, response, "  ")
}

func helloOptions(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderAccept, "GET, POST, OPTIONS")
	return c.String(http.StatusOK, "")
}

func helloPost(c echo.Context) error {
	//read the body
	reqBody := []byte{}
	if c.Request().Body != nil {
		reqBody, _ = ioutil.ReadAll(c.Request().Body)
	}
	// Reset the buffer
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))

	// echo back the body
	response := &EchoResponse{
		Headers: c.Request().Header,
		Method:  "POST",
		Path:    "/",
		Body:    string(reqBody),
	}
	return c.JSONPretty(http.StatusOK, response, "  ")
}

func fakeAuth(c echo.Context) error {
	response := &EchoResponse{
		Headers:  c.Request().Header,
		Method:   "GET",
		Path:     "/fakeAuth",
		Response: "Hello, World!",
	}
	return c.JSONPretty(http.StatusOK, response, "  ")
}

func fakeAuthOptions(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderAccept, "GET, OPTIONS")
	return c.String(http.StatusOK, "")
}
