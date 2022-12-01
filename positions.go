package main

import "strings"

var eboardPositions = map[string]string{
	"Hegemon": "Joshua Janicki",
	"Recruitment Officer": "Xavier Whitlow",
	"Grammateus": "Thomas Schollenberger",
	"Crysophylos": "Nate Hall",
	"Histor": "Jon Orris",
	"Pylortes": "Nic Roseamelia",
	"Hypophetes": "Sukhpal Kingra",
	"Epiprytanis": "Daniel Gardner",
	"Prytanis": "Matthew Faulkner",
}

var committeePositions = map[string]string{
	"Party": "David Sliver",
	"Social": "David Sliver",
	"Philanthropy": "Xavier Cattelona",
	"Community Service": "David Kumar",
	"Public Relations": "David Kumar & Jon Orris",
	"Programming": "Matthew Stasik",
	"Christmas Party": "Austin, Bryan, Jake",
	"Red Carnation Ball": "Brendan & Virgile",
	"Relay For Life": "David Kumar",
	"Jump": "JP Bungart",
	"Last Blast": "Jake Shulroff",
	"Dorms": "Andrew Nicolazzo",
	"Quartermaster": "Jon Orris",
	"Sports": "Coach Blau",
	"IFC": "Connor O'Shea",
	"Sweethearts": "Tobias Kiebala",
	"Restoration": "Bradley Weisfeld",
	"Blacktivities": "Xavier Whitlow",
}

type Int64Range struct {
	StartIndex int64
	EndIndex   int64
}

type Position struct {
	Person string;
	Submitted bool;
	Report string;
	RedColorRanges []Int64Range;
	BulletRanges []Int64Range;
}

func MakeEboardDictionary() map[string]Position {
	eboard := make(map[string]Position)
	for position, person := range eboardPositions {
		eboard[position] = Position{Person: person, Submitted: false, Report: "\n"}
	}
	return eboard
}

func MakeCommitteeDictionary() map[string]Position {
	committee := make(map[string]Position)
	for position, person := range committeePositions {
		committee[position] = Position{Person: person, Submitted: false, Report: "\n"}
	}
	return committee
}

const POSITION_FORMAT = "%s\n%s\n%s\n%s"

func GenerateReport(position string, report string, discussionTopics string) string {
	text := strings.Trim(position, "\n") + "\n" + "Report\n" + report + "\n" + "Discussion Topics\n" + discussionTopics + "\n"

	return text
}
