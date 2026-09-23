# NeoGCS quic-go (BBR for MoQ)

Vendored fork based on [qiulaidongfeng/quic-go](https://github.com/qiulaidongfeng/quic-go) (BBRv1 + pluggable `Config.Congestion`), with `http3.ParseCapsule` restored for [webtransport-go v0.10](https://github.com/quic-go/webtransport-go) compatibility.

Used only by the embedded MoQ WebTransport server (`internal/moqlisten`). Vehicle Manager telemetry WebTransport keeps upstream quic-go.

NeoGCS additions:

- `Conn.MaxBandwidthBitsPerSecond()` / `Conn.PacingRateBitsPerSecond()`
- Matching getters on BBRv1 / CUBIC / MCC senders
- App-limited ACK samples do not shrink BBRv1's max-bandwidth filter

Published on `master` at https://github.com/salomandry/quic-go.

Base: [qiulaidongfeng/quic-go](https://github.com/qiulaidongfeng/quic-go) `1274d131f309d485102b2c6355bdcdf0b6d674f5` (`feat: add mcc V2 dev`).

camera-manager `go.mod` replace: `github.com/quic-go/quic-go => ./third_party/quic-go`
