package controllers

import (
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestHomeControllerGet(t *testing.T) {
	c := &HomeController{}

	c.Ctx = context.NewContext()
	c.Data = make(map[interface{}]interface{})

	c.Get()

	if c.Data["Title"] != "Home" {
		t.Errorf("expected Title to be %q, got %q", "Home", c.Data["Title"])
	}

	if c.Layout != "layout.tpl" {
		t.Errorf("expected Layout to be %q, got %q", "layout.tpl", c.Layout)
	}

	if c.TplName != "home.tpl" {
		t.Errorf("expected TplName to be %q, got %q", "home.tpl", c.TplName)
	}
}
