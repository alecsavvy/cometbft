// Code generated by metricsgen. DO NOT EDIT.

package p2p

import (
	"github.com/go-kit/kit/metrics/discard"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func PrometheusMetrics(namespace string, labelsAndValues ...string) *Metrics {
	labels := []string{}
	for i := 0; i < len(labelsAndValues); i += 2 {
		labels = append(labels, labelsAndValues[i])
	}
	return &Metrics{
		Peers: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "peers",
			Help:      "Number of peers.",
		}, labels).With(labelsAndValues...),
		PeerReceiveBytesTotal: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "peer_receive_bytes_total",
			Help:      "Number of bytes received from a given peer.",
		}, append(labels, "peer_id", "chID", "message_type")).With(labelsAndValues...),
		PeerSendBytesTotal: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "peer_send_bytes_total",
			Help:      "Number of bytes sent to a given peer.",
		}, append(labels, "peer_id", "chID", "message_type")).With(labelsAndValues...),
		PeerPendingSendBytes: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "peer_pending_send_bytes",
			Help:      "Pending bytes to be sent to a given peer.",
		}, append(labels, "peer_id")).With(labelsAndValues...),
		RouterPeerQueueRecv: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "router_peer_queue_recv",
			Help:      "RouterPeerQueueRecv defines the time taken to read off of a peer's queue before sending on the connection.",
		}, labels).With(labelsAndValues...),
		RouterPeerQueueSend: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "router_peer_queue_send",
			Help:      "RouterPeerQueueSend defines the time taken to send on a peer's queue which will later be read and sent on the connection (see RouterPeerQueueRecv).",
		}, labels).With(labelsAndValues...),
		RouterChannelQueueSend: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "router_channel_queue_send",
			Help:      "RouterChannelQueueSend defines the time taken to send on a p2p channel's queue which will later be consued by the corresponding reactor/service.",
		}, labels).With(labelsAndValues...),
		PeerQueueDroppedMsgs: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "router_channel_queue_dropped_msgs",
			Help:      "PeerQueueDroppedMsgs defines the number of messages dropped from a peer's queue for a specific flow (i.e. Channel).",
		}, append(labels, "ch_id")).With(labelsAndValues...),
		PeerQueueMsgSize: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "peer_queue_msg_size",
			Help:      "PeerQueueMsgSize defines the average size of messages sent over a peer's queue for a specific flow (i.e. Channel).",
		}, append(labels, "ch_id")).With(labelsAndValues...),
	}
}

func NopMetrics() *Metrics {
	return &Metrics{
		Peers:                  discard.NewGauge(),
		PeerReceiveBytesTotal:  discard.NewCounter(),
		PeerSendBytesTotal:     discard.NewCounter(),
		PeerPendingSendBytes:   discard.NewGauge(),
		RouterPeerQueueRecv:    discard.NewHistogram(),
		RouterPeerQueueSend:    discard.NewHistogram(),
		RouterChannelQueueSend: discard.NewHistogram(),
		PeerQueueDroppedMsgs:   discard.NewCounter(),
		PeerQueueMsgSize:       discard.NewGauge(),
	}
}
