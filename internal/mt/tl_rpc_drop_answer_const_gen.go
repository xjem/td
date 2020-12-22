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

// RPCDropAnswerRequest represents TL type `rpc_drop_answer#58e4a740`.
type RPCDropAnswerRequest struct {
	// ReqMsgID field of RPCDropAnswerRequest.
	ReqMsgID int64
}

// RPCDropAnswerRequestTypeID is TL type id of RPCDropAnswerRequest.
const RPCDropAnswerRequestTypeID = 0x58e4a740

// String implements fmt.Stringer.
func (r *RPCDropAnswerRequest) String() string {
	if r == nil {
		return "RPCDropAnswerRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("RPCDropAnswerRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tReqMsgID: ")
	sb.WriteString(fmt.Sprint(r.ReqMsgID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (r *RPCDropAnswerRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode rpc_drop_answer#58e4a740 as nil")
	}
	b.PutID(RPCDropAnswerRequestTypeID)
	b.PutLong(r.ReqMsgID)
	return nil
}

// Decode implements bin.Decoder.
func (r *RPCDropAnswerRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode rpc_drop_answer#58e4a740 to nil")
	}
	if err := b.ConsumeID(RPCDropAnswerRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode rpc_drop_answer#58e4a740: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode rpc_drop_answer#58e4a740: field req_msg_id: %w", err)
		}
		r.ReqMsgID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for RPCDropAnswerRequest.
var (
	_ bin.Encoder = &RPCDropAnswerRequest{}
	_ bin.Decoder = &RPCDropAnswerRequest{}
)

// RPCDropAnswer invokes method rpc_drop_answer#58e4a740 returning error if any.
func (c *Client) RPCDropAnswer(ctx context.Context, reqmsgid int64) (RpcDropAnswerClass, error) {
	var result RpcDropAnswerBox

	request := &RPCDropAnswerRequest{
		ReqMsgID: reqmsgid,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.RpcDropAnswer, nil
}
