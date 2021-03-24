/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package vcwallet

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/aries-framework-go/pkg/common/log"
	"github.com/hyperledger/aries-framework-go/pkg/crypto"
	"github.com/hyperledger/aries-framework-go/pkg/doc/verifiable"
	"github.com/hyperledger/aries-framework-go/pkg/framework/aries/api/vdr"
	"github.com/hyperledger/aries-framework-go/pkg/wallet"
	"github.com/hyperledger/aries-framework-go/spi/storage"
)

var logger = log.New("aries-framework/client/vcwallet")

// ErrWalletLocked when key manager related operation attempted on locked wallet.
var ErrWalletLocked = errors.New("wallet locked")

// provider contains dependencies for the verifiable credential wallet client
// and is typically created by using aries.Context().
type provider interface {
	StorageProvider() storage.Provider
	VDRegistry() vdr.Registry
	Crypto() crypto.Crypto
}

// walletAuth is auth function which returns wallet unlock token.
type walletAuth func() (string, error)

// noAuth default auth when wallet is still locked.
// nolint:gochecknoglobals
var noAuth walletAuth = func() (string, error) { return "", ErrWalletLocked }

// Client enable access to verifiable credential wallet features.
type Client struct {
	wallet *wallet.Wallet
	auth   walletAuth
}

// New returns new verifiable credential wallet client for given user.
//
//	Args:
//		- userID : unique user identifier used for login.
//		- provider: dependencies for the verifiable credential wallet client.
//		- options : options for unlocking wallet. Any other existing wallet instance of same wallet user will be locked
//		once this instance is unlocked.
//
// returns error if wallet profile is not found.
// To create a new wallet profile, use `CreateProfile()`.
// To update an existing profile, use `UpdateProfile()`.
func New(userID string, ctx provider, options ...wallet.UnlockOptions) (*Client, error) {
	w, err := wallet.New(userID, ctx)
	if err != nil {
		return nil, err
	}

	client := &Client{wallet: w, auth: noAuth}

	if len(options) > 0 {
		if client.Close() {
			logger.Debugf("wallet was already open, existing wallet instance key manager is now closed")
		}

		err = client.Open(options...)
		if err != nil {
			return nil, err
		}
	}

	return client, nil
}

// CreateProfile creates a new verifiable credential wallet profile for given user.
// returns error if wallet profile is already created.
// Use `UpdateProfile()` for replacing an already created verifiable credential wallet profile.
func CreateProfile(userID string, ctx provider, options ...wallet.ProfileKeyManagerOptions) error {
	return wallet.CreateProfile(userID, ctx, options...)
}

// UpdateProfile updates existing verifiable credential wallet profile.
// Will create new profile if no profile exists for given user.
// Caution: you might lose your existing keys if you change kms options.
func UpdateProfile(userID string, ctx provider, options ...wallet.ProfileKeyManagerOptions) error {
	return wallet.UpdateProfile(userID, ctx, options...)
}

// Open unlocks wallet client's key manager instance and returns a token for subsequent use of wallet features.
//
//	Args:
//		- unlock options for opening wallet.
//
//	Returns token with expiry that can be used for subsequent use of wallet features.
func (c *Client) Open(options ...wallet.UnlockOptions) error {
	authToken, err := c.wallet.Open(options...)
	if err != nil {
		return err
	}

	c.auth = func() (s string, e error) {
		return authToken, nil
	}

	return nil
}

// Close expires token issued to this VC wallet client.
// returns false if token is not found or already expired for this wallet user.
func (c *Client) Close() bool {
	c.auth = noAuth

	return c.wallet.Close()
}

// Export produces a serialized exported wallet representation.
// Only ciphertext wallet contents can be exported.
//
//	Args:
//		- auth: token to be used to lock the wallet before exporting.
//
//	Returns exported locked wallet.
//
// Supported data models:
// 	- https://w3c-ccg.github.io/universal-wallet-interop-spec/#Profile
// 	- https://w3c-ccg.github.io/universal-wallet-interop-spec/#Credential
//
func (c *Client) Export(auth string) (json.RawMessage, error) {
	// TODO to be added #2433
	return nil, fmt.Errorf("to be implemented")
}

// Import Takes a serialized exported wallet representation as input
// and imports all contents into wallet.
//
//	Args:
//		- contents: wallet content to be imported.
//		- auth: token used while exporting the wallet.
//
// Supported data models:
// 	- https://w3c-ccg.github.io/universal-wallet-interop-spec/#Profile
// 	- https://w3c-ccg.github.io/universal-wallet-interop-spec/#Credential
// 	- https://w3c-ccg.github.io/universal-wallet-interop-spec/#CachedDIDDocument
//	- https://w3c-ccg.github.io/universal-wallet-interop-spec/#meta-data
//	- https://w3c-ccg.github.io/universal-wallet-interop-spec/#connection
//
func (c *Client) Import(auth string, contents json.RawMessage) error {
	// TODO to be added #2433
	return fmt.Errorf("to be implemented")
}

// Add adds given data model to wallet contents store.
//
// Supported data models:
// 	- https://w3c-ccg.github.io/universal-wallet-interop-spec/#Profile
// 	- https://w3c-ccg.github.io/universal-wallet-interop-spec/#Credential
// 	- https://w3c-ccg.github.io/universal-wallet-interop-spec/#CachedDIDDocument
//	- https://w3c-ccg.github.io/universal-wallet-interop-spec/#meta-data
//	- https://w3c-ccg.github.io/universal-wallet-interop-spec/#connection
//
// TODO: (#2433) support for correlation between wallet contents (ex: credentials to a profile/collection).
func (c *Client) Add(contentType wallet.ContentType, content json.RawMessage) error {
	return c.wallet.Add(contentType, content)
}

// Remove removes wallet content by content ID.
//
// Supported data models:
// 	- https://w3c-ccg.github.io/universal-wallet-interop-spec/#Profile
// 	- https://w3c-ccg.github.io/universal-wallet-interop-spec/#Credential
// 	- https://w3c-ccg.github.io/universal-wallet-interop-spec/#CachedDIDDocument
//	- https://w3c-ccg.github.io/universal-wallet-interop-spec/#meta-data
//	- https://w3c-ccg.github.io/universal-wallet-interop-spec/#connection
//
func (c *Client) Remove(contentType wallet.ContentType, contentID string) error {
	return c.wallet.Remove(contentType, contentID)
}

// Get fetches a wallet content by content ID.
//
// Supported data models:
// 	- https://w3c-ccg.github.io/universal-wallet-interop-spec/#Profile
// 	- https://w3c-ccg.github.io/universal-wallet-interop-spec/#Credential
// 	- https://w3c-ccg.github.io/universal-wallet-interop-spec/#CachedDIDDocument
//	- https://w3c-ccg.github.io/universal-wallet-interop-spec/#meta-data
//	- https://w3c-ccg.github.io/universal-wallet-interop-spec/#connection
//
func (c *Client) Get(contentType wallet.ContentType, contentID string) (json.RawMessage, error) {
	return c.wallet.Get(contentType, contentID)
}

// Query returns a collection of results based on current wallet contents.
//
// Supported Query Types:
// 	- https://www.w3.org/TR/json-ld11-framing
// 	- https://identity.foundation/presentation-exchange
//
func (c *Client) Query(query *wallet.QueryParams) ([]json.RawMessage, error) {
	// TODO to be added #2433
	return nil, fmt.Errorf("to be implemented")
}

// Issue adds proof to a Verifiable Credential.
//
//	Args:
//		- A verifiable credential with or without proof
//		- Proof options
//
func (c *Client) Issue(credential json.RawMessage,
	options *wallet.ProofOptions) (*verifiable.Credential, error) {
	auth, err := c.auth()
	if err != nil {
		return nil, err
	}

	return c.wallet.Issue(auth, credential, options)
}

// Prove produces a Verifiable Presentation.
//
//	Args:
//		- List of verifiable credentials IDs.
//		- Proof options
//
func (c *Client) Prove(credentialIDs []string, options *wallet.ProofOptions) (json.RawMessage, error) {
	// TODO to be added #2433
	return nil, fmt.Errorf("to be implemented")
}

// Verify takes Takes a Verifiable Credential or Verifiable Presentation as input,.
//
//	Args:
//		- a Verifiable Credential or Verifiable Presentation
//
// Returns: a boolean verified, and an error if verified is false.
func (c *Client) Verify(raw json.RawMessage) (bool, error) {
	// TODO to be added #2433
	return false, fmt.Errorf("to be implemented")
}
