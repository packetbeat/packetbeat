// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs defs_kernel_types.go

package rule

type filter uint32

const (
	userFilter    filter = 0x0
	taskFilter    filter = 0x1
	entryFilter   filter = 0x2
	watchFilter   filter = 0x3
	exitFilter    filter = 0x4
	typeFilter    filter = 0x5
	excludeFilter        = typeFilter

	prependFilter filter = 0x10
)

type action uint32

const (
	neverAction    action = 0x0
	possibleAction action = 0x1
	alwaysAction   action = 0x2
)

type field uint32

const (
	auidField               field = 0x9
	archField               field = 0xb
	arg0Field               field = 0xc8
	arg1Field               field = 0xc9
	arg2Field               field = 0xca
	arg3Field               field = 0xcb
	devMajorField           field = 0x64
	devMinorField           field = 0x65
	dirField                field = 0x6b
	egidField               field = 0x6
	euidField               field = 0x2
	exeField                field = 0x70
	exitField               field = 0x67
	fsgidField              field = 0x8
	fsuidField              field = 0x4
	filetypeField           field = 0x6c
	gidField                field = 0x5
	inodeField              field = 0x66
	keyField                field = 0xd2
	msgTypeField            field = 0xc
	objectGIDField          field = 0x6e
	objectLevelHighField    field = 0x17
	objectLevelLowField     field = 0x16
	objectRoleField         field = 0x14
	objectTypeField         field = 0x15
	objectUIDField          field = 0x6d
	objectUserField         field = 0x13
	pathField               field = 0x69
	pidField                field = 0x0
	ppidField               field = 0x12
	permField               field = 0x6a
	persField               field = 0xa
	sgidField               field = 0x7
	suidField               field = 0x3
	subjectClearanceField   field = 0x11
	subjectRoleField        field = 0xe
	subjectSensitivityField field = 0x10
	subjectTypeField        field = 0xf
	subjectUserField        field = 0xd
	successField            field = 0x68
	uidField                field = 0x1

	fieldCompare field = 0x6f
)

type operator uint32

const (
	bitMaskOperator            operator = 0x8000000
	lessThanOperator           operator = 0x10000000
	greaterThanOperator        operator = 0x20000000
	notEqualOperator           operator = 0x30000000
	equalOperator              operator = 0x40000000
	bitTestOperator            operator = 0x48000000
	lessThanOrEqualOperator    operator = 0x50000000
	greaterThanOrEqualOperator operator = 0x60000000
)

type comparison uint32

const (
	_AUDIT_COMPARE_UID_TO_OBJ_UID   comparison = 0x1
	_AUDIT_COMPARE_GID_TO_OBJ_GID   comparison = 0x2
	_AUDIT_COMPARE_EUID_TO_OBJ_UID  comparison = 0x3
	_AUDIT_COMPARE_EGID_TO_OBJ_GID  comparison = 0x4
	_AUDIT_COMPARE_AUID_TO_OBJ_UID  comparison = 0x5
	_AUDIT_COMPARE_SUID_TO_OBJ_UID  comparison = 0x6
	_AUDIT_COMPARE_SGID_TO_OBJ_GID  comparison = 0x7
	_AUDIT_COMPARE_FSUID_TO_OBJ_UID comparison = 0x8
	_AUDIT_COMPARE_FSGID_TO_OBJ_GID comparison = 0x9

	_AUDIT_COMPARE_UID_TO_AUID  comparison = 0xa
	_AUDIT_COMPARE_UID_TO_EUID  comparison = 0xb
	_AUDIT_COMPARE_UID_TO_FSUID comparison = 0xc
	_AUDIT_COMPARE_UID_TO_SUID  comparison = 0xd

	_AUDIT_COMPARE_AUID_TO_FSUID comparison = 0xe
	_AUDIT_COMPARE_AUID_TO_SUID  comparison = 0xf
	_AUDIT_COMPARE_AUID_TO_EUID  comparison = 0x10

	_AUDIT_COMPARE_EUID_TO_SUID  comparison = 0x11
	_AUDIT_COMPARE_EUID_TO_FSUID comparison = 0x12

	_AUDIT_COMPARE_SUID_TO_FSUID comparison = 0x13

	_AUDIT_COMPARE_GID_TO_EGID  comparison = 0x14
	_AUDIT_COMPARE_GID_TO_FSGID comparison = 0x15
	_AUDIT_COMPARE_GID_TO_SGID  comparison = 0x16

	_AUDIT_COMPARE_EGID_TO_FSGID comparison = 0x17
	_AUDIT_COMPARE_EGID_TO_SGID  comparison = 0x18
	_AUDIT_COMPARE_SGID_TO_FSGID comparison = 0x19
)

type permission uint32

const (
	execPerm  permission = 0x1
	writePerm permission = 0x2
	readPerm  permission = 0x4
	attrPerm  permission = 0x8
)

type filetype uint32

const (
	fileFiletype      filetype = 0x8000
	socketFiletype    filetype = 0xc000
	linkFiletype      filetype = 0xa000
	blockFiletype     filetype = 0x6000
	dirFiletype       filetype = 0x4000
	characterFiletype filetype = 0x2000
	fifoFiletype      filetype = 0x1000
)
