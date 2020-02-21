package traffic_cop

import (
	"errors"
	"github.com/tlarsen7572/Golang-Public/ryx/folders"
)

const browseFolderFunc = `BrowseFolder`
const getToolDataFunc = `GetToolData`
const invalidAppFunc = `invalid app function`

func handleAppFunction(call FunctionCall) FunctionResponse {
	switch call.Function {
	case browseFolderFunc:
		return browseFolder(call)
	case getToolDataFunc:
		return getToolData(call)
	default:
		return FunctionResponse{errors.New(invalidAppFunc), nil}
	}
}

func browseFolder(call FunctionCall) FunctionResponse {
	folderPath, ok := call.Parameters[`FolderPath`].(string)
	if !ok {
		return _errorResponse(_stringParamErr(`FolderPath`))
	}
	controller := folders.InitializeFolderController(call.Config.BrowseFolderRoots...)
	contents, err := controller.ReadFolder(folderPath)
	return FunctionResponse{err, contents}
}

func getToolData(call FunctionCall) FunctionResponse {
	return FunctionResponse{
		Err:      nil,
		Response: call.Config.ToolData,
	}
}
