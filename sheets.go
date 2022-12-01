package main

import (
	"context"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

const REPORT_SUBMISSION string = "1cUM1UxL8crk8s3LedCRKC8xW7JIPUhhNDccE8vbR0F4"

func GetSheetsService() *sheets.Service {
	ctx := context.Background()
	sheetsService, err := sheets.NewService(ctx, option.WithCredentialsFile("credentials.json"))

	if err != nil {
		panic(err)
	}

	return sheetsService
}

func GetSheetValues(sheetsService *sheets.Service, spreadsheetID string, rangeName string) [][]interface{} {
	ctx := context.Background()
	resp, err := sheetsService.Spreadsheets.Values.Get(spreadsheetID, rangeName).Context(ctx).Do()

	if err != nil {
		panic(err)
	}

	return resp.Values
}
