// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package backsheet_test

import (
	"backsheet"
	"sync"
)

// Ensure, that SpreadsheetMock does implement backsheet.Spreadsheet.
// If this is not the case, regenerate this file with moq.
var _ backsheet.Spreadsheet = &SpreadsheetMock{}

// SpreadsheetMock is a mock implementation of backsheet.Spreadsheet.
//
// 	func TestSomethingThatUsesSpreadsheet(t *testing.T) {
//
// 		// make and configure a mocked backsheet.Spreadsheet
// 		mockedSpreadsheet := &SpreadsheetMock{
// 			SheetFunc: func(s string) (backsheet.Sheet, error) {
// 				panic("mock out the Sheet method")
// 			},
// 		}
//
// 		// use mockedSpreadsheet in code that requires backsheet.Spreadsheet
// 		// and then make assertions.
//
// 	}
type SpreadsheetMock struct {
	// SheetFunc mocks the Sheet method.
	SheetFunc func(s string) (backsheet.Sheet, error)

	// calls tracks calls to the methods.
	calls struct {
		// Sheet holds details about calls to the Sheet method.
		Sheet []struct {
			// S is the s argument value.
			S string
		}
	}
	lockSheet sync.RWMutex
}

// Sheet calls SheetFunc.
func (mock *SpreadsheetMock) Sheet(s string) (backsheet.Sheet, error) {
	callInfo := struct {
		S string
	}{
		S: s,
	}
	mock.lockSheet.Lock()
	mock.calls.Sheet = append(mock.calls.Sheet, callInfo)
	mock.lockSheet.Unlock()
	if mock.SheetFunc == nil {
		var (
			sheetOut backsheet.Sheet
			errOut   error
		)
		return sheetOut, errOut
	}
	return mock.SheetFunc(s)
}

// SheetCalls gets all the calls that were made to Sheet.
// Check the length with:
//     len(mockedSpreadsheet.SheetCalls())
func (mock *SpreadsheetMock) SheetCalls() []struct {
	S string
} {
	var calls []struct {
		S string
	}
	mock.lockSheet.RLock()
	calls = mock.calls.Sheet
	mock.lockSheet.RUnlock()
	return calls
}

// Ensure, that SheetMock does implement backsheet.Sheet.
// If this is not the case, regenerate this file with moq.
var _ backsheet.Sheet = &SheetMock{}

// SheetMock is a mock implementation of backsheet.Sheet.
//
// 	func TestSomethingThatUsesSheet(t *testing.T) {
//
// 		// make and configure a mocked backsheet.Sheet
// 		mockedSheet := &SheetMock{
// 			ToJSONFunc: func() string {
// 				panic("mock out the ToJSON method")
// 			},
// 		}
//
// 		// use mockedSheet in code that requires backsheet.Sheet
// 		// and then make assertions.
//
// 	}
type SheetMock struct {
	// ToJSONFunc mocks the ToJSON method.
	ToJSONFunc func() string

	// calls tracks calls to the methods.
	calls struct {
		// ToJSON holds details about calls to the ToJSON method.
		ToJSON []struct {
		}
	}
	lockToJSON sync.RWMutex
}

// ToJSON calls ToJSONFunc.
func (mock *SheetMock) ToJSON() string {
	callInfo := struct {
	}{}
	mock.lockToJSON.Lock()
	mock.calls.ToJSON = append(mock.calls.ToJSON, callInfo)
	mock.lockToJSON.Unlock()
	if mock.ToJSONFunc == nil {
		var (
			sOut string
		)
		return sOut
	}
	return mock.ToJSONFunc()
}

// ToJSONCalls gets all the calls that were made to ToJSON.
// Check the length with:
//     len(mockedSheet.ToJSONCalls())
func (mock *SheetMock) ToJSONCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockToJSON.RLock()
	calls = mock.calls.ToJSON
	mock.lockToJSON.RUnlock()
	return calls
}
