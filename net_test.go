package crab

import (
	"net"
	"testing"

	"github.com/serialt/crab/internal"
)

func TestLenToSubNetMask(t *testing.T) {
	assert := internal.NewAssert(t, "TestLenToSubNetMask")
	want8 := "255.0.0.0"
	want32 := "255.255.255.255"
	got32 := LenToSubNetMask(32)
	assert.Equal(want32, got32)

	got8 := LenToSubNetMask(8)
	assert.Equal(want8, got8)
}

func TestIsPublicIPv4(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsPublicIPv4")

	ip1 := net.ParseIP("8.8.8.8")
	ip2 := net.ParseIP("119.29.29.29")
	ip3 := net.ParseIP("10.10.10.5")
	ip4 := net.ParseIP("127.0.0.1")

	assert.Equal(true, IsPublicIPv4(ip1))
	assert.Equal(true, IsPublicIPv4(ip2))
	assert.Equal(false, IsPublicIPv4(ip3))
	assert.Equal(false, IsPublicIPv4(ip4))
}
