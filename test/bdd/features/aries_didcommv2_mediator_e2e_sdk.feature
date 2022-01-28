#
# Copyright SecureKey Technologies Inc. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

@all
@aries_didcommv2_router_sdk
Feature: DIDComm v2 Transport between two Agents through DIDComm v2 Routers [SDK]

  # https://identity.foundation/didcomm-messaging/spec/#routing
  Scenario Outline: Decentralized Identifier(DID) Exchange between two Edge Agents(without Inbound, DIDComm v2 is one way only) through Routers
    # DID Exchange between Alice and her Router
    Given options "<keyType>" "<keyAgreementType>" "<mediaTypeProfile>"
    # TODO note: original agent initialization ignores the above options. This refactor keeps that behaviour.
    #  Followup should add '"agent" uses configured encryption parameters' step, and validate that it passes.
    Given "Alice-Router" uses DIDComm v2
      And "Alice-Router" is started with a "websocket" DIDComm endpoint
      And "Alice-Router" creates a route exchange client
      And   "Alice-Router" creates did exchange client
      And   "Alice-Router" registers to receive notification for post state event "completed"

    Given "Alice" uses DIDComm v2
      And "Alice" is started as an edge agent
      And   "Alice" creates did exchange client
      And   "Alice" registers to receive notification for post state event "completed"

    When   "Alice-Router" creates invitation
      And   "Alice-Router" validates that invitation service endpoint of type "ws"
      And   "Alice" receives invitation from "Alice-Router"
      And   "Alice" approves invitation request
      And   "Alice-Router" approves did exchange request
      And   "Alice-Router,Alice" waits for post state event "completed"

    Then   "Alice-Router,Alice" retrieves connection record and validates that connection state is "completed"
      And   "Alice" saves connectionID to variable "Alice-router-connID"

     # DID Exchange between Bob and his Router
    Given "Bob" uses DIDComm v2
      And "Bob" is started as an edge agent
      And "Bob" creates did exchange client

    Given "Bob-Router" uses DIDComm v2
      And "Bob-Router" is started with a "websocket" DIDComm endpoint
      And "Bob-Router" creates a route exchange client
      And   "Bob-Router" creates did exchange client
      And   "Bob-Router" registers to receive notification for post state event "completed"
      And   "Bob" registers to receive notification for post state event "completed"

    When   "Bob-Router" creates invitation
      And   "Bob-Router" validates that invitation service endpoint of type "ws"
      And   "Bob" receives invitation from "Bob-Router"
      And   "Bob" approves invitation request
      And   "Bob-Router" approves did exchange request
      And   "Bob-Router,Bob" waits for post state event "completed"

    Then   "Bob-Router,Bob" retrieves connection record and validates that connection state is "completed"
      And   "Bob" saves connectionID to variable "Bob-router-connID"

       # Alice registers her Router
      And   "Alice" creates a route exchange client
      And   "Alice" sets "Alice-router-connID" as the router and "Alice-Router" approves
      And   "Alice" verifies that the router connection id is set to "Alice-router-connID"

       # Bob registers his Router
      And   "Bob" creates a route exchange client
      And   "Bob" sets "Bob-router-connID" as the router and "Bob-Router" approves
      And   "Bob" verifies that the router connection id is set to "Bob-router-connID"

     # DIDExchange between Alice and Bob through routers
    When   "Alice" creates invitation with router "Alice-router-connID"
      And   "Alice" validates that invitation service endpoint of type "ws"
      And   "Bob" receives invitation from "Alice"
      And   "Bob" approves invitation request with router "Bob-router-connID"
      And   "Alice" approves did exchange request with router "Alice-router-connID"
      And   "Alice,Bob" waits for post state event "completed"

    Then   "Alice,Bob" retrieves connection record and validates that connection state is "completed"
    Examples:
      | keyType    | keyAgreementType   | mediaTypeProfile          |
      | "ED25519"  | "X25519ECDHKW"     | "didcomm/aip1"            |
      | "ED25519"  | "X25519ECDHKW"     | "didcomm/aip2;env=rfc19"  |
      | "ED25519"  | "X25519ECDHKW"     | "didcomm/aip2;env=rfc587" |
      | "ED25519"  | "NISTP384ECDHKW"   | "didcomm/v2"              |

  Scenario Outline: DID Exchange between two Edge Agents(without Inbound, DIDComm v2 is one way only) through Routers, without router didexchange
    Given options "<keyType>" "<keyAgreementType>" "<mediaTypeProfile>"

    # DIDComm v2 connection between Alice and her Router
    Given "Alice-Router" uses http-binding did resolver url "${SIDETREE_URL}" which accepts did method "sidetree"
      And "Alice-Router" uses DIDComm v2
      And "Alice-Router" is started with a "websocket" DIDComm endpoint

      And "Alice-Router" creates a route exchange client
      And "Alice-Router" creates public DID for did method "sidetree"
      And "Alice-Router" waits for public did to become available in sidetree for up to 10 seconds

    Given "Alice" uses DIDComm v2
      And "Alice" is started as an edge agent
      And "Alice" creates did exchange client
      And "Alice" registers to receive notification for post state event "completed"

    When   "Alice-Router" creates an out-of-band-v2 invitation
      And   "Alice-Router" sends the request to "Alice", which accepts it

    Then   "Alice" validates that connection to "Alice-Router" has state "completed"
      And   "Alice" saves connectionID to variable "Alice-router-connID"

     # DIDComm v2 connection between Bob and his Router
    Given "Bob-Router" uses http-binding did resolver url "${SIDETREE_URL}" which accepts did method "sidetree"
      And "Bob-Router" uses DIDComm v2
      And "Bob-Router" is started with a "websocket" DIDComm endpoint

      And "Bob-Router" creates a route exchange client
      And "Bob-Router" creates public DID for did method "sidetree"
      And "Bob-Router" waits for public did to become available in sidetree for up to 10 seconds

    Given "Bob" uses DIDComm v2
      And "Bob" is started as an edge agent

      And   "Bob" creates did exchange client
      And   "Bob" registers to receive notification for post state event "completed"

    When   "Bob-Router" creates an out-of-band-v2 invitation
      And   "Bob-Router" sends the request to "Bob", which accepts it

    Then   "Bob" validates that connection to "Bob-Router" has state "completed"
      And   "Bob" saves connectionID to variable "Bob-router-connID"

       # Alice registers her Router
      And   "Alice" creates a route exchange client
      And   "Alice" sets "Alice-router-connID" as the router and "Alice-Router" approves
      And   "Alice" verifies that the router connection id is set to "Alice-router-connID"

       # Bob registers his Router
      And   "Bob" creates a route exchange client
      And   "Bob" sets "Bob-router-connID" as the router and "Bob-Router" approves
      And   "Bob" verifies that the router connection id is set to "Bob-router-connID"

     # DIDExchange between Alice and Bob through routers
    When   "Alice" creates invitation with router "Alice-router-connID"
      And   "Alice" validates that invitation service endpoint of type "ws"
      And   "Bob" receives invitation from "Alice"
      And   "Bob" approves invitation request with router "Bob-router-connID"
      And   "Alice" approves did exchange request with router "Alice-router-connID"
      And   "Alice,Bob" waits for post state event "completed"

    Then   "Alice,Bob" retrieves connection record and validates that connection state is "completed"

    Examples:
      | keyType              | keyAgreementType | mediaTypeProfile |
      | "ECDSAP256IEEEP1363" | "NISTP384ECDHKW" | "didcomm/v2"     |
