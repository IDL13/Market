package time_errors

import (
	"fmt"
	"runtime"
	"time"
)

// const (
// 	BadConnectionError = "Database connection error"
// 	BadDisconnectError = "Connection disconection error"
// 	InsertError        = "Error adding data"
// 	ResetError         = "Data clearing error"
// 	GettingError       = "Error getting all statistics records"
// 	DecodingError      = "Data decoding error"
// 	BrokenData         = "Broken data"
// )

type Error struct {
	body error
}

func (e *Error) Error() string {
	_, file, line, _ := runtime.Caller(1)
	time := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), 0, 0, time.Local)
	return fmt.Sprintf("\n~[ERROR]~\n[BODY]:{%s}\n[TIME]:{%s}\n[LINE]:{%d}\n[FILE]:{%s}\n", e.body, time, line, file)
}

func New(body error) Error {
	return Error{
		body: body,
	}
}
