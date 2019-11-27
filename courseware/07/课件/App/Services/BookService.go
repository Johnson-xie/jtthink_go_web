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
	err := AppInit.GetDB().Limit(req.Size).Order("book_id desc").Find(prods).Error
	if err != nil {
		return nil, err
	}
	return prods, nil
}
func (this *BookService) LoadBookDetail(req *BookDetailRequest) (*Models.Books, error) {
	prods := &Models.Books{}
	if AppInit.GetDB().Find(prods, req.BookID).RowsAffected != 1 {
		return nil, fmt.Errorf("no book")
	}
	return prods, nil
}
