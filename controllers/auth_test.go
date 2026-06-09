package controllers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	beego "github.com/beego/beego/v2/server/web"

	"TravelSphere/services"
)

func postForm(method, path string, data url.Values) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, req)
	return w
}

func TestAuthLoginSuccessAndFailure(t *testing.T) {
	pointStoresAtTemp(t)
	_, err := services.RegisterUser("tester", "hunter2")
	if err != nil {
		t.Fatalf("could not create user: %v", err)
	}
	w := postForm(http.MethodPost, "/login", url.Values{"username": {"tester"}, "password": {"wrong"}})
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 for failed login render, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "Invalid username or password") {
		t.Fatalf("expected error message in body")
	}
	w = postForm(http.MethodPost, "/login", url.Values{"username": {"tester"}, "password": {"hunter2"}})
	if w.Code != http.StatusFound && w.Code != http.StatusSeeOther && w.Code != http.StatusMovedPermanently {
		t.Fatalf("expected redirect for successful login, got %d", w.Code)
	}
	loc := w.Header().Get("Location")
	if loc != "/dashboard" {
		t.Fatalf("expected redirect to /dashboard, got %q", loc)
	}
}
func TestAuthRegisterAndLogout(t *testing.T) {
	pointStoresAtTemp(t)
	w := postForm(http.MethodPost, "/register", url.Values{"username": {"newuser"}, "password": {"abc123"}, "confirm": {"different"}})
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 for register form error, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "Passwords do not match") {
		t.Fatalf("expected mismatch message in body")
	}
	w = postForm(http.MethodPost, "/register", url.Values{"username": {"newuser"}, "password": {"abc123"}, "confirm": {"abc123"}})
	if w.Code != http.StatusFound {
		t.Fatalf("expected redirect after register, got %d", w.Code)
	}
	if w.Header().Get("Location") != "/dashboard" {
		t.Fatalf("expected redirect to /dashboard, got %q", w.Header().Get("Location"))
	}

	// Logout
	w = serve(http.MethodGet, "/logout")
	if w.Code != http.StatusFound {
		t.Fatalf("expected redirect after logout, got %d", w.Code)
	}
	if w.Header().Get("Location") != "/" {
		t.Fatalf("expected redirect to /, got %q", w.Header().Get("Location"))
	}
}
func TestHomePageAfterLogin(t *testing.T) {
	pointStoresAtTemp(t)

	_, err := services.RegisterUser("sessionuser", "pa55word")
	if err != nil {
		t.Fatalf("could not create user: %v", err)
	}

	// login and capture cookie
	w := postForm(http.MethodPost, "/login", url.Values{"username": {"sessionuser"}, "password": {"pa55word"}})
	if w.Code != http.StatusFound {
		t.Fatalf("expected redirect for login, got %d", w.Code)
	}
	cookie := w.Header().Get("Set-Cookie")

	// request home with session cookie
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	if cookie != "" {
		req.Header.Set("Cookie", cookie)
	}
	rw := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(rw, req)
	if rw.Code != http.StatusOK {
		t.Fatalf("expected 200 for home after login, got %d", rw.Code)
	}
	if !strings.Contains(rw.Body.String(), "Hi, sessionuser") {
		t.Fatalf("expected greeting for logged in user in body")
	}
}
