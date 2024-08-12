package history

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

// Service represents a message service.
type Service interface {
	Get(contactURN, channelUUID string, before *time.Time, limit, page int) ([]MessagePayload, error)
	Save(msg MessagePayload) error
	StartHistoryCleaner() error
}

type service struct {
	repo Repo
}

// NewService creates and return a new message service.
func NewService(repo Repo) Service {
	return &service{
		repo: repo,
	}
}

// Get retrieves messages from the given contact URN.
// It returns a slice of messages or nil and any error ocurred while getting messages
func (s *service) Get(contactURN, channelUUID string, before *time.Time, limit, page int) ([]MessagePayload, error) {
	messages, err := s.repo.Get(contactURN, channelUUID, before, limit, page)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

// Save stores the given message. It returns any error occurred while saving.
func (s *service) Save(msg MessagePayload) error {
	err := s.repo.Save(msg)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) StartHistoryCleaner() error {
	scheduleStartTime := "01:00"
	layout := "15:05"
	go func() {
		ticker := time.NewTicker(time.Minute * 2)
		for {
			<-ticker.C
			now := time.Now()
			t, _ := time.Parse(layout, scheduleStartTime)
			startTime := time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), 0, 0, now.Location())
			if now.After(startTime) && now.Before(startTime.Add(time.Hour)) {
				ctx := context.Background()
				retentionPeriod := 24 * 30 * time.Hour // 30 days
				currentTime := time.Now()
				retentionLimit := currentTime.Add(-retentionPeriod)
				deletedCount, err := s.repo.DeleteOlderThan(ctx, retentionLimit, 1000)
				if err != nil {
					log.Error(fmt.Errorf("error on running history cleaner", err.Error()))
				} else {
					log.Info(fmt.Sprintf("deleted %d msgs from history\n", deletedCount, retentionLimit))
				}
			}
		}
	}()
	return nil
}
