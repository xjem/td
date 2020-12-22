// Code generated by gotdgen, DO NOT EDIT.

package mt

import (
	"context"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}

// PingDelayDisconnectRequest represents TL type `ping_delay_disconnect#f3427b8c`.
type PingDelayDisconnectRequest struct {
	// PingID field of PingDelayDisconnectRequest.
	PingID int64
	// DisconnectDelay field of PingDelayDisconnectRequest.
	DisconnectDelay int
}

// PingDelayDisconnectRequestTypeID is TL type id of PingDelayDisconnectRequest.
const PingDelayDisconnectRequestTypeID = 0xf3427b8c

// String implements fmt.Stringer.
func (p *PingDelayDisconnectRequest) String() string {
	if p == nil {
		return "PingDelayDisconnectRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PingDelayDisconnectRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tPingID: ")
	sb.WriteString(fmt.Sprint(p.PingID))
	sb.WriteString(",\n")
	sb.WriteString("\tDisconnectDelay: ")
	sb.WriteString(fmt.Sprint(p.DisconnectDelay))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (p *PingDelayDisconnectRequest) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode ping_delay_disconnect#f3427b8c as nil")
	}
	b.PutID(PingDelayDisconnectRequestTypeID)
	b.PutLong(p.PingID)
	b.PutInt(p.DisconnectDelay)
	return nil
}

// Decode implements bin.Decoder.
func (p *PingDelayDisconnectRequest) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode ping_delay_disconnect#f3427b8c to nil")
	}
	if err := b.ConsumeID(PingDelayDisconnectRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode ping_delay_disconnect#f3427b8c: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode ping_delay_disconnect#f3427b8c: field ping_id: %w", err)
		}
		p.PingID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode ping_delay_disconnect#f3427b8c: field disconnect_delay: %w", err)
		}
		p.DisconnectDelay = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PingDelayDisconnectRequest.
var (
	_ bin.Encoder = &PingDelayDisconnectRequest{}
	_ bin.Decoder = &PingDelayDisconnectRequest{}
)

// PingDelayDisconnect invokes method ping_delay_disconnect#f3427b8c returning error if any.
func (c *Client) PingDelayDisconnect(ctx context.Context, request *PingDelayDisconnectRequest) (*Pong, error) {
	var result Pong

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
