package disburse

import (
	"github.com/iqrahadian/paperid-assesment/common"
	"github.com/iqrahadian/paperid-assesment/model"
	"github.com/iqrahadian/paperid-assesment/model/param"
)

func GetDisburseProcessor(accountType model.AccountType, accountNumber string) (DisburseProcessor, common.Error) {

	err := common.Error{}

	if accountType == model.Internal {
		return InternalProcessor{}, err
	} else {
		// many logic can be applied here
		// like depend on third party disburse success rate, switch between one and other
		return ExternallProcessor{}, err
	}

	return InternalProcessor{}, err

}

type DisburseProcessor interface {
	Disburse(param.DisburseParam) common.Error
}

type InternalProcessor struct {
}

func (i InternalProcessor) Disburse(disburse param.DisburseParam) common.Error {
	return common.Error{}
}

// This External Processor can be extented depend on disburse thirdparty, like Artajasa, Ayoconnect, Etc.
type ExternallProcessor struct {
}

func (i ExternallProcessor) Disburse(disburse param.DisburseParam) common.Error {
	return common.Error{}
}
