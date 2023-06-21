package crab

import (
	"testing"

	"github.com/serialt/crab/internal"
)

func TestOsEnvOperation(t *testing.T) {
	assert := internal.NewAssert(t, "TestOsEnvOperation")

	envNotExist := GetOsEnv("foo")
	assert.Equal("", envNotExist)

	err := SetOsEnv("foo", "foo_value")
	assert.IsNil(err)

	envExist := GetOsEnv("foo")
	assert.Equal("foo_value", envExist)

	assert.Equal(true, CompareOsEnv("foo", "foo_value"))
	assert.Equal(false, CompareOsEnv("foo", "abc"))
	assert.Equal(false, CompareOsEnv("abc", "abc"))
	assert.Equal(false, CompareOsEnv("abc", "abc"))

	err = RemoveOsEnv("foo")
	if err != nil {
		t.Fail()
	}
	assert.Equal(false, CompareOsEnv("foo", "foo_value"))
}

func TestGetOsBits(t *testing.T) {
	osBits := GetOsBits()
	switch osBits {
	case 32, 64:
		t.Logf("os is %d", osBits)
	default:
		t.Error("os is not 32 or 64 bits")
	}
}
