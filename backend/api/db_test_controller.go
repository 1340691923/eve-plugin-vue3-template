package api

import (
	"eve-plugin-vue3-template/backend/dto"
	"eve-plugin-vue3-template/backend/model"
	"eve-plugin-vue3-template/backend/response"
	"github.com/gin-gonic/gin"
)

type DbTestController struct {
	*BaseController
}

func NewDbTestController(baseController *BaseController) *DbTestController {
	return &DbTestController{BaseController: baseController}
}

func (this *DbTestController) InsertAction(ctx *gin.Context) {
	req := new(dto.DbTestInsertReq)
	err := ctx.Bind(req)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	testModel := &model.TestModel{
		Key:   req.Key,
		Value: req.Value,
	}

	err = testModel.Insert()

	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.OperateSuccess, nil)
}

func (this *DbTestController) DeleteAction(ctx *gin.Context) {

	testModel := new(model.TestModel)

	err := testModel.Clean()

	if err != nil {
		this.Error(ctx, err)
		return
	}

	this.Success(ctx, response.OperateSuccess, nil)
}

func (this *DbTestController) SearchAction(ctx *gin.Context) {

	req := new(dto.DbTestSearchReq)
	err := ctx.BindJSON(req)
	if err != nil {
		this.Error(ctx, err)
		return
	}

	testModel := new(model.TestModel)

	list, err := testModel.List(req.Page, req.Limit)

	if err != nil {
		this.Error(ctx, err)
		return
	}

	count, err := testModel.Count()
	if err != nil {
		this.Error(ctx, err)
		return
	}
	this.Success(ctx, response.SearchSuccess, map[string]interface{}{"list": list, "count": count})
}
