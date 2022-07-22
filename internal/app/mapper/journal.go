package mapper

type Journal struct {
	Records []JournalRecord `json:"records"`
}

type JournalRecord struct {
	SourceID string `json:"source_id"`
	TargetID string `json:"target_id"`
	Time     string `json:"time"`
}
