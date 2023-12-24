package tracker

import (
	"gorm.io/gorm"
	"time"
)

type Workspace struct {
	Name       string
	Location   string
	LastLoaded time.Time
	Version    int
}

type WorkspaceMetadata struct {
	gorm.Model
	Name        string
	Description string
	Version     int
}

type Job struct {
	gorm.Model
	Name        string
	Description string
	Rate        float32
	TimeCards   []TimeCard
}

type TimeCard struct {
	gorm.Model
	Started      time.Time
	Ended        time.Time
	LastExported time.Time
	Notes        string
	JobID        uint
}
