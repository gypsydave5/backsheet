package backsheet

//go:generate moq -out spreadsheet_mock_test.go -stub -pkg backsheet_test . Spreadsheet Sheet

type Spreadsheet interface {
	Sheet(string) (Sheet, error)
}

type Sheet interface {
	ToJSON() string
}
