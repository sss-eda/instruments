package instruments

// ListInstruments TODO
func (client *Client) ListInstruments(
	first int,
	after string,
) ([]*Instrument, error) {
	request := client.ListInstrumentsRequest{
		First: first,
		After: after,
	}
	instruments := []*Instrument{}
	client.listingService.ListInstruments(
		&request,
		func(response *listing.ListInstrumentsResponse) error {
			if response.Error != nil {
				return response.Error
			}
			instruments = response.Instruments
			return nil
		},
	)

	return instruments, nil
}
