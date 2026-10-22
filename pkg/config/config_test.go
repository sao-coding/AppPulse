package config

import (
	"os"
	"path/filepath"
	"testing"
)

func load(t *testing.T, body string) (bool, int) {
	t.Helper()
	p := filepath.Join(t.TempDir(), "c.yaml")
	if err := os.WriteFile(p, []byte(body), 0644); err != nil {
		t.Fatal(err)
	}
	c, err := Load(p)
	if err != nil {
		t.Fatal(err)
	}
	return c.Filter.MsedgeInprivate, len(c.Warnings)
}

func TestLoadInPrivate(t *testing.T) {
	if v, w := load(t, "filter:\n  msedge_inprivate: true\n"); !v || w != 0 {
		t.Fatalf("new key: %v %d", v, w)
	}
	if v, w := load(t, "filter:\n  disable_msedge_inprivate_filter: false\n"); !v || w != 0 {
		t.Fatalf("legacy false should enable filter: %v %d", v, w)
	}
	if v, _ := load(t, "filter:\n  disable_msedge_inprivate_filter: true\n"); v {
		t.Fatal("legacy true should disable filter")
	}
	if _, w := load(t, "filter:\n  msedge_inprivat: true\n"); w == 0 {
		t.Fatal("typo should warn")
	}
}

func TestExampleAndRealConfig(t *testing.T) {
	for _, f := range []string{"../../config.example.yaml", "../../config.yaml"} {
		if v, w := load2(t, f); !v || w != 0 {
			t.Fatalf("%s: inprivate=%v warnings=%d", f, v, w)
		}
	}
}

func load2(t *testing.T, p string) (bool, int) {
	c, err := Load(p)
	if err != nil {
		t.Fatal(err)
	}
	return c.Filter.MsedgeInprivate, len(c.Warnings)
}
