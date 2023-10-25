package disburse

import (
	"github.com/iqrahadian/paperid-assesment/common"
	"github.com/iqrahadian/paperid-assesment/model/param"
)

func GetDisburseProcessor(destinationAccount string) (DisburseProcessor, common.Error) {

	err := common.Error{}

	// many logic can be applied here
	// like depend on third party disburse success rate, switch between one and other
	// or depend on the destination account, some switcher cannot process certain account
	return ExternallProcessor{}, err

}

type DisburseProcessor interface {
	Disburse(param.DisburseParam) common.Error
}

// This External Processor can be extented depend on disburse thirdparty, like Artajasa, Ayoconnect, Etc.
type ExternallProcessor struct {
}

func (i ExternallProcessor) Disburse(disburse param.DisburseParam) common.Error {

	// logic for http request to third party API
	// deliberately keep this logic empty
	return common.Error{}
}
