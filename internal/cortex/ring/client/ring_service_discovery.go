// Copyright (c) The Cortex Authors.
// Licensed under the Apache License 2.0.

package client

import (
	"errors"

	"github.com/thanos-io/thanos/internal/cortex/ring"
)

func NewRingServiceDiscovery(r ring.ReadRing) PoolServiceDiscovery {
	return func() ([]string, error) {
		replicationSet, err := r.GetAllHealthy(ring.Reporting)
		if errors.Is(err, ring.ErrEmptyRing) {
			return nil, nil
		}
		if err != nil {
			return nil, err
		}

		var addrs []string
		for _, instance := range replicationSet.Instances {
			addrs = append(addrs, instance.Addr)
		}
		return addrs, nil
	}
}