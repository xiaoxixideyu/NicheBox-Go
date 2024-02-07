package dto

import "time"

type RebuildCacheSubjectCommentIndexMessage struct {
	CreateDate time.Time `json:"create_date"`
	SubjectID  int64     `json:"subject_id"`
}

type RebuildCacheInnerFloorCommentIndexMessage struct {
	RootIDs []int64 `json:"root_ids"`
}
