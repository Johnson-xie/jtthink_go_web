package Services

import (
	"fmt"
	"product.jtthink.com/AppInit"
	"product.jtthink.com/Models"
)

//框架无关性代码
type BookService struct {
}

//加载图书列表
func (this *BookService) LoadBookList(req *BookListRequest) (*Models.BookList, error) {
	prods := &Models.BookList{}

	db := AppInit.GetDB().Limit(req.Size).Order("book_id desc").Find(prods)
	if db.Error != nil {
		return nil, db.Error
	}
	return prods, nil
}
func (this *BookService) LoadBookDetail(req *BookDetailRequest) (*Models.Books, []*Models.BookMetas, error) {
	prods := &Models.Books{}
	if AppInit.GetDB().Find(prods, req.BookID).RowsAffected != 1 {
		return nil, nil, fmt.Errorf("no book")
	}
	Models.NewBookMeta("click", "1", prods.BookID).Save()
	metas := []*Models.BookMetas{}
	AppInit.GetDB().Where("item_id=?", prods.BookID).Find(&metas)
	return prods, metas, nil
}
