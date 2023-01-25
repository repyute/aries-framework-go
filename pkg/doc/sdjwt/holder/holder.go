/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package holder

import (
	"fmt"

	"github.com/go-jose/go-jose/v3/jwt"

	"github.com/hyperledger/aries-framework-go/pkg/doc/jose"
	afgjwt "github.com/hyperledger/aries-framework-go/pkg/doc/jwt"
	"github.com/hyperledger/aries-framework-go/pkg/doc/sdjwt/common"
)

// Claim defines claim.
type Claim struct {
	Disclosure string
	Name       string
	Value      interface{}
}

// jwtParseOpts holds options for the SD-JWT parsing.
type parseOpts struct {
	detachedPayload []byte
	sigVerifier     jose.SignatureVerifier
}

// ParseOpt is the SD-JWT Parser option.
type ParseOpt func(opts *parseOpts)

// WithJWTDetachedPayload option is for definition of JWT detached payload.
func WithJWTDetachedPayload(payload []byte) ParseOpt {
	return func(opts *parseOpts) {
		opts.detachedPayload = payload
	}
}

// WithSignatureVerifier option is for definition of JWT detached payload.
func WithSignatureVerifier(signatureVerifier jose.SignatureVerifier) ParseOpt {
	return func(opts *parseOpts) {
		opts.sigVerifier = signatureVerifier
	}
}

// Parse parses issuer SD-JWT and returns claims that can be selected.
func Parse(combinedFormatForIssuance string, opts ...ParseOpt) ([]*Claim, error) {
	pOpts := &parseOpts{
		sigVerifier: &NoopSignatureVerifier{},
	}

	for _, opt := range opts {
		opt(pOpts)
	}

	var jwtOpts []afgjwt.ParseOpt
	jwtOpts = append(jwtOpts,
		afgjwt.WithSignatureVerifier(pOpts.sigVerifier),
		afgjwt.WithJWTDetachedPayload(pOpts.detachedPayload))

	cfi := common.ParseCombinedFormatForIssuance(combinedFormatForIssuance)

	signedJWT, err := afgjwt.Parse(cfi.SDJWT, jwtOpts...)
	if err != nil {
		return nil, err
	}

	err = common.VerifyDisclosuresInSDJWT(cfi.Disclosures, signedJWT)
	if err != nil {
		return nil, err
	}

	return getClaims(cfi.Disclosures)
}

func getClaims(disclosures []string) ([]*Claim, error) {
	disclosureClaims, err := common.GetDisclosureClaims(disclosures)
	if err != nil {
		return nil, fmt.Errorf("failed to get claims from disclosures: %w", err)
	}

	var claims []*Claim
	for _, disclosure := range disclosureClaims {
		claims = append(claims,
			&Claim{
				Disclosure: disclosure.Disclosure,
				Name:       disclosure.Name,
				Value:      disclosure.Value,
			})
	}

	return claims, nil
}

// BindingPayload represents holder binding payload.
type BindingPayload struct {
	Nonce    string           `json:"nonce,omitempty"`
	Audience string           `json:"aud,omitempty"`
	IssuedAt *jwt.NumericDate `json:"iat,omitempty"`
}

// BindingInfo defines holder binding payload and signer.
type BindingInfo struct {
	Payload BindingPayload
	Signer  jose.Signer
}

// options holds options for holder.
type options struct {
	holderBindingInfo *BindingInfo
}

// Option is a holder option.
type Option func(opts *options)

// WithHolderBinding option to set optional holder binding.
func WithHolderBinding(info *BindingInfo) Option {
	return func(opts *options) {
		opts.holderBindingInfo = info
	}
}

// CreatePresentation is a convenience method to assemble combined format for presentation
// using selected disclosures (claimsToDisclose) and optional holder binding.
// This call assumes that combinedFormatForIssuance has already been parsed and verified using Parse() function.
func CreatePresentation(combinedFormatForIssuance string, claimsToDisclose []string, opts ...Option) (string, error) {
	hOpts := &options{}

	for _, opt := range opts {
		opt(hOpts)
	}

	cfi := common.ParseCombinedFormatForIssuance(combinedFormatForIssuance)

	if len(cfi.Disclosures) == 0 {
		return "", fmt.Errorf("no disclosures found in SD-JWT")
	}

	disclosuresMap := sliceToMap(cfi.Disclosures)

	for _, ctd := range claimsToDisclose {
		if _, ok := disclosuresMap[ctd]; !ok {
			return "", fmt.Errorf("disclosure '%s' not found in SD-JWT", ctd)
		}
	}

	var err error

	var hbJWT string

	if hOpts.holderBindingInfo != nil {
		hbJWT, err = CreateHolderBinding(hOpts.holderBindingInfo)
		if err != nil {
			return "", fmt.Errorf("failed to create holder binding: %w", err)
		}
	}

	cf := common.CombinedFormatForPresentation{
		SDJWT:         cfi.SDJWT,
		Disclosures:   claimsToDisclose,
		HolderBinding: hbJWT,
	}

	return cf.Serialize(), nil
}

// CreateHolderBinding will create holder binding from binding info.
func CreateHolderBinding(info *BindingInfo) (string, error) {
	hbJWT, err := afgjwt.NewSigned(info.Payload, nil, info.Signer)
	if err != nil {
		return "", err
	}

	return hbJWT.Serialize(false)
}

func sliceToMap(ids []string) map[string]bool {
	// convert slice to map
	values := make(map[string]bool)
	for _, id := range ids {
		values[id] = true
	}

	return values
}

// NoopSignatureVerifier is no-op signature verifier (signature will not get checked).
type NoopSignatureVerifier struct {
}

// Verify implements signature verification.
func (sv *NoopSignatureVerifier) Verify(joseHeaders jose.Headers, payload, signingInput, signature []byte) error {
	return nil
}
