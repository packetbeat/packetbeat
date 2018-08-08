package replstatus

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// The status of the replica set from the point of view of the server that processed the command.
type MongoReplStatus struct {
	Ok                      bool      `bson:"ok"`
	Set                     string    `bson:"set"`
	Date                    time.Time `bson:"date"`
	Members                 []Member  `bson:"members"`
	MyState                 int       `bson:"myState"`
	Term                    int       `bson:"term"`
	HeartbeatIntervalMillis int       `bson:"heartbeatIntervalMillis"`
	OpTimes                 struct {
		LastCommitted OpTime `bson:"lastCommittedOpTime"`
		Applied       OpTime `bson:"appliedOpTime"`
		Durable       OpTime `bson:"durableOpTime"`
	} `bson:"optimes"`
}

// Provides information about a member in the replica set.
type Member struct {
	Health        bool      `bson:"health"`
	Name          string    `bson:"name"`
	State         int       `bson:"state"`
	StateStr      string    `bson:"stateStr"`
	Uptime        int       `bson:"uptime"`
	OpTime        OpTime    `bson:"optime"`
	OpTimeDate    time.Time `bson:"optimeDate"`
	ElectionTime  int64     `bson:"electionTime"`
	ElectionDate  time.Time `bson:"electaionDate"`
	ConfigVersion int       `bson:"configVersion"`
	Self          bool      `bson:"self"`
}

// Information regarding the operation from the operation log
type OpTime struct {
	Ts int64 `bson:"ts"` // The timestamp of the last operation applied to this member of the replica set
	T  int   `bson:"t"`  // The term in which the last applied operation was originally generated on the primary.
}

// State of a member in replica set
type MemberState int

const (
	STARTUP    MemberState = 0  // STARTUP state
	PRIMARY    MemberState = 1  // PRIMARY state
	SECONDARY  MemberState = 2  // SECONDARY state
	RECOVERING MemberState = 3  // RECOVERING state
	STARTUP2   MemberState = 5  // STARTUP2 state
	UNKNOWN    MemberState = 6  // UNKNOWN state
	ARBITER    MemberState = 7  // ARBITER state
	DOWN       MemberState = 8  // DOWN state
	ROLLBACK   MemberState = 9  // ROLLBACK state
	REMOVED    MemberState = 10 // REMOVED state
)

func (optime *OpTime) getTimeStamp() int64 {
	return optime.Ts >> 32
}

func getReplicationStatus(mongoSession *mgo.Session) (*MongoReplStatus, error) {
	db := mongoSession.DB("admin")

	var replStatus MongoReplStatus
	if err := db.Run(bson.M{"replSetGetStatus": 1}, &replStatus); err != nil {
		return nil, err
	}

	return &replStatus, nil
}

func findUnhealthyHosts(members []Member) []string {
	var hosts []string

	for _, member := range members {
		if member.Health == false {
			hosts = append(hosts, member.Name)
		}
	}

	return hosts
}

func findHostsByState(members []Member, state MemberState) []string {
	var hosts []string

	for _, member := range members {
		memberState := MemberState(member.State)
		if memberState == state {
			hosts = append(hosts, member.Name)
		}
	}

	return hosts
}

func findMaxLag(members []Member) int64 {
	var minOptime, primaryOptime int64 = 1<<63 - 1, 0

	for _, member := range members {
		memberState := MemberState(member.State)
		if memberState == SECONDARY {
			if minOptime > member.OpTime.Ts {
				minOptime = member.OpTime.Ts
			}
		} else if memberState == PRIMARY {
			primaryOptime = member.OpTime.Ts
		}
	}

	return primaryOptime - minOptime
}

func findMinLag(members []Member) int64 {
	var maxOptime, primaryOptime int64 = 0, 0

	for _, member := range members {
		memberState := MemberState(member.State)
		if memberState == SECONDARY {
			if member.OpTime.Ts > maxOptime {
				maxOptime = member.OpTime.Ts
			}
		} else if memberState == PRIMARY {
			primaryOptime = member.OpTime.Ts
		}
	}

	return primaryOptime - maxOptime
}
