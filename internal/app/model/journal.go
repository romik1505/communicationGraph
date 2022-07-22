package model

import (
	"database/sql"
	"time"
)

type UserJournal struct {
	ID       sql.NullString `db:"id"`
	SourceID sql.NullString `db:"source_id"`
	TargetID sql.NullString `db:"target_id"`
	Time     sql.NullTime   `db:"time"`
}

func NewUserJournalNode(sourceID, targetID, tim string) (UserJournal, error) {
	t, err := time.Parse(time.RFC3339, tim)
	if err != nil {
		return UserJournal{}, err
	}

	if sourceID > targetID {
		sourceID, targetID = targetID, sourceID
	}

	return UserJournal{
		SourceID: NewNullString(sourceID),
		TargetID: NewNullString(targetID),
		Time:     NewNullTime(t),
	}, nil
}
