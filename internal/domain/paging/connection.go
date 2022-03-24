package paging

// Connection TODO
type Connection interface {
	Edges() []Edge
	PageInfo() PageInfo
}

// PageInfo TODO
type PageInfo struct {
	EndCursor   Cursor
	HasNextPage bool
}

// Node TODO
type Node interface {
	ID() ID
}

// ID TODO
type ID string

// Edge TODO
type Edge interface {
	Node() Node
	Cursor() Cursor
}

// ConnectionInput TODO
type ConnectionInput struct {
	First  uint
	After  Cursor
	Last   uint
	Before Cursor
}
