package main

import (
	"context"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func GetGoogleDocFromDrive(documentID string) *drive.File {
	ctx := context.Background()
	driveService, err := drive.NewService(ctx, option.WithCredentialsFile("credentials.json"))

	if err != nil {
		panic(err)
	}

	doc, err := driveService.Files.Get(documentID).Do()

	if err != nil {
		panic(err)
	}

	return doc
}

func AddSharingPermissions(fileID string, email string) {
	ctx := context.Background()
	driveService, err := drive.NewService(ctx, option.WithCredentialsFile("credentials.json"))

	if err != nil {
		panic(err)
	}

	_, err = driveService.Permissions.Create(fileID, &drive.Permission{
		Type:         "user",
		Role:         "writer",
		EmailAddress: email,
	}).Do()

	if err != nil {
		panic(err)
	}
}

func DeleteAllFiles() {
	ctx := context.Background()
	driveService, err := drive.NewService(ctx, option.WithCredentialsFile("credentials.json"))

	if err != nil {
		panic(err)
	}

	files, err := driveService.Files.List().Do()

	if err != nil {
		panic(err)
	}

	for _, file := range files.Files {
		driveService.Files.Delete(file.Id).Do()
	}
}