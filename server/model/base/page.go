package base

//分页 基础对象
type BasePage struct {
	Page   uint64      `form:"page" json:"page"`   //当前页
	Offset uint64      `json:"offset"`             //过滤条数
	Limit  uint64      `form:"limit" json:"limit"` // 每页条数
	Total  uint64      `json:"total"`              // 总条数
	List   interface{} `json:"list"`               //记录

}

//获取offset
func (b BasePage) GetOffset() uint64 {
	return (b.Page - 1) * b.Limit
}
