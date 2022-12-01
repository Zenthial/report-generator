package main

import (
	"context"

	"google.golang.org/api/docs/v1"
	"google.golang.org/api/option"
)

func getDocsService() *docs.Service {
	ctx := context.Background()
	docsService, err := docs.NewService(ctx, option.WithCredentialsFile("credentials.json"))

	if err != nil {
		panic(err)
	}

	return docsService
}

func CreateNewGoogleDoc() *docs.Document {
	docsService := getDocsService()

	doc, err := docsService.Documents.Create(&docs.Document{
		Title: "My Document",
	}).Do()

	if err != nil {
		panic(err)
	}

	return doc
}

func GetGoogleDoc(documentID string) *docs.Document {
	docsService := getDocsService()

	doc, err := docsService.Documents.Get(documentID).Do()

	if err != nil {
		panic(err)
	}

	return doc
}

func MakeRequests(documentID string, requests []*docs.Request) {
	docsService := getDocsService()

	_, err := docsService.Documents.BatchUpdate(documentID, &docs.BatchUpdateDocumentRequest{
		Requests: requests,
	}).Do()

	if err != nil {
		panic(err)
	}
}

func WriteParagraph(documentID string, text string, index int64) []*docs.Request {
	return []*docs.Request{
		{
			InsertText: &docs.InsertTextRequest{
				Location: &docs.Location{
					Index: index,
				},
				Text: text,
			},
		},
	}
}

func UnderlineText(documentId string, index int64, endIndex int64) []*docs.Request {
	return []*docs.Request{
		{
			UpdateTextStyle: &docs.UpdateTextStyleRequest{
				Fields: "*",
				Range: &docs.Range{
					StartIndex: index,
					EndIndex:   endIndex,
				},
				TextStyle: &docs.TextStyle{
					Underline: true,
				},
			},
		},
	}
}

func ColorizeText(documentID string, startIndex int64, endIndex int64) []*docs.Request {
	return []*docs.Request{
		{
			UpdateTextStyle: &docs.UpdateTextStyleRequest{
				Fields: "*",
				Range: &docs.Range{
					StartIndex: startIndex,
					EndIndex:   endIndex,
				},
				TextStyle: &docs.TextStyle{
					ForegroundColor: &docs.OptionalColor{
						Color: &docs.Color{
							RgbColor: &docs.RgbColor{
								Red:   192.0 / 255.0,
								Green: 0.0,
								Blue:  0.0,
							},
						},
					},
				},
			},
		},
		{
			UpdateTextStyle: &docs.UpdateTextStyleRequest{
				Fields: "*",
				Range: &docs.Range{
					StartIndex: endIndex,
					EndIndex:   endIndex + 1,
				},
				TextStyle: &docs.TextStyle{
					ForegroundColor: &docs.OptionalColor{
						Color: &docs.Color{
							RgbColor: &docs.RgbColor{
								Red: 0,
							},
						},
					},
				},
			},
		},
	}
}

func BulletedList(documentID string, startIndex int64, endIndex int64) []*docs.Request {
	return []*docs.Request{
		{
			CreateParagraphBullets: &docs.CreateParagraphBulletsRequest{
				BulletPreset: "BULLET_DISC_CIRCLE_SQUARE",
				Range: &docs.Range{
					StartIndex: startIndex,
					EndIndex:   endIndex,
				},
			},
		},
	}
}
