package event

// ReplicaSegmentWrapped wraps a ReplicationSegment with additional information
type ReplicaSegmentWrapped struct {
	ReplicationSegment
	IDBatch     []string
	SkipIDBatch []string
	SegmentName string
}
