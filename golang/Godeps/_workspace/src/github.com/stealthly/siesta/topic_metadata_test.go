/* Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License. */

package siesta

import "testing"

var emptyMetadataRequestBytes = []byte{0x00, 0x00, 0x00, 0x00}
var asdMetadataRequestBytes = []byte{0x00, 0x00, 0x00, 0x01, 0x00, 0x03, 0x61, 0x73, 0x64}
var multipleTopicsMetadataRequestBytes = []byte{0x00, 0x00, 0x00, 0x03, 0x00, 0x03, 0x61, 0x73, 0x64, 0x00, 0x03, 0x7a, 0x78, 0x63, 0x00, 0x03, 0x71, 0x77, 0x65}

var emptyMetadataResponseBytes = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
var brokerMetadataResponseBytes = []byte{0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x00, 0x00, 0x23, 0x84, 0x00, 0x00, 0x00, 0x00}
var topicMetadataResponseBytes = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x05, 0x6c, 0x6f, 0x67, 0x73, 0x31, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00}

var invalidBrokersLengthMetadataResponseBytes = []byte{0x00, 0x00, 0x01}
var invalidMetadataLengthMetadataResponseBytes = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}
var invalidBrokerNodeIdMetadataResponseBytes = []byte{0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00}
var invalidBrokerHostMetadataResponseBytes = []byte{0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f}
var invalidBrokerPortMetadataResponseBytes = []byte{0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x00, 0x00, 0x23}
var invalidTopicMetadataErrorCodeMetadataResponseBytes = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00}
var invalidTopicMetadataTopicNameMetadataResponseBytes = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x05, 0x6c, 0x6f, 0x67}
var invalidPartitionMetadataLengthMetadataResponseBytes = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x05, 0x6c, 0x6f, 0x67, 0x73, 0x31, 0x00, 0x00, 0x00}
var invalidPartitionMetadataErrorCodeMetadataResponseBytes = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x05, 0x6c, 0x6f, 0x67, 0x73, 0x31, 0x00, 0x00, 0x00, 0x02, 0x00}
var invalidPartitionMetadataPartitionMetadataResponseBytes = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x05, 0x6c, 0x6f, 0x67, 0x73, 0x31, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00}
var invalidPartitionMetadataLeaderMetadataResponseBytes = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x05, 0x6c, 0x6f, 0x67, 0x73, 0x31, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00}
var invalidPartitionMetadataReplicasLengthMetadataResponseBytes = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x05, 0x6c, 0x6f, 0x67, 0x73, 0x31, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
var invalidPartitionMetadataReplicaMetadataResponseBytes = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x05, 0x6c, 0x6f, 0x67, 0x73, 0x31, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00}
var invalidPartitionMetadataIsrLengthMetadataResponseBytes = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x05, 0x6c, 0x6f, 0x67, 0x73, 0x31, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
var invalidPartitionMetadataIsrMetadataResponseBytes = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x05, 0x6c, 0x6f, 0x67, 0x73, 0x31, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00}

func TestTopicMetadataRequest(t *testing.T) {
	emptyMetadataRequest := new(TopicMetadataRequest)
	testRequest(t, emptyMetadataRequest, emptyMetadataRequestBytes)

	asdMetadataRequest := NewTopicMetadataRequest([]string{"asd"})
	testRequest(t, asdMetadataRequest, asdMetadataRequestBytes)

	multipleTopicsMetadataRequest := NewTopicMetadataRequest([]string{"asd", "zxc", "qwe"})
	testRequest(t, multipleTopicsMetadataRequest, multipleTopicsMetadataRequestBytes)
}

func TestTopicMetadataResponse(t *testing.T) {
	emptyMetadataResponse := new(TopicMetadataResponse)
	decode(t, emptyMetadataResponse, emptyMetadataResponseBytes)
	assertFatal(t, len(emptyMetadataResponse.Brokers), 0)
	assertFatal(t, len(emptyMetadataResponse.TopicMetadata), 0)

	brokerMetadataResponse := new(TopicMetadataResponse)
	decode(t, brokerMetadataResponse, brokerMetadataResponseBytes)
	assertFatal(t, len(brokerMetadataResponse.Brokers), 1)
	broker := brokerMetadataResponse.Brokers[0]
	assert(t, broker.NodeId, int32(0))
	assert(t, broker.Host, "localhost")
	assert(t, broker.Port, int32(9092))
	assertFatal(t, len(brokerMetadataResponse.TopicMetadata), 0)

	topicMetadataResponse := new(TopicMetadataResponse)
	decode(t, topicMetadataResponse, topicMetadataResponseBytes)
	assertFatal(t, len(topicMetadataResponse.Brokers), 0)
	assertFatal(t, len(topicMetadataResponse.TopicMetadata), 1)
	meta := topicMetadataResponse.TopicMetadata[0]
	assert(t, meta.TopicName, "logs1")
	assert(t, meta.Error, NoError)
	assertFatal(t, len(meta.PartitionMetadata), 2)
	partition0 := meta.PartitionMetadata[1]
	assert(t, partition0.PartitionId, int32(0))
	assert(t, partition0.Error, NoError)
	assert(t, partition0.Isr, []int32{0})
	assert(t, partition0.Leader, int32(0))
	assert(t, partition0.Replicas, []int32{0})

	partition1 := meta.PartitionMetadata[0]
	assert(t, partition1.PartitionId, int32(1))
	assert(t, partition1.Error, NoError)
	assert(t, partition1.Isr, []int32{0})
	assert(t, partition1.Leader, int32(0))
	assert(t, partition1.Replicas, []int32{0})

	decodeErr(t, new(TopicMetadataResponse), invalidBrokersLengthMetadataResponseBytes, NewDecodingError(EOF, reason_InvalidBrokersLength))
	decodeErr(t, new(TopicMetadataResponse), invalidMetadataLengthMetadataResponseBytes, NewDecodingError(EOF, reason_InvalidMetadataLength))
	decodeErr(t, new(TopicMetadataResponse), invalidBrokerNodeIdMetadataResponseBytes, NewDecodingError(EOF, reason_InvalidBrokerNodeId))
	decodeErr(t, new(TopicMetadataResponse), invalidBrokerHostMetadataResponseBytes, NewDecodingError(EOF, reason_InvalidBrokerHost))
	decodeErr(t, new(TopicMetadataResponse), invalidBrokerPortMetadataResponseBytes, NewDecodingError(EOF, reason_InvalidBrokerPort))
	decodeErr(t, new(TopicMetadataResponse), invalidTopicMetadataErrorCodeMetadataResponseBytes, NewDecodingError(EOF, reason_InvalidTopicMetadataErrorCode))
	decodeErr(t, new(TopicMetadataResponse), invalidTopicMetadataTopicNameMetadataResponseBytes, NewDecodingError(EOF, reason_InvalidTopicMetadataTopicName))
	decodeErr(t, new(TopicMetadataResponse), invalidPartitionMetadataLengthMetadataResponseBytes, NewDecodingError(EOF, reason_InvalidPartitionMetadataLength))
	decodeErr(t, new(TopicMetadataResponse), invalidPartitionMetadataErrorCodeMetadataResponseBytes, NewDecodingError(EOF, reason_InvalidPartitionMetadataErrorCode))
	decodeErr(t, new(TopicMetadataResponse), invalidPartitionMetadataPartitionMetadataResponseBytes, NewDecodingError(EOF, reason_InvalidPartitionMetadataPartition))
	decodeErr(t, new(TopicMetadataResponse), invalidPartitionMetadataLeaderMetadataResponseBytes, NewDecodingError(EOF, reason_InvalidPartitionMetadataLeader))
	decodeErr(t, new(TopicMetadataResponse), invalidPartitionMetadataReplicasLengthMetadataResponseBytes, NewDecodingError(EOF, reason_InvalidPartitionMetadataReplicasLength))
	decodeErr(t, new(TopicMetadataResponse), invalidPartitionMetadataReplicaMetadataResponseBytes, NewDecodingError(EOF, reason_InvalidPartitionMetadataReplica))
	decodeErr(t, new(TopicMetadataResponse), invalidPartitionMetadataIsrLengthMetadataResponseBytes, NewDecodingError(EOF, reason_InvalidPartitionMetadataIsrLength))
	decodeErr(t, new(TopicMetadataResponse), invalidPartitionMetadataIsrMetadataResponseBytes, NewDecodingError(EOF, reason_InvalidPartitionMetadataIsr))
}
