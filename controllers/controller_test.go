package controllers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	beego "github.com/beego/beego/v2/server/web"

	"TravelSphere/utils"
)

// Beego needs templates and routes registered before controller tests run.
func TestMain(m *testing.M) {
	wd, err := os.Getwd()
	if err != nil {
		panic("could not get working dir: " + err.Error())
	}
	root := filepath.Dir(wd)

	viewsPath := filepath.Join(root, "views")
	staticPath := filepath.Join(root, "static")

	beego.SetStaticPath("/static", staticPath)
	beego.BConfig.WebConfig.ViewsPath = viewsPath
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.AddFuncMap("population", utils.FormatPopulation)
	beego.TestBeegoInit(root)

	// Templates are initialized by TestBeegoInit,no further setup needed here

	beego.Router("/", &HomeController{})
	beego.Router("/countries", &CountryController{})
	beego.Router("/countries/:slug", &CountryController{}, "get:Detail")
	beego.Router("/wishlist", &WishlistController{})
	beego.Router("/dashboard", &DashboardController{})
	beego.Router("/login", &AuthController{})
	beego.Router("/logout", &AuthController{}, "get:Logout")
	beego.Router("/register", &AuthController{}, "get:RegisterForm;post:Register")
	os.Exit(m.Run())
}
func serve(method, path string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, req)
	return w
}

func pointStoresAtTemp(t *testing.T) {
	t.Helper()

	dir := t.TempDir()
	restoreWishlist := utils.SetWishlistFile(filepath.Join(dir, "wishlist.json"))
	restoreUsers := utils.SetUserFile(filepath.Join(dir, "users.json"))

	t.Cleanup(func() {
		restoreWishlist()
		restoreUsers()
	})
}
func TestHomePage(t *testing.T) {
	w := serve(http.MethodGet, "/")

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
}

func TestLoginPage(t *testing.T) {
	w := serve(http.MethodGet, "/login")

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
}

func TestRegisterForm(t *testing.T) {
	w := serve(http.MethodGet, "/register")

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
}

func TestWishlistPage(t *testing.T) {
	pointStoresAtTemp(t)

	w := serve(http.MethodGet, "/wishlist")

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
}

func TestDashboardPage(t *testing.T) {
	pointStoresAtTemp(t)

	w := serve(http.MethodGet, "/dashboard")

	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
}

func TestCountryDetailNotFound(t *testing.T) {
	w := serve(http.MethodGet, "/countries/this-is-not-a-real-place-xyz")

	if w.Code != http.StatusNotFound {
		t.Errorf("status = %d, want 404", w.Code)
	}
}
