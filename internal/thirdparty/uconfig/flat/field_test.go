package flat_test

import (
	"errors"
	"testing"

	"github.com/arquivei/arqbeam-app/internal/thirdparty/uconfig/flat"
	"github.com/stretchr/testify/assert"
)

func TestField(t *testing.T) {
	type Config struct {
		First  string `default:"first" test:"test-tag"`
		Second error
	}

	conf := Config{}
	fs, err := flat.View(&conf)

	if err != nil {
		t.Fatal(err)
	}

	firstField := fs[0]

	if name := firstField.Name(); name != "First" {
		t.Errorf("expected First but got %v", name)
	}

	tag, ok := firstField.Tag("test")
	if !ok {
		t.Error("expected test tag on firstField but not found")
	}

	if tag != "test-tag" {
		t.Errorf("expected tag test to be test-tag but got %v", tag)
	}

	if !firstField.IsZero() {
		t.Error("expected IsZero() to return true")
	}

	meta1 := firstField.Meta()
	meta2 := firstField.Meta()

	meta1["test"] = "okay"

	assert.Equal(t, meta1, meta2)

	if def := firstField.String(); def != "first" {
		t.Errorf("expected String() to return default tag value but got %v", def)
	}

	if err := firstField.Set("some-value"); err != nil {
		t.Errorf("expected Set() to return nil but got: %v", err)
	}

	if firstField.IsZero() {
		t.Error("expected IsZero() to return false")
	}

	secondField := fs[1]

	if !secondField.IsZero() {
		t.Error("expected IsZero() to return true")
	}

	conf.Second = errors.New("oh no")

	if secondField.IsZero() {
		t.Error("expected IsZero() to return false")
	}
}
