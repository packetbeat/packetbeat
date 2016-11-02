package nfs

var nfsOpnum3 = [...]string{
	"NULL",
	"GETATTR",
	"SETATTR",
	"LOOKUP",
	"ACCESS",
	"READLINK",
	"READ",
	"WRITE",
	"CREATE",
	"MKDIR",
	"SYM_LINK",
	"MKNODE",
	"REMOVE",
	"RMDIR",
	"RENAME",
	"LINK",
	"READDIR",
	"READDIRPLUS",
	"FSSTAT",
	"FSINFO",
	"PATHINFO",
	"COMMIT",
}

func (nfs *NFS) getV3Opcode(proc int) string {
	if proc < len(nfsOpnum3) {
		return nfsOpnum3[proc]
	} else {
		return "ILLEGAL"
	}
}
