package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/madneal/gshark/global"
	"github.com/madneal/gshark/model"
	"github.com/madneal/gshark/model/request"
	"github.com/madneal/gshark/model/response"
	"github.com/madneal/gshark/service"
	"go.uber.org/zap"
)

func CreateSearchResult(c *gin.Context) {
	var searchResult model.SearchResult
	_ = c.ShouldBindJSON(&searchResult)
	if err := service.CreateSearchResult(searchResult); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func DeleteSearchResult(c *gin.Context) {
	var searchResult model.SearchResult
	_ = c.ShouldBindJSON(&searchResult)
	if err := service.DeleteSearchResult(searchResult); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func DeleteSearchResultByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteSearchResultByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func UpdateSearchResultByIds(c *gin.Context) {
	var batchUpdateReq request.BatchUpdateReq
	_ = c.ShouldBindJSON(&batchUpdateReq)
	if err := service.UpdateSearchResultByIds(batchUpdateReq); err != nil {
		global.GVA_LOG.Error("批量更新状态失败！", zap.Any("err", err))
		response.FailWithMessage("批量更新状态失败", c)
	} else {
		response.OkWithMessage("批量更新状态成功", c)
	}
}

func UpdateSearchResult(c *gin.Context) {
	var updateReq request.UpdateReq
	_ = c.ShouldBindJSON(&updateReq)
	if err := service.UpdateSearchResult(updateReq); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func FindSearchResult(c *gin.Context) {
	var searchResult model.SearchResult
	_ = c.ShouldBindQuery(&searchResult)
	if err, researchResult := service.GetSearchResult(searchResult.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"researchResult": researchResult}, c)
	}
}

func GetSearchResultList(c *gin.Context) {
	var pageInfo request.SearchResultSearch
	_ = c.ShouldBindQuery(&pageInfo)
	//if err != nil {
	//	global.GVA_LOG.Error("GetSearchResultList bind query error", zap.Any("err", err))
	//}
	if err, list, total := service.GetSearchResultInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
