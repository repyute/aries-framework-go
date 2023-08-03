/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package common

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
	afjwt "github.com/hyperledger/aries-framework-go/component/models/jwt"
	"testing"

	"github.com/stretchr/testify/require"

	afjose "github.com/hyperledger/aries-framework-go/component/kmscrypto/doc/jose"
)

func TestVerifySigningAlgorithm(t *testing.T) {
	r := require.New(t)

	t.Run("success - EdDSA signing algorithm", func(t *testing.T) {
		headers := make(afjose.Headers)
		headers["alg"] = "EdDSA"
		err := VerifySigningAlg(headers, []string{"EdDSA"})
		r.NoError(err)
	})

	t.Run("error - signing algorithm can not be empty", func(t *testing.T) {
		headers := make(afjose.Headers)
		err := VerifySigningAlg(headers, []string{"RS256"})
		r.Error(err)
		r.Contains(err.Error(), "missing alg")
	})

	t.Run("success - EdDSA signing algorithm not in allowed list", func(t *testing.T) {
		headers := make(afjose.Headers)
		headers["alg"] = "EdDSA"
		err := VerifySigningAlg(headers, []string{"RS256"})
		r.Error(err)
		r.Contains(err.Error(), "alg 'EdDSA' is not in the allowed list")
	})

	t.Run("error - signing algorithm can not be none", func(t *testing.T) {
		headers := make(afjose.Headers)
		headers["alg"] = "none"
		err := VerifySigningAlg(headers, []string{"RS256"})
		r.Error(err)
		r.Contains(err.Error(), "alg value cannot be 'none'")
	})
}

func TestVerifyDisclosuresInSDJWT(t *testing.T) {
	r := require.New(t)

	_, privKey, err := ed25519.GenerateKey(rand.Reader)
	r.NoError(err)

	signer := afjwt.NewEd25519Signer(privKey)

	t.Run("success", func(t *testing.T) {
		sdJWT := ParseCombinedFormatForIssuance(testCombinedFormatForIssuance)
		require.Equal(t, 1, len(sdJWT.Disclosures))

		signedJWT, _, err := afjwt.Parse(sdJWT.SDJWT, afjwt.WithSignatureVerifier(&NoopSignatureVerifier{}))
		require.NoError(t, err)

		err = VerifyDisclosuresInSDJWT(sdJWT.Disclosures, signedJWT)
		r.NoError(err)
	})

	t.Run("success V5", func(t *testing.T) {
		sdJWT := ParseCombinedFormatForIssuance(testCombinedFormatForIssuanceV5)
		require.Equal(t, 6, len(sdJWT.Disclosures))

		signedJWT, _, err := afjwt.Parse(sdJWT.SDJWT, afjwt.WithSignatureVerifier(&NoopSignatureVerifier{}))
		require.NoError(t, err)

		err = VerifyDisclosuresInSDJWT(sdJWT.Disclosures, signedJWT)
		r.NoError(err)
	})

	t.Run("success - complex struct(spec example 2b)", func(t *testing.T) {
		specExample2bPresentation := fmt.Sprintf("%s%s", specExample2bJWT, specExample2bDisclosures)

		sdJWT := ParseCombinedFormatForPresentation(specExample2bPresentation)

		signedJWT, _, err := afjwt.Parse(sdJWT.SDJWT, afjwt.WithSignatureVerifier(&NoopSignatureVerifier{}))
		require.NoError(t, err)

		err = VerifyDisclosuresInSDJWT(sdJWT.Disclosures, signedJWT)
		r.NoError(err)
	})

	t.Run("success - no selective disclosures(valid case)", func(t *testing.T) {
		jwtPayload := &payload{
			Issuer: "issuer",
			SDAlg:  "sha-256",
		}

		signedJWT, err := afjwt.NewSigned(jwtPayload, nil, signer)
		r.NoError(err)

		err = VerifyDisclosuresInSDJWT(nil, signedJWT)
		r.NoError(err)
	})

	t.Run("success - selective disclosures nil", func(t *testing.T) {
		payload := make(map[string]interface{})
		payload[SDAlgorithmKey] = testAlg
		payload[SDKey] = nil

		signedJWT, err := afjwt.NewSigned(payload, nil, signer)
		r.NoError(err)

		err = VerifyDisclosuresInSDJWT(nil, signedJWT)
		r.NoError(err)
	})

	t.Run("error - disclosure not present in SD-JWT", func(t *testing.T) {
		sdJWT := ParseCombinedFormatForIssuance(testCombinedFormatForIssuance)
		require.Equal(t, 1, len(sdJWT.Disclosures))

		signedJWT, _, err := afjwt.Parse(sdJWT.SDJWT, afjwt.WithSignatureVerifier(&NoopSignatureVerifier{}))
		require.NoError(t, err)

		err = VerifyDisclosuresInSDJWT(append(sdJWT.Disclosures, additionalDisclosure), signedJWT)
		r.Error(err)
		r.Contains(err.Error(),
			"disclosure digest 'X9yH0Ajrdm1Oij4tWso9UzzKJvPoDxwmuEcO3XAdRC0' not found in SD-JWT disclosure digests")
	})

	t.Run("error - disclosure not present in SD-JWT without selective disclosures", func(t *testing.T) {
		jwtPayload := &payload{
			Issuer: "issuer",
			SDAlg:  testAlg,
		}

		signedJWT, err := afjwt.NewSigned(jwtPayload, nil, signer)
		r.NoError(err)

		err = VerifyDisclosuresInSDJWT([]string{additionalDisclosure}, signedJWT)
		r.Error(err)
		r.Contains(err.Error(),
			"disclosure digest 'X9yH0Ajrdm1Oij4tWso9UzzKJvPoDxwmuEcO3XAdRC0' not found in SD-JWT disclosure digests")
	})

	t.Run("error - missing algorithm", func(t *testing.T) {
		jwtPayload := &payload{
			Issuer: "issuer",
		}

		signedJWT, err := afjwt.NewSigned(jwtPayload, nil, signer)
		r.NoError(err)

		err = VerifyDisclosuresInSDJWT(nil, signedJWT)
		r.Error(err)
		r.Contains(err.Error(), "_sd_alg must be present in SD-JWT", SDAlgorithmKey)
	})

	t.Run("error - invalid algorithm", func(t *testing.T) {
		jwtPayload := payload{
			Issuer: "issuer",
			SDAlg:  "SHA-XXX",
		}

		signedJWT, err := afjwt.NewSigned(jwtPayload, nil, signer)
		r.NoError(err)

		err = VerifyDisclosuresInSDJWT(nil, signedJWT)
		r.Error(err)
		r.Contains(err.Error(), "_sd_alg 'SHA-XXX' not supported")
	})

	t.Run("error - algorithm is not a string", func(t *testing.T) {
		payload := make(map[string]interface{})
		payload[SDAlgorithmKey] = 18

		signedJWT, err := afjwt.NewSigned(payload, nil, signer)
		r.NoError(err)

		err = VerifyDisclosuresInSDJWT(nil, signedJWT)
		r.Error(err)
		r.Contains(err.Error(), "_sd_alg must be a string")
	})

	t.Run("error - selective disclosures must be an array", func(t *testing.T) {
		payload := make(map[string]interface{})
		payload[SDAlgorithmKey] = testAlg
		payload[SDKey] = "test"

		signedJWT, err := afjwt.NewSigned(payload, nil, signer)
		r.NoError(err)

		err = VerifyDisclosuresInSDJWT([]string{additionalDisclosure}, signedJWT)
		r.Error(err)
		r.Contains(err.Error(), "get disclosure digests: entry type[string] is not an array")
	})

	t.Run("error - selective disclosures must be a string", func(t *testing.T) {
		payload := make(map[string]interface{})
		payload[SDAlgorithmKey] = testAlg
		payload[SDKey] = []float64{123}

		signedJWT, err := afjwt.NewSigned(payload, nil, signer)
		r.NoError(err)

		err = VerifyDisclosuresInSDJWT([]string{additionalDisclosure}, signedJWT)
		r.Error(err)
		r.Contains(err.Error(), "get disclosure digests: entry item type[float64] is not a string")
	})
}
