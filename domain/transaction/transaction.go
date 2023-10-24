package transaction

import (
	"github.com/iqrahadian/paperid-assesment/common"
	"github.com/iqrahadian/paperid-assesment/common/ctx"
	"github.com/iqrahadian/paperid-assesment/model/param"
)

func Disburse(carrier *ctx.Carrier, disburseParam param.DisburseParam) common.Error
