package filters

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	beego "github.com/beego/beego/v2/server/web"
	beecontext "github.com/beego/beego/v2/server/web/context"

	"TravelSphere/utils"
)

func TestMain(m *testing.M) {
	wd, _ := os.Getwd()
	root := filepath.Dir(wd)

	beego.AddFuncMap("population", utils.FormatPopulation)
	beego.TestBeegoInit(root)

	os.Exit(m.Run())
}

func newCtx(method, path string) (*beecontext.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, path, nil)
	w := httptest.NewRecorder()

	ctx := beecontext.NewContext()
	ctx.Reset(w, req)

	return ctx, w
}

func TestLogStartSetsTime(t *testing.T) {
	ctx, _ := newCtx(http.MethodGet, "/countries")
	LogStart(ctx)

	if _, ok := ctx.Input.GetData(requestStartKey).(time.Time); !ok {
		t.Error("LogStart should store a start time")
	}
}

func TestLogFinishWithStart(t *testing.T) {
	ctx, _ := newCtx(http.MethodGet, "/countries")
	LogStart(ctx)
	LogFinish(ctx)
}

func TestLogFinishWithoutStart(t *testing.T) {
	ctx, _ := newCtx(http.MethodGet, "/countries")
	LogFinish(ctx)
}

func TestRequireAuth_NoSession(t *testing.T) {
	ctx, w := newCtx(http.MethodGet, "/wishlist")
	RequireAuth(ctx)

	if w.Code != http.StatusFound {
		t.Errorf("expected 302 redirect, got %d", w.Code)
	}

	if loc := w.Header().Get("Location"); loc != "/login" {
		t.Errorf("expected redirect to /login, got %q", loc)
	}
}

func TestRequireAuthAPI_NoSession(t *testing.T) {
	ctx, w := newCtx(http.MethodGet, "/api/wishlist")
	RequireAuthAPI(ctx)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("expected 401, got %d", w.Code)
	}
}