package standalone

import (
	"github.com/actionscore/cli/pkg/age"
	"github.com/actionscore/cli/pkg/rundata"
	"github.com/actionscore/cli/utils"
)

type ListOutput struct {
	AppID       string `csv:"APP ID"`
	ActionsPort int    `csv:"ACTIONS PORT"`
	AppPort     int    `csv:"APP PORT"`
	Command     string `csv:"COMMAND"`
	Age         string `csv:"AGE"`
	Created     string `csv:"CREATED"`
	PID         int
}

func List() ([]ListOutput, error) {
	list := []ListOutput{}

	runtimeData, err := rundata.ReadAllRunData()
	if err != nil {
		return nil, err
	}

	for _, runtimeLine := range *runtimeData {
		// TODO: Call to /metadata and validate the runtime data
		var listRow = ListOutput{
			AppID:       runtimeLine.AppId,
			ActionsPort: runtimeLine.ActionsPort,
			Command:     utils.TruncateString(runtimeLine.Command, 20),
			Created:     runtimeLine.Created.Format("2006-01-02 15:04.05"),
			PID:         runtimeLine.PID,
		}
		if runtimeLine.AppPort > 0 {
			listRow.AppPort = runtimeLine.AppPort
		}
		listRow.Age = age.GetAge(runtimeLine.Created)
		list = append(list, listRow)
	}

	return list, nil
}