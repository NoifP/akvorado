//go:build !release

package console

import (
	"testing"

	"github.com/benbjohnson/clock"

	"akvorado/common/clickhousedb"
	"akvorado/common/clickhousedb/mocks"
	"akvorado/common/daemon"
	"akvorado/common/http"
	"akvorado/common/reporter"
	"akvorado/console/authentication"
)

// NewMock instantiantes a new authentication component
func NewMock(t *testing.T, config Configuration) (*Component, *http.Component, *mocks.MockConn, *clock.Mock) {
	t.Helper()
	r := reporter.NewMock(t)
	h := http.NewMock(t, r)
	ch, mockConn := clickhousedb.NewMock(t, r)
	mockClock := clock.NewMock()
	c, err := New(r, config, Dependencies{
		Daemon:       daemon.NewMock(t),
		HTTP:         h,
		ClickHouseDB: ch,
		Clock:        mockClock,
		Auth:         authentication.NewMock(t, r),
	})
	if err != nil {
		t.Fatalf("New() error:\n%+v", err)
	}
	return c, h, mockConn, mockClock
}
