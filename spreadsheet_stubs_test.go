package backsheet_test

import (
	"backsheet"
	"errors"
)

type StubSpreadsheet struct {
	sheet backsheet.Sheet
}

type StubMissingSpreadsheet struct {
}

func (s StubMissingSpreadsheet) Sheet(_ string) (backsheet.Sheet, error) {
	return StubSheet{}, errors.New("")
}

func (ss StubSpreadsheet) Sheet(_ string) (backsheet.Sheet, error) {
	return ss.sheet, nil
}

type StubSheet struct {
	json string
}

func (s StubSheet) ToJSON() string {
	return s.json
}

func newStubSpreadsheet(json string) StubSpreadsheet {
	return StubSpreadsheet{StubSheet{json}}
}
