package main

import (
	"fmt"

	pbrc "github.com/brotherlogic/recordcollection/proto"
	"golang.org/x/net/context"
)

func (s *Server) fullMatch(ctx context.Context, r1, r2 *pbrc.ReleaseMetadata) bool {
	if r1.DateAdded == 0 {
		s.RaiseIssue(ctx, "Bad Date", fmt.Sprintf("%v has no add date", r1.InstanceId), false)
		return true
	}

	if r2.DateAdded == 0 {
		s.RaiseIssue(ctx, "Bad Date", fmt.Sprintf("%v has no add date", r2.InstanceId), false)
		return true
	}

	// This is a different record entirely
	if r1.InstanceId != r2.InstanceId {
		return false
	}

	if r1.Cost != r2.Cost ||
		r1.Category != r2.Category ||
		r1.GoalFolder != r2.GoalFolder ||
		r1.SaleId != r2.SaleId {
		return false
	}

	return true

}

func (s *Server) procRecords(ctx context.Context) error {
	recs, err := s.getter.getRecords(ctx)
	if err != nil {
		return err
	}

	count := 0
	for _, r := range recs {
		if r.GetMetadata() != nil {
			count++
			var match *pbrc.ReleaseMetadata
			for _, meta := range s.config.Metadata {
				if s.fullMatch(ctx, r.GetMetadata(), meta) {
					match = meta
				}
			}

			if match == nil {
				s.config.Metadata = append(s.config.Metadata, r.GetMetadata())
			}
		}
	}

	s.Log(fmt.Sprintf("Processed %v records", count))

	s.save(ctx)
	return nil
}
