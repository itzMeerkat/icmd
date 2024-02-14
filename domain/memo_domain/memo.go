package memo_domain

import (
	"github.com/itzmeerkat/icmd/domain/memo_domain/model"
	"github.com/itzmeerkat/icmd/infra/database"
)

func init() {
	database.MigrateType(&model.MemoDTO{})
}

func AddMemo(me model.Memo) {
	dto := model.ToMemoDTO(me)
	database.CreateRecord(&dto)
}

func GetAllMemo() []model.Memo {
	var memos []model.MemoDTO
	database.GetRecords(&memos, nil)
	var res []model.Memo
	for _, e := range memos {
		res = append(res, model.FromMemoDTO(e))
	}
	return res
}

func GetAllTodo() []model.Memo {
	var memos []model.MemoDTO
	database.GetRecords(&memos, model.MemoDTO{Memo: model.Memo{IsTodo: true}})
	var res []model.Memo
	for _, e := range memos {
		res = append(res, model.FromMemoDTO(e))
	}
	return res
}

func GetAllWishlist() []model.Memo {
	var memos []model.MemoDTO
	database.GetRecords(&memos, model.MemoDTO{Memo: model.Memo{IsWishlist: true}})
	var res []model.Memo
	for _, e := range memos {
		res = append(res, model.FromMemoDTO(e))
	}
	return res
}

func GetAllPlain() []model.Memo {
	var memos []model.MemoDTO
	database.GetRecords(&memos, model.MemoDTO{Memo: model.Memo{IsWishlist: false, IsTodo: false}})
	var res []model.Memo
	for _, e := range memos {
		res = append(res, model.FromMemoDTO(e))
	}
	return res
}
