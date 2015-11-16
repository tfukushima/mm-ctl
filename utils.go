// Copyright 2015 Midokura SARL
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/binary"
	"errors"

	uuid "github.com/satori/go.uuid"
	models "github.com/tfukushima/mm-ctl/org_midonet_cluster_models"
)

func convertUuidToProto(hostUuid string) (*models.UUID, error) {
	u, err := uuid.FromString(hostUuid)
	if err != nil {
		return nil, err
	}
	bytes := u.Bytes()
	msbBytes := bytes[0:7]
	lsbBytes := bytes[8:15]

	msb, nbytes := binary.Uvarint(msbBytes)
	if nbytes != 8 {
		return nil, errors.New("Invalid MSB in the UUID of the host.")
	}
	lsb, nbytes := binary.Uvarint(lsbBytes)
	if nbytes != 8 {
		return nil, errors.New("Invalid LSB in the UUID of the host.")
	}

	id := &models.UUID{}
	id.Msb = &msb
	id.Lsb = &lsb
	return id, nil
}
