package time_errors

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		body     error
		response string
	}{{
		name: "OK",
		body: errors.New("test"),
		response: fmt.Sprintf(
			"\n~[ERROR]~\n[BODY]:{%s}\n[TIME]:{%s}\n[LINE]:{%d}\n[FILE]:{%s}\n",
			"test",
			time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), 0, 0, time.Local),
			32,
			"/home/root-user/Desktop/programs/go/Market/pkg/time_errors/errors_test.go",
		),
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			newError := New(test.body)
			res := newError.Error()
			assert.Equal(t, test.response, res)
		})
	}
}
