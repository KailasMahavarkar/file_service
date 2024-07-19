package s3

import (
	"time"
)

var sizeRanges = map[string]FilterSizeRange{
	"0-10MB":    {0, 10 * 1024 * 1024},
	"10-100MB":  {10 * 1024 * 1024, 100 * 1024 * 1024},
	"100MB-1GB": {100 * 1024 * 1024, 1024 * 1024 * 1024},
	"1GB-10GB":  {1024 * 1024 * 1024, 10 * 1024 * 1024 * 1024},
	"10GB+":     {10 * 1024 * 1024 * 1024, -1}, // -1 represents unlimited size
}

var timeRanges = map[string]time.Duration{
	"today":        0,
	"yesterday":    -24 * time.Hour,
	"last 7 days":  -7 * 24 * time.Hour,
	"last 30 days": -30 * 24 * time.Hour,
	"last 90 days": -90 * 24 * time.Hour,
	"last 1 year":  -365 * 24 * time.Hour,
	"custom":       -365 * 24 * time.Hour, // Modify this value for your custom range
}
