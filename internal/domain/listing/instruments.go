package listing

import (
	"context"

	"github.com/sss-eda/instruments/internal/domain/paging"
)

// Instruments TODO
func Instruments(
	repository InstrumentRepository,
) func(context.Context, *InstrumentsInput) (*InstrumentsOutput, error) {
	return func(
		ctx context.Context,
		input *InstrumentsInput,
	) (*InstrumentsOutput, error) {
		instruments, err := repository.GetInstruments(ctx, input.ConnectionInput)
		if err != nil {
			// TODO: Return an error, but also return the available instruments
			return nil, err
		}

		edges := make([]*InstrumentsEdge, len(instruments))
		for i, instrument := range instruments {
			edges[i] = &InstrumentsEdge{
				node:   instrument,
				cursor: paging.Cursor(instrument.ID()),
			}
		}

		return &InstrumentsOutput{
			edges: edges,
			pageInfo: &paging.PageInfo{
				EndCursor:   edges[len(edges)-1].cursor,
				HasNextPage: false,
			},
		}, nil
	}
}

// InstrumentsInput TODO
type InstrumentsInput struct {
	paging.ConnectionInput
}

// InstrumentsOutput TODO
type InstrumentsOutput InstrumentsConnection

// InstrumentsConnection TODO
type InstrumentsConnection struct {
	edges    []*InstrumentsEdge
	pageInfo *paging.PageInfo
}

// InstrumentsEdge - Because we define the edge here, we also need to defined
// the cursor here: It is used to uniquely identify and EDGE and not a NODE! So
// it could be a combination of values from different places.
type InstrumentsEdge struct {
	node   *Instrument
	cursor paging.Cursor
}

// Edges TODO
func (connection *InstrumentsConnection) Edges() []*InstrumentsEdge {
	return connection.edges
}

// PageInfo TODO
func (connection *InstrumentsConnection) PageInfo() *paging.PageInfo {
	return connection.pageInfo
}
