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

type AssetDynamicDataID struct {
	ObjectID
}

func (p AssetDynamicDataID) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(p.Instance())); err != nil {
		return errors.Annotate(err, "encode instance")
	}

	return nil
}

func (p *AssetDynamicDataID) Unmarshal(dec *util.TypeDecoder) error {
	var instance uint64
	if err := dec.DecodeUVarint(&instance); err != nil {
		return errors.Annotate(err, "decode instance")
	}

	p.number = UInt64((uint64(SpaceTypeProtocol) << 56) | (uint64(ObjectTypeAssetDynamicData) << 48) | instance)
	return nil
}

type AssetDynamicDataIDs []AssetDynamicDataID

func (p AssetDynamicDataIDs) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	for _, ex := range p {
		if err := enc.Encode(ex); err != nil {
			return errors.Annotate(err, "encode AssetDynamicDataID")
		}
	}

	return nil
}

func AssetDynamicDataIDFromObject(ob GrapheneObject) AssetDynamicDataID {
	id, ok := ob.(*AssetDynamicDataID)
	if ok {
		return *id
	}

	p := AssetDynamicDataID{}
	p.MustFromObject(ob)
	if p.ObjectType() != ObjectTypeAssetDynamicData {
		panic(fmt.Sprintf("invalid ObjectType: %q has no ObjectType 'ObjectTypeAssetDynamicData'", p.ID()))
	}

	return p
}

//NewAssetDynamicDataID creates an new AssetDynamicDataID object
func NewAssetDynamicDataID(id string) GrapheneObject {
	gid := new(AssetDynamicDataID)
	if err := gid.Parse(id); err != nil {
		logging.Errorf(
			"AssetDynamicDataID parser error %v",
			errors.Annotate(err, "Parse"),
		)
		return nil
	}

	if gid.ObjectType() != ObjectTypeAssetDynamicData {
		logging.Errorf(
			"AssetDynamicDataID parser error %s",
			fmt.Sprintf("%q has no ObjectType 'ObjectTypeAssetDynamicData'", id),
		)
		return nil
	}

	return gid
}
