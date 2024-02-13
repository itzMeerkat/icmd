package memo_domain

import (
	"github.com/itzmeerkat/icmd/infra"
	"gorm.io/gorm"
)

type Memo struct {
	Content    string
	DisplayId  uint
	IsWishlist bool
	IsTodo     bool
	Status     int
}

func init() {
	infra.MigrateType(&MemoDTO{})
}

func AddMemo(me Memo) {
	dto := ToMemoDTO(me)
	infra.CreateRecord(&dto)
}

func GetAllMemo() []Memo {
	var memos []MemoDTO
	infra.GetRecords(&memos, nil)
	var res []Memo
	for _, e := range memos {
		res = append(res, FromMemoDTO(e))
	}
	return res
}

func GetAllTodo() []Memo {
	var memos []MemoDTO
	infra.GetRecords(&memos, MemoDTO{Memo{IsTodo: true}, gorm.Model{}})
	var res []Memo
	for _, e := range memos {
		res = append(res, FromMemoDTO(e))
	}
	return res
}

func GetAllWishlist() []Memo {
	var memos []MemoDTO
	infra.GetRecords(&memos, MemoDTO{Memo{IsWishlist: true}, gorm.Model{}})
	var res []Memo
	for _, e := range memos {
		res = append(res, FromMemoDTO(e))
	}
	return res
}
