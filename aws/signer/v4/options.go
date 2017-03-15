package v4

// Option is functional options that will dictate how a signer is specifically
// signed.
type Option func(*Signer)

// WithUnsignedPayload will return an option that will not sign
// the payload.
func WithUnsignedPayload(b bool) Option {
	return func(v4 *Signer) {
		v4.UnsignedPayload = b
	}
}
