package dcs

import (
	"context"
	"crypto/rand"
	"io"
	"net"

	"golang.org/x/xerrors"

	"github.com/gotd/td/internal/mtproxy"
	"github.com/gotd/td/internal/mtproxy/obfuscator"
	"github.com/gotd/td/internal/proto/codec"
	"github.com/gotd/td/tg"
	"github.com/gotd/td/transport"
)

type mtProxy struct {
	dialer        transport.Dialer
	transport     *transport.Transport
	addr, network string

	secret mtproxy.Secret
	tag    [4]byte
	rand   io.Reader
}

func (m mtProxy) Primary(ctx context.Context, dc int, _ []tg.DCOption) (transport.Conn, error) {
	return m.resolve(ctx, dc)
}

func (m mtProxy) MediaOnly(ctx context.Context, dc int, _ []tg.DCOption) (transport.Conn, error) {
	return m.resolve(ctx, dc+10000)
}

func (m mtProxy) CDN(ctx context.Context, dc int, _ []tg.DCOption) (transport.Conn, error) {
	return m.resolve(ctx, dc)
}

func (m mtProxy) resolve(ctx context.Context, dc int) (transport.Conn, error) {
	c, err := m.dialer.DialContext(ctx, m.network, m.addr)
	if err != nil {
		return nil, xerrors.Errorf("connect to the MTProxy %q: %w", m.addr, err)
	}

	conn, err := m.handshake(c, dc)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// handshake inits given net.Conn as MTProto connection.
func (m mtProxy) handshake(c net.Conn, dc int) (transport.Conn, error) {
	var conn *obfuscator.Conn
	switch m.secret.Type {
	case mtproxy.Simple, mtproxy.Secured:
		conn = obfuscator.Obfuscated2(m.rand, c)
	case mtproxy.TLS:
		conn = obfuscator.FakeTLS(m.rand, c)
	default:
		return nil, xerrors.Errorf("unknown MTProxy secret type: %d", m.secret.Type)
	}

	secret := m.secret
	secret.DC = dc
	if err := conn.Handshake(m.tag, secret); err != nil {
		return nil, xerrors.Errorf("MTProxy handshake: %w", err)
	}

	return m.transport.Handshake(conn)
}

// MTProxyOptions is MTProxy resolver creation options.
type MTProxyOptions struct {
	// Dialer to use. net.Dialer will be used by default.
	Dialer transport.Dialer
	// Network to use.
	Network string
	// Random source for MTProxy obfuscator.
	Rand io.Reader
}

func (m *MTProxyOptions) setDefaults() {
	if m.Dialer == nil {
		m.Dialer = &net.Dialer{}
	}
	if m.Network == "" {
		m.Network = "tcp"
	}
	if m.Rand == nil {
		m.Rand = rand.Reader
	}
}

// MTProxyResolver creates MTProxy obfuscated DC resolver.
//
// See https://core.telegram.org/mtproto/mtproto-transports#transport-obfuscation.
func MTProxyResolver(addr string, secret []byte, opts MTProxyOptions) (Resolver, error) {
	s, err := mtproxy.ParseSecret(2, secret)
	if err != nil {
		return nil, err
	}

	cdc := codec.PaddedIntermediate{}
	opts.setDefaults()
	return mtProxy{
		dialer:  opts.Dialer,
		addr:    addr,
		network: opts.Network,
		transport: transport.NewTransport(func() transport.Codec {
			return codec.NoHeader{Codec: cdc}
		}),
		secret: s,
		tag:    cdc.ObfuscatedTag(),
		rand:   opts.Rand,
	}, nil
}