package acl

import (
	memo_model "github.com/itzmeerkat/icmd/domain/memo_domain/model"
	page_render_model "github.com/itzmeerkat/icmd/domain/page_render_domain/model"
)

func NewPageContent(wishlist, todo, plain []memo_model.Memo) page_render_model.PageContent {
	return page_render_model.PageContent{
		Wishlist:  wishlist,
		Todo:      todo,
		PlainMemo: plain,
	}
}
