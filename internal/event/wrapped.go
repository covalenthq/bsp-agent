package event

type ReplicaSegmentWrapped struct {
	ReplicationSegment
	IdBatch     []string
	SkipIDBatch []string
	SegmentName string
}
