package types

import (
	"gin-vben-admin/dao/models"
	"gin-vben-admin/pkg/e"
	"github.com/bwmarrin/snowflake"
	"time"
)

type DeptItem struct {
	ID        snowflake.ID  `json:"id" swaggertype:"string"`
	Name      string        `json:"name"`
	Sort      int64         `json:"sort"`
	Status    models.Status `json:"status"`
	Remark    string        `json:"remark"`
	ParentID  snowflake.ID  `json:"parentId,omitempty" swaggertype:"string"`
	Children  []*DeptItem   `json:"children,omitempty"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
}

type DeptItemsData struct {
	Items []*DeptItem `json:"items"`
}

func RespDeptItems(depts []*models.Dept) e.Response {
	data := DeptItemsData{
		Items: []*DeptItem{},
	}
	for _, dept := range depts {
		data.Items = append(data.Items, buildDeptData(dept))
	}
	return e.RespOK.WithData(data)
}

type DeptsTreeData struct {
	Total int64       `json:"total"`
	Items []*DeptItem `json:"items"`
}

func RespDeptsTree(depts []*models.Dept, count int64) e.Response {
	data := DeptsTreeData{
		Total: count,
		Items: []*DeptItem{},
	}
	for _, dept := range depts {
		data.Items = append(data.Items, buildDeptData(dept))
	}

	return e.RespOK.WithData(data)
}

func buildDeptData(dept *models.Dept) *DeptItem {
	item := &DeptItem{
		ID:        dept.ID,
		Name:      dept.Name,
		Sort:      dept.Sort,
		Status:    dept.Status,
		Remark:    dept.Remark,
		CreatedAt: dept.CreatedAt,
		ParentID:  dept.ParentID,
		UpdatedAt: dept.UpdatedAt,
		Children:  []*DeptItem{},
	}
	if dept.Depts != nil {
		for _, d := range dept.Depts {
			item.Children = append(item.Children, buildDeptData(d))
		}
	}
	return item
}
