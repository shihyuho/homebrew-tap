package brew

import "testing"

func TestFormula_Guess(t *testing.T) {
	f, err := guess("slctl-v1.2.3-darwin.tgz")
	if err != nil {
		t.Error(err)
		t.SkipNow()
	}
	if v := f.Name; v != "slctl" {
		t.Errorf("name must be slctl, but got %s", v)
	}
	if v := f.Version; v != "v1.2.3" {
		t.Errorf("name must be v1.2.3, but got %s", v)
	}
}
