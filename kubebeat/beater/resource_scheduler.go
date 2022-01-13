package beater

type ResourceScheduler interface {
	ScheduleResources(resourcesMap map[string][]interface{}, resourceFunc func(interface{}))
}

type SynchronousScheduler struct {
}

func NewSynchronousScheduler() ResourceScheduler {
	return &SynchronousScheduler{}
}

func (s *SynchronousScheduler) ScheduleResources(resourcesMap map[string][]interface{}, resourceFunc func(interface{})) {
	for _, resources := range resourcesMap {
		for _, r := range resources {
			resourceFunc(r)
		}
	}
}
