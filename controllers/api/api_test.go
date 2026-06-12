package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	beego "github.com/beego/beego/v2/server/web"

	"TravelSphere/utils"
)

func TestMain(m *testing.M) {
	wd, _ := os.Getwd()
	root := filepath.Dir(filepath.Dir(wd))

	beego.AddFuncMap("population", utils.FormatPopulation)
	beego.TestBeegoInit(root)

	beego.Router("/api/countries", &CountryController{})
	beego.Router("/api/wishlist", &WishlistController{})
	beego.Router("/api/wishlist/:id", &WishlistController{})
	beego.Router("/api/dashboard/summary", &DashboardController{})

	os.Exit(m.Run())
}

func do(method, path, body string) *httptest.ResponseRecorder {
	var req *http.Request

	if body != "" {
		req = httptest.NewRequest(method, path, strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
	} else {
		req = httptest.NewRequest(method, path, nil)
	}

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, req)

	return w
}

func tempStores(t *testing.T) {
	t.Helper()

	dir := t.TempDir()
	r1 := utils.SetWishlistFile(filepath.Join(dir, "wishlist.json"))
	r2 := utils.SetUserFile(filepath.Join(dir, "users.json"))

	t.Cleanup(func() {
		r1()
		r2()
	})
}

func mockCountries(t *testing.T) {
	t.Helper()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"data":{"objects":[
			{"names":{"common":"France"},"codes":{"alpha_2":"FR"},"capitals":[{"name":"Paris"}],"region":"Europe","coordinates":{"lat":46,"lng":2}},
			{"names":{"common":"Japan"},"codes":{"alpha_2":"JP"},"capitals":[{"name":"Tokyo"}],"region":"Asia","coordinates":{"lat":36,"lng":138}}
		],"meta":{"more":false}}}`))
	}))

	restore := utils.SetRestCountriesBaseURL(srv.URL)

	t.Cleanup(func() {
		restore()
		srv.Close()
	})
}

func TestCountriesAPI_All(t *testing.T) {
	mockCountries(t)

	w := do(http.MethodGet, "/api/countries", "")
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", w.Code)
	}

	if !strings.Contains(w.Body.String(), "France") {
		t.Error("expected France in response")
	}
}

func TestCountriesAPI_Search(t *testing.T) {
	mockCountries(t)

	w := do(http.MethodGet, "/api/countries?search=japan", "")
	body := w.Body.String()

	if !strings.Contains(body, "Japan") || strings.Contains(body, "France") {
		t.Errorf("search should return only Japan, got %s", body)
	}
}

func TestCountriesAPI_RegionFilter(t *testing.T) {
	mockCountries(t)

	w := do(http.MethodGet, "/api/countries?region=Asia", "")
	body := w.Body.String()

	if !strings.Contains(body, "Japan") || strings.Contains(body, "France") {
		t.Errorf("region filter should return only Asia, got %s", body)
	}
}

func TestCountriesAPI_ClientError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()
	defer utils.SetRestCountriesBaseURL(srv.URL)()

	w := do(http.MethodGet, "/api/countries", "")
	if w.Code != http.StatusInternalServerError {
		t.Errorf("status = %d, want 500", w.Code)
	}
}

func TestWishlistAPI_CreateReadUpdateDelete(t *testing.T) {
	tempStores(t)

	w := do(http.MethodPost, "/api/wishlist", `{"country_name":"France","note":"trip","status":"Planned"}`)
	if w.Code != http.StatusCreated {
		t.Fatalf("create status = %d, want 201", w.Code)
	}

	var created struct {
		ID string `json:"id"`
	}

	if err := json.Unmarshal(w.Body.Bytes(), &created); err != nil {
		t.Fatalf("could not parse create response: %v", err)
	}

	if created.ID == "" {
		t.Fatalf("no id in create response: %s", w.Body.String())
	}

	id := created.ID

	w = do(http.MethodGet, "/api/wishlist", "")
	if w.Code != http.StatusOK || !strings.Contains(w.Body.String(), "France") {
		t.Errorf("read failed: %d %s", w.Code, w.Body.String())
	}

	w = do(http.MethodGet, "/api/wishlist/"+id, "")
	if w.Code != http.StatusOK {
		t.Errorf("read by id status = %d, want 200", w.Code)
	}

	w = do(http.MethodPut, "/api/wishlist/"+id, `{"note":"updated","status":"Visited"}`)
	if w.Code != http.StatusOK || !strings.Contains(w.Body.String(), "updated") {
		t.Errorf("update failed: %d %s", w.Code, w.Body.String())
	}

	w = do(http.MethodDelete, "/api/wishlist/"+id, "")
	if w.Code != http.StatusOK {
		t.Errorf("delete status = %d, want 200", w.Code)
	}
}

func TestWishlistAPI_CreateBadJSON(t *testing.T) {
	tempStores(t)

	w := do(http.MethodPost, "/api/wishlist", `{not json`)
	if w.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want 400", w.Code)
	}
}

func TestWishlistAPI_CreateValidationError(t *testing.T) {
	tempStores(t)

	w := do(http.MethodPost, "/api/wishlist", `{"country_name":"","status":"Planned"}`)
	if w.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want 400", w.Code)
	}
}

func TestWishlistAPI_GetByIDNotFound(t *testing.T) {
	tempStores(t)

	w := do(http.MethodGet, "/api/wishlist/nope", "")
	if w.Code != http.StatusNotFound {
		t.Errorf("status = %d, want 404", w.Code)
	}
}

func TestWishlistAPI_UpdateBadJSON(t *testing.T) {
	tempStores(t)

	w := do(http.MethodPut, "/api/wishlist/someid", `{bad`)
	if w.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want 400", w.Code)
	}
}

func TestWishlistAPI_UpdateNotFound(t *testing.T) {
	tempStores(t)

	w := do(http.MethodPut, "/api/wishlist/missing", `{"note":"x","status":"Planned"}`)
	if w.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want 400", w.Code)
	}
}

func TestWishlistAPI_DeleteNotFound(t *testing.T) {
	tempStores(t)

	w := do(http.MethodDelete, "/api/wishlist/missing", "")
	if w.Code != http.StatusNotFound {
		t.Errorf("status = %d, want 404", w.Code)
	}
}

func TestDashboardAPI_Summary(t *testing.T) {
	tempStores(t)

	do(http.MethodPost, "/api/wishlist", `{"country_name":"France","status":"Planned"}`)

	w := do(http.MethodGet, "/api/dashboard/summary", "")
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", w.Code)
	}

	if !strings.Contains(w.Body.String(), `"total"`) {
		t.Errorf("expected total in summary, got %s", w.Body.String())
	}
}