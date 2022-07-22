package store

import (
	"context"

	"github.com/romik1505/communicationGraph/internal/app/model"
)

func (s Storage) AddJournalNode(ctx context.Context, node model.UserJournal) (model.UserJournal, error) {
	query := s.Builder().Insert("user_journal").
		Columns("source_id", "target_id", "time").
		Values(node.SourceID, node.TargetID, node.Time).
		Suffix("RETURNING id")

	row := query.QueryRow()
	err := row.Scan(&node.ID)
	if err != nil {
		return model.UserJournal{}, err
	}

	return node, nil
}

func (s Storage) GetCommunicationCounts(ctx context.Context) ([]model.Communication, error) {
	query := s.Builder().
		Select("c.count, c.source_id, c.target_id, min(c.count) OVER(), max(c.count) OVER(), avg(c.count) OVER(), sum(c.count) OVER() as total").
		FromSelect(
			s.Builder().Select("COUNT(*) as count, source_id, target_id").
				From("user_journal").
				GroupBy("source_id", "target_id"), "c",
		)

	rows, err := query.Query()
	if err != nil {
		return nil, err
	}

	result := make([]model.Communication, 0, 10)
	var buff model.Communication

	for rows.Next() {
		err := rows.Scan(&buff.Count, &buff.SourceID, &buff.TargetID, &buff.Min, &buff.Max, &buff.Avg, &buff.Total)
		if err != nil {
			return nil, err
		}
		result = append(result, buff)
	}

	return result, nil
}
