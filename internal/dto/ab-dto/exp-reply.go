package abDto

type ExperimentsReply struct {
	ExperimentsReply []ExperimentReply `json:"experiments_reply"`
}

type ExperimentReply struct {
	ExperimentName string `json:"experiment_name"`
	GroupName      string `json:"group_name"`
}
