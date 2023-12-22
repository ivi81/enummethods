package withstringarray

//go:generate enummethods -type ObjectType
type ObjectType int

var (
//objectType = [...]string{"case", "Case", "case_artifact", "Artifact", "case_task", "Task", "case_task_log", "Log", "alert", "Alert", "organisation", "Organisation"}
//hueobjectType = [...]string{"case1", "Case1", "case_artifact1", "Artifact1", "case_task1", "Task1", "case_task_log1", "Log1", "alert1", "Alert1", "organisation1", "Organisation1"}
)

const (
	HIVECASE = ObjectType(iota)
	HIVECASEAPI1
	HIVEARTIFACT
	HIVEARTIFACTAPI1
	HIVETASK
	HIVETASKAPI1
	HIVETASKLOG
	HIVETASKLOGAPI1
	HIVEALERT
	HIVEALERTAPI1
	HIVEORG
	HIVEORGAPI1
)
