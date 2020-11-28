// Code generated by gotdgen, DO NOT EDIT.

package td

import (
	"context"
	"fmt"

	"github.com/ernado/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// GetUpdatesResp represents TL type `getUpdatesResp#2b4b45c`.
type GetUpdatesResp struct {
	// Updates field of GetUpdatesResp.
	Updates []AbstractMessageClass
}

// GetUpdatesRespTypeID is TL type id of GetUpdatesResp.
const GetUpdatesRespTypeID = 0x2b4b45c

// Encode implements bin.Encoder.
func (g *GetUpdatesResp) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getUpdatesResp#2b4b45c as nil")
	}
	b.PutID(GetUpdatesRespTypeID)
	b.PutVectorHeader(len(g.Updates))
	for idx, v := range g.Updates {
		if v == nil {
			return fmt.Errorf("unable to encode getUpdatesResp#2b4b45c: field updates element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode getUpdatesResp#2b4b45c: field updates element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetUpdatesResp) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getUpdatesResp#2b4b45c to nil")
	}
	if err := b.ConsumeID(GetUpdatesRespTypeID); err != nil {
		return fmt.Errorf("unable to decode getUpdatesResp#2b4b45c: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode getUpdatesResp#2b4b45c: field updates: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeAbstractMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode getUpdatesResp#2b4b45c: field updates: %w", err)
			}
			g.Updates = append(g.Updates, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for GetUpdatesResp.
var (
	_ bin.Encoder = &GetUpdatesResp{}
	_ bin.Decoder = &GetUpdatesResp{}
)