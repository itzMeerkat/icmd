package model

import "github.com/itzmeerkat/icmd/domain/memo_domain/model"

type PageContent struct {
	Wishlist  []model.Memo
	Todo      []model.Memo
	PlainMemo []model.Memo
}
