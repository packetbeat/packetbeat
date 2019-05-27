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

// go run mkpdh_defs.go
// MACHINE GENERATED BY THE ABOVE COMMAND; DO NOT EDIT

// +build ignore

package perfmon

/*
#include <pdh.h>
#include <pdhmsg.h>
#cgo LDFLAGS: -lpdh
*/
import "C"

type PdhErrno uintptr

// PDH Error Codes
const (
	PDH_CSTATUS_VALID_DATA                     PdhErrno = C.PDH_CSTATUS_VALID_DATA
	PDH_CSTATUS_NEW_DATA                       PdhErrno = C.PDH_CSTATUS_NEW_DATA
	PDH_CSTATUS_NO_MACHINE                     PdhErrno = C.PDH_CSTATUS_NO_MACHINE
	PDH_CSTATUS_NO_INSTANCE                    PdhErrno = C.PDH_CSTATUS_NO_INSTANCE
	PDH_MORE_DATA                              PdhErrno = C.PDH_MORE_DATA
	PDH_CSTATUS_ITEM_NOT_VALIDATED             PdhErrno = C.PDH_CSTATUS_ITEM_NOT_VALIDATED
	PDH_RETRY                                  PdhErrno = C.PDH_RETRY
	PDH_NO_DATA                                PdhErrno = C.PDH_NO_DATA
	PDH_CALC_NEGATIVE_DENOMINATOR              PdhErrno = C.PDH_CALC_NEGATIVE_DENOMINATOR
	PDH_CALC_NEGATIVE_TIMEBASE                 PdhErrno = C.PDH_CALC_NEGATIVE_TIMEBASE
	PDH_CALC_NEGATIVE_VALUE                    PdhErrno = C.PDH_CALC_NEGATIVE_VALUE
	PDH_DIALOG_CANCELLED                       PdhErrno = C.PDH_DIALOG_CANCELLED
	PDH_END_OF_LOG_FILE                        PdhErrno = C.PDH_END_OF_LOG_FILE
	PDH_ASYNC_QUERY_TIMEOUT                    PdhErrno = C.PDH_ASYNC_QUERY_TIMEOUT
	PDH_CANNOT_SET_DEFAULT_REALTIME_DATASOURCE PdhErrno = C.PDH_CANNOT_SET_DEFAULT_REALTIME_DATASOURCE
	PDH_CSTATUS_NO_OBJECT                      PdhErrno = C.PDH_CSTATUS_NO_OBJECT
	PDH_CSTATUS_NO_COUNTER                     PdhErrno = C.PDH_CSTATUS_NO_COUNTER
	PDH_CSTATUS_INVALID_DATA                   PdhErrno = C.PDH_CSTATUS_INVALID_DATA
	PDH_MEMORY_ALLOCATION_FAILURE              PdhErrno = C.PDH_MEMORY_ALLOCATION_FAILURE
	PDH_INVALID_HANDLE                         PdhErrno = C.PDH_INVALID_HANDLE
	PDH_INVALID_ARGUMENT                       PdhErrno = C.PDH_INVALID_ARGUMENT
	PDH_FUNCTION_NOT_FOUND                     PdhErrno = C.PDH_FUNCTION_NOT_FOUND
	PDH_CSTATUS_NO_COUNTERNAME                 PdhErrno = C.PDH_CSTATUS_NO_COUNTERNAME
	PDH_CSTATUS_BAD_COUNTERNAME                PdhErrno = C.PDH_CSTATUS_BAD_COUNTERNAME
	PDH_INVALID_BUFFER                         PdhErrno = C.PDH_INVALID_BUFFER
	PDH_INSUFFICIENT_BUFFER                    PdhErrno = C.PDH_INSUFFICIENT_BUFFER
	PDH_CANNOT_CONNECT_MACHINE                 PdhErrno = C.PDH_CANNOT_CONNECT_MACHINE
	PDH_INVALID_PATH                           PdhErrno = C.PDH_INVALID_PATH
	PDH_INVALID_INSTANCE                       PdhErrno = C.PDH_INVALID_INSTANCE
	PDH_INVALID_DATA                           PdhErrno = C.PDH_INVALID_DATA
	PDH_NO_DIALOG_DATA                         PdhErrno = C.PDH_NO_DIALOG_DATA
	PDH_CANNOT_READ_NAME_STRINGS               PdhErrno = C.PDH_CANNOT_READ_NAME_STRINGS
	PDH_LOG_FILE_CREATE_ERROR                  PdhErrno = C.PDH_LOG_FILE_CREATE_ERROR
	PDH_LOG_FILE_OPEN_ERROR                    PdhErrno = C.PDH_LOG_FILE_OPEN_ERROR
	PDH_LOG_TYPE_NOT_FOUND                     PdhErrno = C.PDH_LOG_TYPE_NOT_FOUND
	PDH_NO_MORE_DATA                           PdhErrno = C.PDH_NO_MORE_DATA
	PDH_ENTRY_NOT_IN_LOG_FILE                  PdhErrno = C.PDH_ENTRY_NOT_IN_LOG_FILE
	PDH_DATA_SOURCE_IS_LOG_FILE                PdhErrno = C.PDH_DATA_SOURCE_IS_LOG_FILE
	PDH_DATA_SOURCE_IS_REAL_TIME               PdhErrno = C.PDH_DATA_SOURCE_IS_REAL_TIME
	PDH_UNABLE_READ_LOG_HEADER                 PdhErrno = C.PDH_UNABLE_READ_LOG_HEADER
	PDH_FILE_NOT_FOUND                         PdhErrno = C.PDH_FILE_NOT_FOUND
	PDH_FILE_ALREADY_EXISTS                    PdhErrno = C.PDH_FILE_ALREADY_EXISTS
	PDH_NOT_IMPLEMENTED                        PdhErrno = C.PDH_NOT_IMPLEMENTED
	PDH_STRING_NOT_FOUND                       PdhErrno = C.PDH_STRING_NOT_FOUND
	PDH_UNABLE_MAP_NAME_FILES                  PdhErrno = C.PDH_UNABLE_MAP_NAME_FILES
	PDH_UNKNOWN_LOG_FORMAT                     PdhErrno = C.PDH_UNKNOWN_LOG_FORMAT
	PDH_UNKNOWN_LOGSVC_COMMAND                 PdhErrno = C.PDH_UNKNOWN_LOGSVC_COMMAND
	PDH_LOGSVC_QUERY_NOT_FOUND                 PdhErrno = C.PDH_LOGSVC_QUERY_NOT_FOUND
	PDH_LOGSVC_NOT_OPENED                      PdhErrno = C.PDH_LOGSVC_NOT_OPENED
	PDH_WBEM_ERROR                             PdhErrno = C.PDH_WBEM_ERROR
	PDH_ACCESS_DENIED                          PdhErrno = C.PDH_ACCESS_DENIED
	PDH_LOG_FILE_TOO_SMALL                     PdhErrno = C.PDH_LOG_FILE_TOO_SMALL
	PDH_INVALID_DATASOURCE                     PdhErrno = C.PDH_INVALID_DATASOURCE
	PDH_INVALID_SQLDB                          PdhErrno = C.PDH_INVALID_SQLDB
	PDH_NO_COUNTERS                            PdhErrno = C.PDH_NO_COUNTERS
	PDH_SQL_ALLOC_FAILED                       PdhErrno = C.PDH_SQL_ALLOC_FAILED
	PDH_SQL_ALLOCCON_FAILED                    PdhErrno = C.PDH_SQL_ALLOCCON_FAILED
	PDH_SQL_EXEC_DIRECT_FAILED                 PdhErrno = C.PDH_SQL_EXEC_DIRECT_FAILED
	PDH_SQL_FETCH_FAILED                       PdhErrno = C.PDH_SQL_FETCH_FAILED
	PDH_SQL_ROWCOUNT_FAILED                    PdhErrno = C.PDH_SQL_ROWCOUNT_FAILED
	PDH_SQL_MORE_RESULTS_FAILED                PdhErrno = C.PDH_SQL_MORE_RESULTS_FAILED
	PDH_SQL_CONNECT_FAILED                     PdhErrno = C.PDH_SQL_CONNECT_FAILED
	PDH_SQL_BIND_FAILED                        PdhErrno = C.PDH_SQL_BIND_FAILED
	PDH_CANNOT_CONNECT_WMI_SERVER              PdhErrno = C.PDH_CANNOT_CONNECT_WMI_SERVER
	PDH_PLA_COLLECTION_ALREADY_RUNNING         PdhErrno = C.PDH_PLA_COLLECTION_ALREADY_RUNNING
	PDH_PLA_ERROR_SCHEDULE_OVERLAP             PdhErrno = C.PDH_PLA_ERROR_SCHEDULE_OVERLAP
	PDH_PLA_COLLECTION_NOT_FOUND               PdhErrno = C.PDH_PLA_COLLECTION_NOT_FOUND
	PDH_PLA_ERROR_SCHEDULE_ELAPSED             PdhErrno = C.PDH_PLA_ERROR_SCHEDULE_ELAPSED
	PDH_PLA_ERROR_NOSTART                      PdhErrno = C.PDH_PLA_ERROR_NOSTART
	PDH_PLA_ERROR_ALREADY_EXISTS               PdhErrno = C.PDH_PLA_ERROR_ALREADY_EXISTS
	PDH_PLA_ERROR_TYPE_MISMATCH                PdhErrno = C.PDH_PLA_ERROR_TYPE_MISMATCH
	PDH_PLA_ERROR_FILEPATH                     PdhErrno = C.PDH_PLA_ERROR_FILEPATH
	PDH_PLA_SERVICE_ERROR                      PdhErrno = C.PDH_PLA_SERVICE_ERROR
	PDH_PLA_VALIDATION_ERROR                   PdhErrno = C.PDH_PLA_VALIDATION_ERROR
	PDH_PLA_VALIDATION_WARNING                 PdhErrno = C.PDH_PLA_VALIDATION_WARNING
	PDH_PLA_ERROR_NAME_TOO_LONG                PdhErrno = C.PDH_PLA_ERROR_NAME_TOO_LONG
	PDH_INVALID_SQL_LOG_FORMAT                 PdhErrno = C.PDH_INVALID_SQL_LOG_FORMAT
	PDH_COUNTER_ALREADY_IN_QUERY               PdhErrno = C.PDH_COUNTER_ALREADY_IN_QUERY
	PDH_BINARY_LOG_CORRUPT                     PdhErrno = C.PDH_BINARY_LOG_CORRUPT
	PDH_LOG_SAMPLE_TOO_SMALL                   PdhErrno = C.PDH_LOG_SAMPLE_TOO_SMALL
	PDH_OS_LATER_VERSION                       PdhErrno = C.PDH_OS_LATER_VERSION
	PDH_OS_EARLIER_VERSION                     PdhErrno = C.PDH_OS_EARLIER_VERSION
	PDH_INCORRECT_APPEND_TIME                  PdhErrno = C.PDH_INCORRECT_APPEND_TIME
	PDH_UNMATCHED_APPEND_COUNTER               PdhErrno = C.PDH_UNMATCHED_APPEND_COUNTER
	PDH_SQL_ALTER_DETAIL_FAILED                PdhErrno = C.PDH_SQL_ALTER_DETAIL_FAILED
	PDH_QUERY_PERF_DATA_TIMEOUT                PdhErrno = C.PDH_QUERY_PERF_DATA_TIMEOUT
)

var pdhErrors = map[PdhErrno]struct{}{
	PDH_CSTATUS_VALID_DATA:                     struct{}{},
	PDH_CSTATUS_NEW_DATA:                       struct{}{},
	PDH_CSTATUS_NO_MACHINE:                     struct{}{},
	PDH_CSTATUS_NO_INSTANCE:                    struct{}{},
	PDH_MORE_DATA:                              struct{}{},
	PDH_CSTATUS_ITEM_NOT_VALIDATED:             struct{}{},
	PDH_RETRY:                                  struct{}{},
	PDH_NO_DATA:                                struct{}{},
	PDH_CALC_NEGATIVE_DENOMINATOR:              struct{}{},
	PDH_CALC_NEGATIVE_TIMEBASE:                 struct{}{},
	PDH_CALC_NEGATIVE_VALUE:                    struct{}{},
	PDH_DIALOG_CANCELLED:                       struct{}{},
	PDH_END_OF_LOG_FILE:                        struct{}{},
	PDH_ASYNC_QUERY_TIMEOUT:                    struct{}{},
	PDH_CANNOT_SET_DEFAULT_REALTIME_DATASOURCE: struct{}{},
	PDH_CSTATUS_NO_OBJECT:                      struct{}{},
	PDH_CSTATUS_NO_COUNTER:                     struct{}{},
	PDH_CSTATUS_INVALID_DATA:                   struct{}{},
	PDH_MEMORY_ALLOCATION_FAILURE:              struct{}{},
	PDH_INVALID_HANDLE:                         struct{}{},
	PDH_INVALID_ARGUMENT:                       struct{}{},
	PDH_FUNCTION_NOT_FOUND:                     struct{}{},
	PDH_CSTATUS_NO_COUNTERNAME:                 struct{}{},
	PDH_CSTATUS_BAD_COUNTERNAME:                struct{}{},
	PDH_INVALID_BUFFER:                         struct{}{},
	PDH_INSUFFICIENT_BUFFER:                    struct{}{},
	PDH_CANNOT_CONNECT_MACHINE:                 struct{}{},
	PDH_INVALID_PATH:                           struct{}{},
	PDH_INVALID_INSTANCE:                       struct{}{},
	PDH_INVALID_DATA:                           struct{}{},
	PDH_NO_DIALOG_DATA:                         struct{}{},
	PDH_CANNOT_READ_NAME_STRINGS:               struct{}{},
	PDH_LOG_FILE_CREATE_ERROR:                  struct{}{},
	PDH_LOG_FILE_OPEN_ERROR:                    struct{}{},
	PDH_LOG_TYPE_NOT_FOUND:                     struct{}{},
	PDH_NO_MORE_DATA:                           struct{}{},
	PDH_ENTRY_NOT_IN_LOG_FILE:                  struct{}{},
	PDH_DATA_SOURCE_IS_LOG_FILE:                struct{}{},
	PDH_DATA_SOURCE_IS_REAL_TIME:               struct{}{},
	PDH_UNABLE_READ_LOG_HEADER:                 struct{}{},
	PDH_FILE_NOT_FOUND:                         struct{}{},
	PDH_FILE_ALREADY_EXISTS:                    struct{}{},
	PDH_NOT_IMPLEMENTED:                        struct{}{},
	PDH_STRING_NOT_FOUND:                       struct{}{},
	PDH_UNABLE_MAP_NAME_FILES:                  struct{}{},
	PDH_UNKNOWN_LOG_FORMAT:                     struct{}{},
	PDH_UNKNOWN_LOGSVC_COMMAND:                 struct{}{},
	PDH_LOGSVC_QUERY_NOT_FOUND:                 struct{}{},
	PDH_LOGSVC_NOT_OPENED:                      struct{}{},
	PDH_WBEM_ERROR:                             struct{}{},
	PDH_ACCESS_DENIED:                          struct{}{},
	PDH_LOG_FILE_TOO_SMALL:                     struct{}{},
	PDH_INVALID_DATASOURCE:                     struct{}{},
	PDH_INVALID_SQLDB:                          struct{}{},
	PDH_NO_COUNTERS:                            struct{}{},
	PDH_SQL_ALLOC_FAILED:                       struct{}{},
	PDH_SQL_ALLOCCON_FAILED:                    struct{}{},
	PDH_SQL_EXEC_DIRECT_FAILED:                 struct{}{},
	PDH_SQL_FETCH_FAILED:                       struct{}{},
	PDH_SQL_ROWCOUNT_FAILED:                    struct{}{},
	PDH_SQL_MORE_RESULTS_FAILED:                struct{}{},
	PDH_SQL_CONNECT_FAILED:                     struct{}{},
	PDH_SQL_BIND_FAILED:                        struct{}{},
	PDH_CANNOT_CONNECT_WMI_SERVER:              struct{}{},
	PDH_PLA_COLLECTION_ALREADY_RUNNING:         struct{}{},
	PDH_PLA_ERROR_SCHEDULE_OVERLAP:             struct{}{},
	PDH_PLA_COLLECTION_NOT_FOUND:               struct{}{},
	PDH_PLA_ERROR_SCHEDULE_ELAPSED:             struct{}{},
	PDH_PLA_ERROR_NOSTART:                      struct{}{},
	PDH_PLA_ERROR_ALREADY_EXISTS:               struct{}{},
	PDH_PLA_ERROR_TYPE_MISMATCH:                struct{}{},
	PDH_PLA_ERROR_FILEPATH:                     struct{}{},
	PDH_PLA_SERVICE_ERROR:                      struct{}{},
	PDH_PLA_VALIDATION_ERROR:                   struct{}{},
	PDH_PLA_VALIDATION_WARNING:                 struct{}{},
	PDH_PLA_ERROR_NAME_TOO_LONG:                struct{}{},
	PDH_INVALID_SQL_LOG_FORMAT:                 struct{}{},
	PDH_COUNTER_ALREADY_IN_QUERY:               struct{}{},
	PDH_BINARY_LOG_CORRUPT:                     struct{}{},
	PDH_LOG_SAMPLE_TOO_SMALL:                   struct{}{},
	PDH_OS_LATER_VERSION:                       struct{}{},
	PDH_OS_EARLIER_VERSION:                     struct{}{},
	PDH_INCORRECT_APPEND_TIME:                  struct{}{},
	PDH_UNMATCHED_APPEND_COUNTER:               struct{}{},
	PDH_SQL_ALTER_DETAIL_FAILED:                struct{}{},
	PDH_QUERY_PERF_DATA_TIMEOUT:                struct{}{},
}

type PdhCounterFormat uint32

// PDH Counter Formats
const (
	// PdhFmtDouble returns data as a double-precision floating point real.
	PdhFmtDouble PdhCounterFormat = C.PDH_FMT_DOUBLE
	// PdhFmtLarge returns data as a 64-bit integer.
	PdhFmtLarge PdhCounterFormat = C.PDH_FMT_LARGE
	// PdhFmtLong returns data as a long integer.
	PdhFmtLong PdhCounterFormat = C.PDH_FMT_LONG

	// Use bitwise operators to combine these values with the counter type to scale the value.

	// Do not apply the counter's default scaling factor.
	PdhFmtNoScale PdhCounterFormat = C.PDH_FMT_NOSCALE
	// Counter values greater than 100 (for example, counter values measuring
	// the processor load on multiprocessor computers) will not be reset to 100.
	// The default behavior is that counter values are capped at a value of 100.
	PdhFmtNoCap100 PdhCounterFormat = C.PDH_FMT_NOCAP100
	// Multiply the actual value by 1,000.
	PdhFmtMultiply1000 PdhCounterFormat = C.PDH_FMT_1000
)

// PdhRawCounter is the structure that receives the raw counter.
type PdhRawCounter C.PDH_RAW_COUNTER

type PdhFileTime C.FILETIME
