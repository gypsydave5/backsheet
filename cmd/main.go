package main

import (
	"context"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func main() {
	sheetsService, _ := sheets.NewService(context.TODO(), option.WithCredentialsFile("google-credentials.json"))
	meta, err := sheetsService.Spreadsheets.Get("1P7fkmK4l_29L9K86olO16HUVrWU5jhUkXuZ5PnzX1_0").IncludeGridData(true).Do()
	fmt.Println(len(meta.Sheets))
	for _, sheet := range meta.Sheets {
		fmt.Println(sheet.Properties.Title)
		fmt.Println(sheet.Data[0].RowData)                             // use these
		fmt.Println(sheet.Data[0].RowData[0].Values[0].EffectiveValue) // to sniff for the total data available in a sheet
		for r, row := range sheet.Data[0].RowData {
			for c, cell := range row.Values {
				fmt.Printf("%v%v: %q\n", r, c, cell.FormattedValue)
			}
		}
	}

	sheet, err := sheetsService.Spreadsheets.Values.Get("1P7fkmK4l_29L9K86olO16HUVrWU5jhUkXuZ5PnzX1_0", "Sheet1!A1:D4").Do()
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println(sheet.Values)
	for r, row := range sheet.Values {
		for c, cell := range row {
			fmt.Printf("%v%v: %q\n", r, c, cell)
		}
	}
}
