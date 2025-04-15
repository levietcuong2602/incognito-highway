package proto

import "highway/common"

func (req *RegisterRequest) SetUUID(uuid string) {
	req.UUID = uuid
}

func (req *GetBlockBeaconByHashRequest) GetCID() int32 {
	return int32(common.BEACONID)
}

func (req *GetBlockShardByHashRequest) GetCID() int32 {
	return req.Shard
}

func (req *GetBlockShardByHashRequest) SetUUID(uuid string) {
	req.UUID = uuid
}

func (req *GetBlockBeaconByHashRequest) SetUUID(uuid string) {
	req.UUID = uuid
}

func (req *GetBlockCrossShardByHashRequest) SetUUID(uuid string) {
	req.UUID = uuid
}

func (req *BlockByHeightRequest) SetUUID(uuid string) {
	req.UUID = uuid
}

func (req *BlockByHashRequest) SetUUID(uuid string) {
	req.UUID = uuid
}
