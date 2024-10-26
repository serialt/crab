package crab

import (
	"testing"
	"time"

	"github.com/serialt/crab/internal"
)

func TestTime(t *testing.T) {
	assert := internal.NewAssert(t, "TestTime")
	ctime := time.Time{}

	t.Logf("date : %v", GetCrabDay(ctime))
	t.Logf("crabSecond: %v", GetCrabSecond(ctime))
	t.Logf("DateTime %v", GetDateTime(ctime))
	t.Logf("DateTimeV2 %v", GetDateTimeV2(ctime))
	cases := []string{
		GetCrabDay(ctime),
		GetCrabSecond(ctime),
		GetDateTime(ctime),
		GetDateTimeV2(ctime),
	}
	expected := []string{
		"00010101",
		"00010101000000",
		"0001-01-01 00:00:00",
		"0001/01/01 00:00:00",
	}
	for i, v := range cases {
		assert.Equal(expected[i], v)
	}
}
