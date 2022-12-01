package main

import (
	"fmt"
	"strings"

	"google.golang.org/api/docs/v1"
)

func replaceDashWithTabs(s string) string {
	return strings.ReplaceAll(s, "-", "\t")
}

func main() {
	// currentYear := time.Now().Year()
	// currentDate := time.Now().Format("01/02")
	// fmt.Println(currentDate + "/" + strconv.Itoa(currentYear))
	fakeDate := "11/27/2022"

	// DeleteAllFiles()

	doc := CreateNewGoogleDoc()
	AddSharingPermissions(doc.DocumentId, "tkexiupsilongram@gmail.com")
	AddSharingPermissions(doc.DocumentId, "iamzenthial@gmail.com")

	eboardPositions := MakeEboardDictionary()
	committeePositions := MakeCommitteeDictionary()

	reportText := "Leadership\n"

	sheetsService := GetSheetsService()
	values := GetSheetValues(sheetsService, REPORT_SUBMISSION, "A1:Z1000")
	for _, row := range values {
		timestamp := strings.Split(row[0].(string), " ")[0]
		if timestamp == fakeDate {
			position := row[1].(string)
			oldPosition := strings.Clone(position)
			report := row[2].(string)
			discussion := ""
			if len(row) > 3 {
				discussion = row[3].(string)
			}

			
			report = replaceDashWithTabs(report)
			discussion = replaceDashWithTabs(discussion)

			if _, ok := eboardPositions[oldPosition]; ok {
				position += "(" + eboardPositions[oldPosition].Person + ")"
			} else if _, ok := committeePositions[oldPosition]; ok {
				position += "(" + committeePositions[oldPosition].Person + ")"
			}

			fullReportText := GenerateReport(position, report, discussion)
			
			// for both of these ranges, the actual endindex is calculated by adding the startindex, but for space saving purposes, we only store the offset
			colorRanges := []Int64Range{
				// we add 2 for the newline characters
				{StartIndex: 0, EndIndex: int64(len(position) + len("Report") + 2)},
				{StartIndex: int64(len(position) + len("Report") + 2 + len(report) + 1), EndIndex: int64(len("Discussion Topics") + 1)},
			}

			bulletRanges := []Int64Range{
				{StartIndex: int64(len(position) + len("Report") + 2), EndIndex: int64(len(report) + 1)},
				{StartIndex: int64(len(position) + len("Report") + 2 + len(report) + 1 + len("Discussion Topics") + 1), EndIndex: int64(len(discussion) + 1)},
			}

			if _, ok := eboardPositions[oldPosition]; ok {
				eboardPositions[oldPosition] = Position{Submitted: true, Report: fullReportText, RedColorRanges: colorRanges, BulletRanges: bulletRanges}
			} else if _, ok := committeePositions[oldPosition]; ok {
				committeePositions[oldPosition] = Position{Submitted: true, Report: fullReportText, RedColorRanges: colorRanges, BulletRanges: bulletRanges}
			}
		}
	}

	requests := []*docs.Request{}
	var currIndex int64 = 12

	positionFunction := func(position string, positionInfo Position) {
		if positionInfo.Submitted {
			reportText += positionInfo.Report
			for _, rangeInfo := range positionInfo.RedColorRanges {
				requests = append(requests, ColorizeText(doc.DocumentId, rangeInfo.StartIndex + currIndex, rangeInfo.EndIndex + rangeInfo.StartIndex + currIndex)...)
			}
	
			for _, rangeInfo := range positionInfo.BulletRanges {
				requests = append(requests, BulletedList(doc.DocumentId, rangeInfo.StartIndex + currIndex, rangeInfo.EndIndex + rangeInfo.StartIndex + currIndex)...)
			}
	
			currIndex += int64(len(positionInfo.Report))
		} else {
			reportText += position + "\nReport\n\nDiscussion Topics\n\n"
	
			requests = append(requests, ColorizeText(doc.DocumentId, currIndex, currIndex + int64(len(position + "\nReport\n")))...)
			currIndex += int64(len(position + "\nReport\n"))
			requests = append(requests, BulletedList(doc.DocumentId, currIndex, currIndex + int64(len("\n")))...)
			currIndex += int64(len("\n"))
			requests = append(requests, ColorizeText(doc.DocumentId, currIndex, currIndex + int64(len("Discussion Topics\n")))...)
			currIndex += int64(len("Discussion Topics\n"))
			requests = append(requests, BulletedList(doc.DocumentId, currIndex, currIndex + int64(len("\n")))...)
			currIndex += int64(len("\n"))
		}
	}

	for position, positionInfo := range eboardPositions {
		positionFunction(position, positionInfo)
	}

	reportText += "\nCommittees\n"
	currIndex += int64(len("\nCommittees\n"))
	for position, positionInfo := range committeePositions {
		positionFunction(position, positionInfo)
	}

	requests = append(WriteParagraph(doc.DocumentId, reportText, 1), requests...)
	MakeRequests(doc.DocumentId, requests)

	fmt.Println("https://docs.google.com/document/d/" + doc.DocumentId)
}
