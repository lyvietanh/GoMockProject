package util

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// set time format
const (
	dateFormat = "2006-01-02"
	timeFormat = "2006-01-02T15:04:05"
)

// Custom type
type JsonDate time.Time
type JsonDateTime time.Time

func ConvertToIsoDate(text string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", text)
	return t, err
}

func StartOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// JsonDate deserialization
func (t *JsonDate) UnmarshalJSON(data []byte) (err error) {
	s := strings.Trim(string(data), "\"")
	log.Printf("UnmarshalJSON:%v", s)
	newTime, err := time.Parse(dateFormat, s)
	*t = JsonDate(newTime)
	return
}

// JsonDate serialization
func (t JsonDate) MarshalJSON() ([]byte, error) {
	log.Println("MarshalJSON")
	timeStr := fmt.Sprintf("\"%s\"", time.Time(t).Format(dateFormat))
	return []byte(timeStr), nil
}

// JsonDateTime deserialization
func (t *JsonDateTime) UnmarshalJSON(data []byte) (err error) {
	s := strings.Trim(string(data), "\"")
	log.Printf("UnmarshalJSON:%v", s)
	newTime, err := time.Parse("\""+timeFormat+"\"", string(data))
	*t = JsonDateTime(newTime)
	return
}

// JsonDateTime serialization
func (t JsonDateTime) MarshalJSON() ([]byte, error) {
	log.Println("MarshalJSON")
	timeStr := fmt.Sprintf("\"%s\"", time.Time(t).Format(timeFormat))
	return []byte(timeStr), nil
}
