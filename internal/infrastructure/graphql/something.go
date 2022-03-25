package graphql

var (
	instrument1          SomeInstrument
	instrument2          SomeInstrument
	instrument1Edge      SomeInstrumentEdge
	instrument2Edge      SomeInstrumentEdge
	instrumentConnection SomeInstrumentConnection
)

// SomeInstrument TODO
type SomeInstrument struct {
	ID string
}

// SomeInstrumentEdge TODO
type SomeInstrumentEdge struct {
	Cursor string
	Node   SomeInstrument
}

// SomeInstrumentConnection TODO
type SomeInstrumentConnection struct {
	Edges      []SomeInstrumentEdge
	Nodes      []SomeInstrument
	PageInfo   SomePageInfo
	TotalCount int
}

// SomePageInfo TODO
type SomePageInfo struct {
	EndCursor       string
	HasNextPage     bool
	HasPreviousPage bool
	StartCursor     string
}

func init() {
	instrument1 = SomeInstrument{
		ID: "1",
	}
	instrument2 = SomeInstrument{
		ID: "2",
	}

	instrument1Edge = SomeInstrumentEdge{
		Cursor: "1",
		Node:   instrument1,
	}
	instrument2Edge = SomeInstrumentEdge{
		Cursor: "2",
		Node:   instrument2,
	}

	instrumentConnection = SomeInstrumentConnection{
		Edges: []SomeInstrumentEdge{instrument1Edge, instrument2Edge},
		Nodes: []SomeInstrument{instrument1, instrument2},
		PageInfo: SomePageInfo{
			EndCursor:       "2",
			HasNextPage:     false,
			HasPreviousPage: false,
			StartCursor:     "1",
		},
		TotalCount: 2,
	}
}
