package v4

// Option is functional options that will dictate how a signer is specifically
// signed.
type Option func(*Signer)

// WithUnsignedPayload will enable and set the UnsignedPayload field to
// true of the signer.
func WithUnsignedPayload(v4 *Signer) {
	v4.UnsignedPayload = true
}
