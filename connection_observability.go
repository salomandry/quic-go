package quic

import "net"

// UnderlyingPacketConn returns the underlying net.PacketConn for the current active path.
// This enables socket-level operations (e.g. SyscallConn, SetReadBuffer, SetWriteBuffer)
// and observability such as socket stats and buffer sizes.
//
// For connection migration, this returns the PacketConn of the current path.
// It returns nil if the underlying type does not expose a net.PacketConn.
func (c *Conn) UnderlyingPacketConn() net.PacketConn {
	return c.conn.UnderlyingPacketConn()
}
