package service

import (
	"github.com/madneal/gshark/global"
	"github.com/madneal/gshark/model"
	"github.com/madneal/gshark/model/request"
)

func CreateSearchResult(searchResult model.SearchResult) (err error) {
	err = global.GVA_DB.Create(&searchResult).Error
	return err
}

func DeleteSearchResult(searchResult model.SearchResult) (err error) {
	err = global.GVA_DB.Delete(&searchResult).Error
	return err
}

func DeleteSearchResultByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.SearchResult{}, "id in ?", ids.Ids).Error
	return err
}

func UpdateSearchResultByIds(req request.BatchUpdateReq) (err error) {
	err = global.GVA_DB.Table("search_result").Where("id in ?", req.Ids).
		UpdateColumn("status", req.Status).Error
	return err
}

func UpdateSearchResult(updateReq request.UpdateReq) (err error) {
	err = global.GVA_DB.Table("search_result").Where("repo = ?", updateReq.Repo).
		UpdateColumn("status", updateReq.Status).Error
	return err
}

func UpdateSearchResultById(id, status int) (err error) {
	err = global.GVA_DB.UpdateColumn("status", status).Where("id = ?", id).Error
	return err
}

func GetSearchResult(id uint) (err error, searchResult model.SearchResult) {
	err = global.GVA_DB.Where("id = ?", id).First(&searchResult).Error
	return
}

func GetSearchResultInfoList(info request.SearchResultSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.SearchResult{})
	var searchResults []model.SearchResult
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Query != "" {
		db = db.Where("`repo` LIKE ? or `text_matches_json` LIKE ?", "%"+info.Query+"%", "%"+info.Query+"%")
	}
	if info.Keyword != "" {
		db = db.Where("`keyword` = ?", info.Keyword)
	}
	if info.Status >= 0 {
		db = db.Where("`status` = ?", info.Status)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&searchResults).Error
	return err, searchResults, total
}

func CheckExistOfSearchResult(searchResult *model.SearchResult) bool {
	urlExist := searchResult.CheckPathExists()
	repoExists := searchResult.CheckRepoExists()
	return urlExist || repoExists
}
