// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package types

import (
	"fmt"

	"github.com/bitcard/go-bitshares/util"
	"github.com/denkhaus/logging"
	"github.com/juju/errors"
)

type DynamicGlobalPropertyID struct {
	ObjectID
}

func (p DynamicGlobalPropertyID) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(p.Instance())); err != nil {
		return errors.Annotate(err, "encode instance")
	}

	return nil
}

func (p *DynamicGlobalPropertyID) Unmarshal(dec *util.TypeDecoder) error {
	var instance uint64
	if err := dec.DecodeUVarint(&instance); err != nil {
		return errors.Annotate(err, "decode instance")
	}

	p.number = UInt64((uint64(SpaceTypeProtocol) << 56) | (uint64(ObjectTypeDynamicGlobalProperty) << 48) | instance)
	return nil
}

type DynamicGlobalPropertyIDs []DynamicGlobalPropertyID

func (p DynamicGlobalPropertyIDs) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	for _, ex := range p {
		if err := enc.Encode(ex); err != nil {
			return errors.Annotate(err, "encode DynamicGlobalPropertyID")
		}
	}

	return nil
}

func DynamicGlobalPropertyIDFromObject(ob GrapheneObject) DynamicGlobalPropertyID {
	id, ok := ob.(*DynamicGlobalPropertyID)
	if ok {
		return *id
	}

	p := DynamicGlobalPropertyID{}
	p.MustFromObject(ob)
	if p.ObjectType() != ObjectTypeDynamicGlobalProperty {
		panic(fmt.Sprintf("invalid ObjectType: %q has no ObjectType 'ObjectTypeDynamicGlobalProperty'", p.ID()))
	}

	return p
}

//NewDynamicGlobalPropertyID creates an new DynamicGlobalPropertyID object
func NewDynamicGlobalPropertyID(id string) GrapheneObject {
	gid := new(DynamicGlobalPropertyID)
	if err := gid.Parse(id); err != nil {
		logging.Errorf(
			"DynamicGlobalPropertyID parser error %v",
			errors.Annotate(err, "Parse"),
		)
		return nil
	}

	if gid.ObjectType() != ObjectTypeDynamicGlobalProperty {
		logging.Errorf(
			"DynamicGlobalPropertyID parser error %s",
			fmt.Sprintf("%q has no ObjectType 'ObjectTypeDynamicGlobalProperty'", id),
		)
		return nil
	}

	return gid
}
