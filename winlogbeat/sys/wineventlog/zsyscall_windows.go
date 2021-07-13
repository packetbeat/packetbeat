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

// Code generated by 'go generate'; DO NOT EDIT.

package wineventlog

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modwevtapi = windows.NewLazySystemDLL("wevtapi.dll")

	procEvtClearLog                     = modwevtapi.NewProc("EvtClearLog")
	procEvtClose                        = modwevtapi.NewProc("EvtClose")
	procEvtCreateBookmark               = modwevtapi.NewProc("EvtCreateBookmark")
	procEvtCreateRenderContext          = modwevtapi.NewProc("EvtCreateRenderContext")
	procEvtFormatMessage                = modwevtapi.NewProc("EvtFormatMessage")
	procEvtGetEventMetadataProperty     = modwevtapi.NewProc("EvtGetEventMetadataProperty")
	procEvtGetObjectArrayProperty       = modwevtapi.NewProc("EvtGetObjectArrayProperty")
	procEvtGetObjectArraySize           = modwevtapi.NewProc("EvtGetObjectArraySize")
	procEvtGetPublisherMetadataProperty = modwevtapi.NewProc("EvtGetPublisherMetadataProperty")
	procEvtNext                         = modwevtapi.NewProc("EvtNext")
	procEvtNextChannelPath              = modwevtapi.NewProc("EvtNextChannelPath")
	procEvtNextEventMetadata            = modwevtapi.NewProc("EvtNextEventMetadata")
	procEvtNextPublisherId              = modwevtapi.NewProc("EvtNextPublisherId")
	procEvtOpenChannelEnum              = modwevtapi.NewProc("EvtOpenChannelEnum")
	procEvtOpenEventMetadataEnum        = modwevtapi.NewProc("EvtOpenEventMetadataEnum")
	procEvtOpenLog                      = modwevtapi.NewProc("EvtOpenLog")
	procEvtOpenPublisherEnum            = modwevtapi.NewProc("EvtOpenPublisherEnum")
	procEvtOpenPublisherMetadata        = modwevtapi.NewProc("EvtOpenPublisherMetadata")
	procEvtQuery                        = modwevtapi.NewProc("EvtQuery")
	procEvtRender                       = modwevtapi.NewProc("EvtRender")
	procEvtSeek                         = modwevtapi.NewProc("EvtSeek")
	procEvtSubscribe                    = modwevtapi.NewProc("EvtSubscribe")
	procEvtUpdateBookmark               = modwevtapi.NewProc("EvtUpdateBookmark")
)

func _EvtClearLog(session EvtHandle, channelPath *uint16, targetFilePath *uint16, flags uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procEvtClearLog.Addr(), 4, uintptr(session), uintptr(unsafe.Pointer(channelPath)), uintptr(unsafe.Pointer(targetFilePath)), uintptr(flags), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func _EvtClose(object EvtHandle) (err error) {
	r1, _, e1 := syscall.Syscall(procEvtClose.Addr(), 1, uintptr(object), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func _EvtCreateBookmark(bookmarkXML *uint16) (handle EvtHandle, err error) {
	r0, _, e1 := syscall.Syscall(procEvtCreateBookmark.Addr(), 1, uintptr(unsafe.Pointer(bookmarkXML)), 0, 0)
	handle = EvtHandle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

func _EvtCreateRenderContext(ValuePathsCount uint32, valuePaths uintptr, flags EvtRenderContextFlag) (handle EvtHandle, err error) {
	r0, _, e1 := syscall.Syscall(procEvtCreateRenderContext.Addr(), 3, uintptr(ValuePathsCount), uintptr(valuePaths), uintptr(flags))
	handle = EvtHandle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

func _EvtFormatMessage(publisherMetadata EvtHandle, event EvtHandle, messageID uint32, valueCount uint32, values uintptr, flags EvtFormatMessageFlag, bufferSize uint32, buffer *byte, bufferUsed *uint32) (err error) {
	r1, _, e1 := syscall.Syscall9(procEvtFormatMessage.Addr(), 9, uintptr(publisherMetadata), uintptr(event), uintptr(messageID), uintptr(valueCount), uintptr(values), uintptr(flags), uintptr(bufferSize), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(bufferUsed)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func _EvtGetEventMetadataProperty(eventMetadata EvtHandle, propertyID EvtEventMetadataPropertyID, flags uint32, bufferSize uint32, variant *EvtVariant, bufferUsed *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procEvtGetEventMetadataProperty.Addr(), 6, uintptr(eventMetadata), uintptr(propertyID), uintptr(flags), uintptr(bufferSize), uintptr(unsafe.Pointer(variant)), uintptr(unsafe.Pointer(bufferUsed)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func _EvtGetObjectArrayProperty(objectArray EvtObjectArrayPropertyHandle, propertyID EvtPublisherMetadataPropertyID, arrayIndex uint32, flags uint32, bufferSize uint32, evtVariant *EvtVariant, bufferUsed *uint32) (err error) {
	r1, _, e1 := syscall.Syscall9(procEvtGetObjectArrayProperty.Addr(), 7, uintptr(objectArray), uintptr(propertyID), uintptr(arrayIndex), uintptr(flags), uintptr(bufferSize), uintptr(unsafe.Pointer(evtVariant)), uintptr(unsafe.Pointer(bufferUsed)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func _EvtGetObjectArraySize(objectArray EvtObjectArrayPropertyHandle, arraySize *uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procEvtGetObjectArraySize.Addr(), 2, uintptr(objectArray), uintptr(unsafe.Pointer(arraySize)), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func _EvtGetPublisherMetadataProperty(publisherMetadata EvtHandle, propertyID EvtPublisherMetadataPropertyID, flags uint32, bufferSize uint32, variant *EvtVariant, bufferUsed *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procEvtGetPublisherMetadataProperty.Addr(), 6, uintptr(publisherMetadata), uintptr(propertyID), uintptr(flags), uintptr(bufferSize), uintptr(unsafe.Pointer(variant)), uintptr(unsafe.Pointer(bufferUsed)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func _EvtNext(resultSet EvtHandle, eventArraySize uint32, eventArray *EvtHandle, timeout uint32, flags uint32, numReturned *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procEvtNext.Addr(), 6, uintptr(resultSet), uintptr(eventArraySize), uintptr(unsafe.Pointer(eventArray)), uintptr(timeout), uintptr(flags), uintptr(unsafe.Pointer(numReturned)))
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func _EvtNextChannelPath(channelEnum EvtHandle, channelPathBufferSize uint32, channelPathBuffer *uint16, channelPathBufferUsed *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procEvtNextChannelPath.Addr(), 4, uintptr(channelEnum), uintptr(channelPathBufferSize), uintptr(unsafe.Pointer(channelPathBuffer)), uintptr(unsafe.Pointer(channelPathBufferUsed)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func _EvtNextEventMetadata(enumerator EvtHandle, flags uint32) (handle EvtHandle, err error) {
	r0, _, e1 := syscall.Syscall(procEvtNextEventMetadata.Addr(), 2, uintptr(enumerator), uintptr(flags), 0)
	handle = EvtHandle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

func _EvtNextPublisherId(enumerator EvtHandle, bufferSize uint32, buffer *uint16, bufferUsed *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procEvtNextPublisherId.Addr(), 4, uintptr(enumerator), uintptr(bufferSize), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(bufferUsed)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func _EvtOpenChannelEnum(session EvtHandle, flags uint32) (handle EvtHandle, err error) {
	r0, _, e1 := syscall.Syscall(procEvtOpenChannelEnum.Addr(), 2, uintptr(session), uintptr(flags), 0)
	handle = EvtHandle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

func _EvtOpenEventMetadataEnum(publisherMetadata EvtHandle, flags uint32) (handle EvtHandle, err error) {
	r0, _, e1 := syscall.Syscall(procEvtOpenEventMetadataEnum.Addr(), 2, uintptr(publisherMetadata), uintptr(flags), 0)
	handle = EvtHandle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

func _EvtOpenLog(session EvtHandle, path *uint16, flags uint32) (handle EvtHandle, err error) {
	r0, _, e1 := syscall.Syscall(procEvtOpenLog.Addr(), 3, uintptr(session), uintptr(unsafe.Pointer(path)), uintptr(flags))
	handle = EvtHandle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

func _EvtOpenPublisherEnum(session EvtHandle, flags uint32) (handle EvtHandle, err error) {
	r0, _, e1 := syscall.Syscall(procEvtOpenPublisherEnum.Addr(), 2, uintptr(session), uintptr(flags), 0)
	handle = EvtHandle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

func _EvtOpenPublisherMetadata(session EvtHandle, publisherIdentity *uint16, logFilePath *uint16, locale uint32, flags uint32) (handle EvtHandle, err error) {
	r0, _, e1 := syscall.Syscall6(procEvtOpenPublisherMetadata.Addr(), 5, uintptr(session), uintptr(unsafe.Pointer(publisherIdentity)), uintptr(unsafe.Pointer(logFilePath)), uintptr(locale), uintptr(flags), 0)
	handle = EvtHandle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

func _EvtQuery(session EvtHandle, path *uint16, query *uint16, flags uint32) (handle EvtHandle, err error) {
	r0, _, e1 := syscall.Syscall6(procEvtQuery.Addr(), 4, uintptr(session), uintptr(unsafe.Pointer(path)), uintptr(unsafe.Pointer(query)), uintptr(flags), 0, 0)
	handle = EvtHandle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

func _EvtRender(context EvtHandle, fragment EvtHandle, flags EvtRenderFlag, bufferSize uint32, buffer *byte, bufferUsed *uint32, propertyCount *uint32) (err error) {
	r1, _, e1 := syscall.Syscall9(procEvtRender.Addr(), 7, uintptr(context), uintptr(fragment), uintptr(flags), uintptr(bufferSize), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(bufferUsed)), uintptr(unsafe.Pointer(propertyCount)), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func _EvtSeek(resultSet EvtHandle, position int64, bookmark EvtHandle, timeout uint32, flags uint32) (success bool, err error) {
	var (
		r0 uintptr
		e1 syscall.Errno
	)
	if unsafe.Sizeof(uintptr(0)) == unsafe.Sizeof(uint64(0)) {
		r0, _, e1 = syscall.Syscall6(procEvtSeek.Addr(), 5, uintptr(resultSet), uintptr(position), uintptr(bookmark), uintptr(timeout), uintptr(flags), 0)
	} else {
		var _p0 [2]uintptr = *(*[2]uintptr)(unsafe.Pointer(&position))

		r0, _, e1 = syscall.Syscall6(procEvtSeek.Addr(), 6, uintptr(resultSet), _p0[0], _p0[1], uintptr(bookmark), uintptr(timeout), uintptr(flags))
	}

	success = r0 != 0
	if !success {
		err = errnoErr(e1)
	}
	return
}

func _EvtSubscribe(session EvtHandle, signalEvent uintptr, channelPath *uint16, query *uint16, bookmark EvtHandle, context uintptr, callback syscall.Handle, flags EvtSubscribeFlag) (handle EvtHandle, err error) {
	r0, _, e1 := syscall.Syscall9(procEvtSubscribe.Addr(), 8, uintptr(session), uintptr(signalEvent), uintptr(unsafe.Pointer(channelPath)), uintptr(unsafe.Pointer(query)), uintptr(bookmark), uintptr(context), uintptr(callback), uintptr(flags), 0)
	handle = EvtHandle(r0)
	if handle == 0 {
		err = errnoErr(e1)
	}
	return
}

func _EvtUpdateBookmark(bookmark EvtHandle, event EvtHandle) (err error) {
	r1, _, e1 := syscall.Syscall(procEvtUpdateBookmark.Addr(), 2, uintptr(bookmark), uintptr(event), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}
