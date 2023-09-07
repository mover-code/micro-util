package validate

import (
	"testing"
	"time"
)

func TestCodeValidate(t *testing.T) {
	email := "xxx@gmail.com"
	codeV := NewCodeValidate(NewMemoryStore(0), Config{Expire: time.Second * 60 * 10, CodeType: 2})
	code, err := codeV.Generate(email)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(code, "---", email)
	isValid, err := codeV.Validate(email, code)
	if err != nil {
		t.Error(err)
		return
	}
	if !isValid {
		t.Error("Validate error")
		return
	}
}
