// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package mntr

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"
	mbtest "github.com/elastic/beats/v7/metricbeat/mb/testing"
)

var mntrTestInputZooKeeper35 = `zk_version	3.5.3-beta-8ce24f9e675cbefffb8f21a47e06b42864475a60, built on 04/03/2017 16:19 GMT
zk_avg_latency	0
zk_max_latency	29
zk_min_latency	0
zk_packets_received	18835
zk_packets_sent	19012
zk_num_alive_connections	16
zk_outstanding_requests	0
zk_server_state	leader
zk_znode_count	489
zk_watch_count	418
zk_ephemerals_count	76
zk_approximate_data_size	706622
zk_open_file_descriptor_count	76
zk_max_file_descriptor_count	1048576
zk_followers	1
zk_synced_followers	1
zk_pending_syncs	0`

var mntrTestInputZooKeeper37 = `zk_version	3.7.0-e3704b390a6697bfdf4b0bef79e3da7a4f6bac4b, built on 2021-03-17 09:46 UTC
zk_server_state	leader
zk_peer_state	leading - broadcast
zk_ephemerals_count	80
zk_num_alive_connections	6
zk_avg_latency	0.5714
zk_min_proposal_size	36
zk_outstanding_requests	0
zk_znode_count	508
zk_global_sessions	53
zk_non_mtls_remote_conn_count	0
zk_last_client_response_size	105
zk_packets_sent	274
zk_packets_received	270
zk_max_proposal_size	8754
zk_max_client_response_size	9926
zk_synced_followers	1
zk_connection_drop_probability	0.0
zk_watch_count	113
zk_synced_non_voting_followers	0
zk_synced_observers	0
zk_auth_failed_count	0
zk_min_latency	0
zk_max_file_descriptor_count	1048576
zk_approximate_data_size	764962
zk_open_file_descriptor_count	71
zk_last_proposal_size	8600
zk_pending_syncs	0
zk_leader_uptime	648
zk_local_sessions	0
zk_uptime	749
zk_max_latency	29
zk_outstanding_tls_handshake	0
zk_learners	1
zk_min_client_response_size	16
zk_non_mtls_local_conn_count	0
zk_quorum_size	2
zk_proposal_count	6
zk_watch_bytes	253
zk_outstanding_changes_removed	4
zk_throttled_ops	0
zk_stale_requests_dropped	0
zk_large_requests_rejected	0
zk_insecure_admin_count	0
zk_connection_rejected	0
zk_cnxn_closed_without_zk_server_running	0
zk_sessionless_connections_expired	0
zk_looking_count	1
zk_dead_watchers_queued	0
zk_stale_requests	0
zk_connection_drop_count	0
zk_learner_proposal_received_count	0
zk_digest_mismatches_count	0
zk_dead_watchers_cleared	0
zk_response_packet_cache_hits	15
zk_bytes_received_count	39688
zk_add_dead_watcher_stall_time	0
zk_request_throttle_wait_count	0
zk_requests_not_forwarded_to_commit_processor	0
zk_response_packet_cache_misses	36
zk_ensemble_auth_success	0
zk_prep_processor_request_queued	214
zk_learner_commit_received_count	0
zk_stale_replies	0
zk_connection_request_count	5
zk_response_bytes	82547
zk_ensemble_auth_fail	0
zk_diff_count	1
zk_response_packet_get_children_cache_misses	0
zk_connection_revalidate_count	5
zk_quit_leading_due_to_disloyal_voter	0
zk_snap_count	0
zk_unrecoverable_error_count	0
zk_unsuccessful_handshake	0
zk_commit_count	6
zk_stale_sessions_expired	0
zk_response_packet_get_children_cache_hits	0
zk_sync_processor_request_queued	6
zk_outstanding_changes_queued	4
zk_request_commit_queued	6
zk_ensemble_auth_skip	0
zk_skip_learner_request_to_next_processor_count	0
zk_tls_handshake_exceeded	0
zk_revalidate_count	5
zk_avg_socket_closing_time	0.0
zk_min_socket_closing_time	0
zk_max_socket_closing_time	0
zk_cnt_socket_closing_time	0
zk_sum_socket_closing_time	0
zk_avg_proposal_process_time	0.0327
zk_min_proposal_process_time	0
zk_max_proposal_process_time	1
zk_cnt_proposal_process_time	214
zk_sum_proposal_process_time	7
zk_avg_leader_unavailable_time	395.0
zk_min_leader_unavailable_time	395
zk_max_leader_unavailable_time	395
zk_cnt_leader_unavailable_time	1
zk_sum_leader_unavailable_time	395
zk_avg_node_created_watch_count	0.0
zk_min_node_created_watch_count	0
zk_max_node_created_watch_count	0
zk_cnt_node_created_watch_count	0
zk_sum_node_created_watch_count	0
zk_avg_session_queues_drained	1.0
zk_min_session_queues_drained	1
zk_max_session_queues_drained	1
zk_cnt_session_queues_drained	6
zk_sum_session_queues_drained	6
zk_avg_write_commit_proc_req_queued	0.0909
zk_min_write_commit_proc_req_queued	0
zk_max_write_commit_proc_req_queued	2
zk_cnt_write_commit_proc_req_queued	220
zk_sum_write_commit_proc_req_queued	20
zk_avg_connection_token_deficit	0.0
zk_min_connection_token_deficit	0
zk_max_connection_token_deficit	0
zk_cnt_connection_token_deficit	5
zk_sum_connection_token_deficit	0
zk_avg_read_commit_proc_req_queued	0.9591
zk_min_read_commit_proc_req_queued	0
zk_max_read_commit_proc_req_queued	2
zk_cnt_read_commit_proc_req_queued	220
zk_sum_read_commit_proc_req_queued	211
zk_avg_node_deleted_watch_count	0.0
zk_min_node_deleted_watch_count	0
zk_max_node_deleted_watch_count	0
zk_cnt_node_deleted_watch_count	0
zk_sum_node_deleted_watch_count	0
zk_avg_startup_txns_load_time	24.0
zk_min_startup_txns_load_time	24
zk_max_startup_txns_load_time	24
zk_cnt_startup_txns_load_time	1
zk_sum_startup_txns_load_time	24
zk_avg_sync_processor_queue_size	0.0
zk_min_sync_processor_queue_size	0
zk_max_sync_processor_queue_size	0
zk_cnt_sync_processor_queue_size	7
zk_sum_sync_processor_queue_size	0
zk_avg_follower_sync_time	0.0
zk_min_follower_sync_time	0
zk_max_follower_sync_time	0
zk_cnt_follower_sync_time	0
zk_sum_follower_sync_time	0
zk_avg_prep_processor_queue_size	0.014
zk_min_prep_processor_queue_size	0
zk_max_prep_processor_queue_size	1
zk_cnt_prep_processor_queue_size	215
zk_sum_prep_processor_queue_size	3
zk_avg_fsynctime	0.8333
zk_min_fsynctime	0
zk_max_fsynctime	3
zk_cnt_fsynctime	6
zk_sum_fsynctime	5
zk_avg_inflight_snap_count	0.0
zk_min_inflight_snap_count	0
zk_max_inflight_snap_count	0
zk_cnt_inflight_snap_count	0
zk_sum_inflight_snap_count	0
zk_avg_reads_issued_from_session_queue	0.1667
zk_min_reads_issued_from_session_queue	0
zk_max_reads_issued_from_session_queue	1
zk_cnt_reads_issued_from_session_queue	6
zk_sum_reads_issued_from_session_queue	1
zk_avg_learner_request_processor_queue_size	0.0
zk_min_learner_request_processor_queue_size	0
zk_max_learner_request_processor_queue_size	0
zk_cnt_learner_request_processor_queue_size	0
zk_sum_learner_request_processor_queue_size	0
zk_avg_snapshottime	29.0
zk_min_snapshottime	29
zk_max_snapshottime	29
zk_cnt_snapshottime	1
zk_sum_snapshottime	29
zk_avg_unavailable_time	395.0
zk_min_unavailable_time	395
zk_max_unavailable_time	395
zk_cnt_unavailable_time	1
zk_sum_unavailable_time	395
zk_avg_startup_txns_loaded	226.0
zk_min_startup_txns_loaded	226
zk_max_startup_txns_loaded	226
zk_cnt_startup_txns_loaded	1
zk_sum_startup_txns_loaded	226
zk_avg_reads_after_write_in_session_queue	0.1667
zk_min_reads_after_write_in_session_queue	0
zk_max_reads_after_write_in_session_queue	1
zk_cnt_reads_after_write_in_session_queue	6
zk_sum_reads_after_write_in_session_queue	1
zk_avg_requests_in_session_queue	1.7692
zk_min_requests_in_session_queue	1
zk_max_requests_in_session_queue	3
zk_cnt_requests_in_session_queue	13
zk_sum_requests_in_session_queue	23
zk_avg_write_commit_proc_issued	1.0
zk_min_write_commit_proc_issued	1
zk_max_write_commit_proc_issued	1
zk_cnt_write_commit_proc_issued	6
zk_sum_write_commit_proc_issued	6
zk_avg_prep_process_time	0.028
zk_min_prep_process_time	0
zk_max_prep_process_time	3
zk_cnt_prep_process_time	214
zk_sum_prep_process_time	6
zk_avg_pending_session_queue_size	1.0
zk_min_pending_session_queue_size	1
zk_max_pending_session_queue_size	1
zk_cnt_pending_session_queue_size	6
zk_sum_pending_session_queue_size	6
zk_avg_time_waiting_empty_pool_in_commit_processor_read_ms	0.0
zk_min_time_waiting_empty_pool_in_commit_processor_read_ms	0
zk_max_time_waiting_empty_pool_in_commit_processor_read_ms	0
zk_cnt_time_waiting_empty_pool_in_commit_processor_read_ms	6
zk_sum_time_waiting_empty_pool_in_commit_processor_read_ms	0
zk_avg_commit_process_time	0.0409
zk_min_commit_process_time	0
zk_max_commit_process_time	2
zk_cnt_commit_process_time	220
zk_sum_commit_process_time	9
zk_avg_dbinittime	61.0
zk_min_dbinittime	61
zk_max_dbinittime	61
zk_cnt_dbinittime	1
zk_sum_dbinittime	61
zk_avg_inflight_diff_count	0.5
zk_min_inflight_diff_count	0
zk_max_inflight_diff_count	1
zk_cnt_inflight_diff_count	2
zk_sum_inflight_diff_count	1
zk_avg_netty_queued_buffer_capacity	0.0
zk_min_netty_queued_buffer_capacity	0
zk_max_netty_queued_buffer_capacity	0
zk_cnt_netty_queued_buffer_capacity	0
zk_sum_netty_queued_buffer_capacity	0
zk_avg_election_time	252.0
zk_min_election_time	252
zk_max_election_time	252
zk_cnt_election_time	1
zk_sum_election_time	252
zk_avg_commit_commit_proc_req_queued	0.0273
zk_min_commit_commit_proc_req_queued	0
zk_max_commit_commit_proc_req_queued	1
zk_cnt_commit_commit_proc_req_queued	220
zk_sum_commit_commit_proc_req_queued	6
zk_avg_sync_processor_batch_size	1.0
zk_min_sync_processor_batch_size	1
zk_max_sync_processor_batch_size	1
zk_cnt_sync_processor_batch_size	6
zk_sum_sync_processor_batch_size	6
zk_avg_node_children_watch_count	0.0
zk_min_node_children_watch_count	0
zk_max_node_children_watch_count	0
zk_cnt_node_children_watch_count	0
zk_sum_node_children_watch_count	0
zk_avg_write_batch_time_in_commit_processor	0.5
zk_min_write_batch_time_in_commit_processor	0
zk_max_write_batch_time_in_commit_processor	2
zk_cnt_write_batch_time_in_commit_processor	6
zk_sum_write_batch_time_in_commit_processor	3
zk_avg_read_commit_proc_issued	0.9409
zk_min_read_commit_proc_issued	0
zk_max_read_commit_proc_issued	1
zk_cnt_read_commit_proc_issued	220
zk_sum_read_commit_proc_issued	207
zk_avg_concurrent_request_processing_in_commit_processor	0.0
zk_min_concurrent_request_processing_in_commit_processor	0
zk_max_concurrent_request_processing_in_commit_processor	0
zk_cnt_concurrent_request_processing_in_commit_processor	0
zk_sum_concurrent_request_processing_in_commit_processor	0
zk_avg_node_changed_watch_count	1.0
zk_min_node_changed_watch_count	1
zk_max_node_changed_watch_count	1
zk_cnt_node_changed_watch_count	2
zk_sum_node_changed_watch_count	2
zk_avg_sync_process_time	0.3333
zk_min_sync_process_time	0
zk_max_sync_process_time	2
zk_cnt_sync_process_time	6
zk_sum_sync_process_time	2
zk_avg_startup_snap_load_time	20.0
zk_min_startup_snap_load_time	20
zk_max_startup_snap_load_time	20
zk_cnt_startup_snap_load_time	1
zk_sum_startup_snap_load_time	20
zk_avg_prep_processor_queue_time_ms	0.0514
zk_min_prep_processor_queue_time_ms	0
zk_max_prep_processor_queue_time_ms	3
zk_cnt_prep_processor_queue_time_ms	214
zk_sum_prep_processor_queue_time_ms	11
zk_p50_prep_processor_queue_time_ms	0
zk_p95_prep_processor_queue_time_ms	0
zk_p99_prep_processor_queue_time_ms	1
zk_p999_prep_processor_queue_time_ms	3
zk_avg_jvm_pause_time_ms	0.0
zk_min_jvm_pause_time_ms	0
zk_max_jvm_pause_time_ms	0
zk_cnt_jvm_pause_time_ms	0
zk_sum_jvm_pause_time_ms	0
zk_p50_jvm_pause_time_ms	0
zk_p95_jvm_pause_time_ms	0
zk_p99_jvm_pause_time_ms	0
zk_p999_jvm_pause_time_ms	0
zk_avg_close_session_prep_time	0.0
zk_min_close_session_prep_time	0
zk_max_close_session_prep_time	0
zk_cnt_close_session_prep_time	0
zk_sum_close_session_prep_time	0
zk_p50_close_session_prep_time	0
zk_p95_close_session_prep_time	0
zk_p99_close_session_prep_time	0
zk_p999_close_session_prep_time	0
zk_avg_read_commitproc_time_ms	0.2548
zk_min_read_commitproc_time_ms	0
zk_max_read_commitproc_time_ms	7
zk_cnt_read_commitproc_time_ms	208
zk_sum_read_commitproc_time_ms	53
zk_p50_read_commitproc_time_ms	0
zk_p95_read_commitproc_time_ms	2
zk_p99_read_commitproc_time_ms	5
zk_p999_read_commitproc_time_ms	7
zk_avg_updatelatency	6.0
zk_min_updatelatency	6
zk_max_updatelatency	6
zk_cnt_updatelatency	2
zk_sum_updatelatency	12
zk_p50_updatelatency	6
zk_p95_updatelatency	6
zk_p99_updatelatency	6
zk_p999_updatelatency	6
zk_avg_local_write_committed_time_ms	0.0
zk_min_local_write_committed_time_ms	0
zk_max_local_write_committed_time_ms	0
zk_cnt_local_write_committed_time_ms	6
zk_sum_local_write_committed_time_ms	0
zk_p50_local_write_committed_time_ms	0
zk_p95_local_write_committed_time_ms	0
zk_p99_local_write_committed_time_ms	0
zk_p999_local_write_committed_time_ms	0
zk_avg_request_throttle_queue_time_ms	0.0429
zk_min_request_throttle_queue_time_ms	0
zk_max_request_throttle_queue_time_ms	1
zk_cnt_request_throttle_queue_time_ms	210
zk_sum_request_throttle_queue_time_ms	9
zk_p50_request_throttle_queue_time_ms	0
zk_p95_request_throttle_queue_time_ms	0
zk_p99_request_throttle_queue_time_ms	1
zk_p999_request_throttle_queue_time_ms	1
zk_avg_readlatency	0.5192
zk_min_readlatency	0
zk_max_readlatency	8
zk_cnt_readlatency	208
zk_sum_readlatency	108
zk_p50_readlatency	0
zk_p95_readlatency	3
zk_p99_readlatency	5
zk_p999_readlatency	8
zk_avg_quorum_ack_latency	4.0
zk_min_quorum_ack_latency	1
zk_max_quorum_ack_latency	6
zk_cnt_quorum_ack_latency	6
zk_sum_quorum_ack_latency	24
zk_p50_quorum_ack_latency	5
zk_p95_quorum_ack_latency	6
zk_p99_quorum_ack_latency	6
zk_p999_quorum_ack_latency	6
zk_avg_om_commit_process_time_ms	0.0
zk_min_om_commit_process_time_ms	0
zk_max_om_commit_process_time_ms	0
zk_cnt_om_commit_process_time_ms	0
zk_sum_om_commit_process_time_ms	0
zk_p50_om_commit_process_time_ms	0
zk_p95_om_commit_process_time_ms	0
zk_p99_om_commit_process_time_ms	0
zk_p999_om_commit_process_time_ms	0
zk_avg_read_final_proc_time_ms	0.351
zk_min_read_final_proc_time_ms	0
zk_max_read_final_proc_time_ms	11
zk_cnt_read_final_proc_time_ms	208
zk_sum_read_final_proc_time_ms	73
zk_p50_read_final_proc_time_ms	0
zk_p95_read_final_proc_time_ms	1
zk_p99_read_final_proc_time_ms	8
zk_p999_read_final_proc_time_ms	11
zk_avg_commit_propagation_latency	0.0
zk_min_commit_propagation_latency	0
zk_max_commit_propagation_latency	0
zk_cnt_commit_propagation_latency	0
zk_sum_commit_propagation_latency	0
zk_p50_commit_propagation_latency	0
zk_p95_commit_propagation_latency	0
zk_p99_commit_propagation_latency	0
zk_p999_commit_propagation_latency	0
zk_avg_dead_watchers_cleaner_latency	0.0
zk_min_dead_watchers_cleaner_latency	0
zk_max_dead_watchers_cleaner_latency	0
zk_cnt_dead_watchers_cleaner_latency	0
zk_sum_dead_watchers_cleaner_latency	0
zk_p50_dead_watchers_cleaner_latency	0
zk_p95_dead_watchers_cleaner_latency	0
zk_p99_dead_watchers_cleaner_latency	0
zk_p999_dead_watchers_cleaner_latency	0
zk_avg_write_final_proc_time_ms	0.5
zk_min_write_final_proc_time_ms	0
zk_max_write_final_proc_time_ms	2
zk_cnt_write_final_proc_time_ms	6
zk_sum_write_final_proc_time_ms	3
zk_p50_write_final_proc_time_ms	0
zk_p95_write_final_proc_time_ms	2
zk_p99_write_final_proc_time_ms	2
zk_p999_write_final_proc_time_ms	2
zk_avg_proposal_ack_creation_latency	2.8333
zk_min_proposal_ack_creation_latency	0
zk_max_proposal_ack_creation_latency	6
zk_cnt_proposal_ack_creation_latency	6
zk_sum_proposal_ack_creation_latency	17
zk_p50_proposal_ack_creation_latency	3
zk_p95_proposal_ack_creation_latency	6
zk_p99_proposal_ack_creation_latency	6
zk_p999_proposal_ack_creation_latency	6
zk_avg_proposal_latency	0.0
zk_min_proposal_latency	0
zk_max_proposal_latency	0
zk_cnt_proposal_latency	0
zk_sum_proposal_latency	0
zk_p50_proposal_latency	0
zk_p95_proposal_latency	0
zk_p99_proposal_latency	0
zk_p999_proposal_latency	0
zk_avg_om_proposal_process_time_ms	0.0
zk_min_om_proposal_process_time_ms	0
zk_max_om_proposal_process_time_ms	0
zk_cnt_om_proposal_process_time_ms	0
zk_sum_om_proposal_process_time_ms	0
zk_p50_om_proposal_process_time_ms	0
zk_p95_om_proposal_process_time_ms	0
zk_p99_om_proposal_process_time_ms	0
zk_p999_om_proposal_process_time_ms	0
zk_avg_sync_processor_queue_and_flush_time_ms	1.6667
zk_min_sync_processor_queue_and_flush_time_ms	0
zk_max_sync_processor_queue_and_flush_time_ms	4
zk_cnt_sync_processor_queue_and_flush_time_ms	6
zk_sum_sync_processor_queue_and_flush_time_ms	10
zk_p50_sync_processor_queue_and_flush_time_ms	2
zk_p95_sync_processor_queue_and_flush_time_ms	4
zk_p99_sync_processor_queue_and_flush_time_ms	4
zk_p999_sync_processor_queue_and_flush_time_ms	4
zk_avg_propagation_latency	4.5
zk_min_propagation_latency	1
zk_max_propagation_latency	6
zk_cnt_propagation_latency	6
zk_sum_propagation_latency	27
zk_p50_propagation_latency	6
zk_p95_propagation_latency	6
zk_p99_propagation_latency	6
zk_p999_propagation_latency	6
zk_avg_server_write_committed_time_ms	0.0
zk_min_server_write_committed_time_ms	0
zk_max_server_write_committed_time_ms	0
zk_cnt_server_write_committed_time_ms	0
zk_sum_server_write_committed_time_ms	0
zk_p50_server_write_committed_time_ms	0
zk_p95_server_write_committed_time_ms	0
zk_p99_server_write_committed_time_ms	0
zk_p999_server_write_committed_time_ms	0
zk_avg_sync_processor_queue_time_ms	0.3333
zk_min_sync_processor_queue_time_ms	0
zk_max_sync_processor_queue_time_ms	1
zk_cnt_sync_processor_queue_time_ms	6
zk_sum_sync_processor_queue_time_ms	2
zk_p50_sync_processor_queue_time_ms	0
zk_p95_sync_processor_queue_time_ms	1
zk_p99_sync_processor_queue_time_ms	1
zk_p999_sync_processor_queue_time_ms	1
zk_avg_sync_processor_queue_flush_time_ms	1.0
zk_min_sync_processor_queue_flush_time_ms	0
zk_max_sync_processor_queue_flush_time_ms	3
zk_cnt_sync_processor_queue_flush_time_ms	6
zk_sum_sync_processor_queue_flush_time_ms	6
zk_p50_sync_processor_queue_flush_time_ms	1
zk_p95_sync_processor_queue_flush_time_ms	3
zk_p99_sync_processor_queue_flush_time_ms	3
zk_p999_sync_processor_queue_flush_time_ms	3
zk_avg_write_commitproc_time_ms	3.3333
zk_min_write_commitproc_time_ms	1
zk_max_write_commitproc_time_ms	6
zk_cnt_write_commitproc_time_ms	6
zk_sum_write_commitproc_time_ms	20
zk_p50_write_commitproc_time_ms	4
zk_p95_write_commitproc_time_ms	6
zk_p99_write_commitproc_time_ms	6
zk_p999_write_commitproc_time_ms	6
zk_avg_10_learner_handler_qp_size	0.1429
zk_min_10_learner_handler_qp_size	0
zk_max_10_learner_handler_qp_size	2
zk_cnt_10_learner_handler_qp_size	21
zk_sum_10_learner_handler_qp_size	3
zk_avg_zookeeper_read_per_namespace	212.0
zk_min_zookeeper_read_per_namespace	212
zk_max_zookeeper_read_per_namespace	212
zk_cnt_zookeeper_read_per_namespace	1
zk_sum_zookeeper_read_per_namespace	212
zk_avg_v1_read_per_namespace	430.9095
zk_min_v1_read_per_namespace	71
zk_max_v1_read_per_namespace	9951
zk_cnt_v1_read_per_namespace	199
zk_sum_v1_read_per_namespace	85751
zk_avg_zookeeper_write_per_namespace	141.3333
zk_min_zookeeper_write_per_namespace	136
zk_max_zookeeper_write_per_namespace	144
zk_cnt_zookeeper_write_per_namespace	3
zk_sum_zookeeper_write_per_namespace	424
zk_avg_v1_write_per_namespace	2464.5476
zk_min_v1_write_per_namespace	50
zk_max_v1_write_per_namespace	11669
zk_cnt_v1_write_per_namespace	126
zk_sum_v1_write_per_namespace	310533
zk_avg_10_learner_handler_qp_time_ms	1.0
zk_min_10_learner_handler_qp_time_ms	1
zk_max_10_learner_handler_qp_time_ms	1
zk_cnt_10_learner_handler_qp_time_ms	1
zk_sum_10_learner_handler_qp_time_ms	1
zk_p50_10_learner_handler_qp_time_ms	1
zk_p95_10_learner_handler_qp_time_ms	1
zk_p99_10_learner_handler_qp_time_ms	1
zk_p999_10_learner_handler_qp_time_ms	1`

func TestEventMapping(t *testing.T) {

	type TestCase struct {
		Version    string
		MntrSample string
	}

	mntrSamples := []TestCase{{"3.5.3", mntrTestInputZooKeeper35}, {"3.7.0", mntrTestInputZooKeeper37}}

	logger := logp.NewLogger("mntr_test")

	for i, sample := range mntrSamples {
		versionMsg := fmt.Sprintf("With ZooKeeper [%s]", sample.Version)

		reporter := &mbtest.CapturingReporterV2{}

		eventMapping(fmt.Sprint(i), bytes.NewReader([]byte(sample.MntrSample)), reporter, logger)

		assert.Empty(t, reporter.GetErrors(), versionMsg)

		events := reporter.GetEvents()
		assert.Len(t, events, 1, versionMsg)

		event := events[len(events)-1]

		// Check leader fields
		assert.Equal(t, event.MetricSetFields["learners"], event.MetricSetFields["followers"], versionMsg)
		assert.Equal(t, int64(1), event.MetricSetFields["learners"], versionMsg)

		// Check latency fields
		latencyFields := event.MetricSetFields["latency"].(common.MapStr)
		assert.Equal(t, float64(29), latencyFields["max"], versionMsg)
		if sample.Version != "3.5.3" {
			// Float precission is only provided from ZK 3.6
			assert.Equal(t, float64(0.5714), latencyFields["avg"], versionMsg)
		}
		assert.Equal(t, float64(0), latencyFields["min"], versionMsg)
	}

}
