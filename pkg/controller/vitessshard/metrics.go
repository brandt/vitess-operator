/*
Copyright 2019 PlanetScale Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vitessshard

import (
	"planetscale.dev/vitess-operator/pkg/operator/metrics"
	"github.com/prometheus/client_golang/prometheus"

	planetscalev2 "planetscale.dev/vitess-operator/pkg/apis/planetscale/v2"
)

const (
	metricsSubsystemName = "shard"
)

var (
	shardMetricLabels = []string{
		metrics.ClusterLabel,
		metrics.KeyspaceLabel,
		metrics.ShardLabel,
		metrics.ResultLabel,
	}

	reconcileCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: metrics.Namespace,
		Subsystem: metricsSubsystemName,
		Name:      "reconcile_count",
		Help:      "Reconciliation attempts for a VitessShard",
	}, shardMetricLabels)
)

func init() {
	metrics.Registry.MustRegister(
		reconcileCount,
	)
}

func metricLabels(vts *planetscalev2.VitessShard, err error) []string {
	return []string{
		vts.Labels[planetscalev2.ClusterLabel],
		vts.Labels[planetscalev2.KeyspaceLabel],
		vts.Spec.Name,
		metrics.Result(err),
	}
}
