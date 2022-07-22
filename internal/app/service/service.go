package service

import (
	"context"

	"github.com/romik1505/communicationGraph/internal/app/mapper"
	"github.com/romik1505/communicationGraph/internal/app/model"
	"github.com/romik1505/communicationGraph/internal/app/store"
)

type CommunicationService struct {
	store store.Storage
}

func NewCommunicationService(store store.Storage) *CommunicationService {
	return &CommunicationService{
		store: store,
	}
}

func (c CommunicationService) AddRecord(ctx context.Context, record mapper.JournalRecord) error {
	node, err := model.NewUserJournalNode(
		record.SourceID,
		record.TargetID,
		record.Time,
	)
	if err != nil {
		return err
	}

	_, err = c.store.AddJournalNode(ctx, node)

	return err
}

func (c CommunicationService) ListRecords(ctx context.Context) (mapper.Journal, error) {
	return mapper.Journal{}, nil
}

func (c CommunicationService) GetGraph(ctx context.Context) (mapper.Graph, error) {
	comms, err := c.store.GetCommunicationCounts(ctx)
	if err != nil {
		return mapper.Graph{}, err
	}
	if len(comms) == 0 {
		return mapper.Graph{}, nil
	}

	return mapper.Graph{
		Nodes: mapper.ConvertNodes(comms),
		Links: mapper.ConverterLinks(comms),
		Detail: mapper.Detail{
			Total: int(comms[0].Total.Int32),
			Max:   int(comms[0].Max.Int32),
			Min:   int(comms[0].Min.Int32),
			Avg:   float32(comms[0].Avg.Float64),
		},
	}, nil
}
