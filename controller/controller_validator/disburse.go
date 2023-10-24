package controllervalidator

import (
	"net/http"

	"github.com/iqrahadian/paperid-assesment/common"
	"github.com/iqrahadian/paperid-assesment/common/ctx"
	"github.com/iqrahadian/paperid-assesment/model/param"
)

func PostDirburseValidator(carrier *ctx.Carrier, req *http.Request) (param.DisburseParam, common.Error)
