package main

import (
	"crypto/rand"

	"github.com/agl/pond/client/ratchet"
)

type Contact struct {
	publicKey   string
	identityKey string
}

type client struct {
	publicKey   string
	identityKey string
}

// {signedPreKey:{KeyID:0 PublicKey: Signature:} IdentityKey:BY2dPfoCI05vyQIpcf1Z73HAQuWeHUhHbKhumhiNxOR6 LastResortKey:{KeyID:16777215 PublicKey:BSlICrcRxxMntiwFgMj5SMlX+hWSDHL9NIhzviT2meQh} Keys:[]}

func (c *client) newRatchet(contact *Contact) *ratchet.Ratchet {
	r := ratchet.New(rand.Reader)
	//r.MyIdentityPrivate = &c.identityKey
	//r.MySigningPublic = &c.publicKey
	//r.TheirIdentityPublic = &contact.identityKey
	//r.TheirSigningPublic = &contact.publicKey
	return r
}
