package memo_domain

import (
	"gorm.io/gorm"
)

type MemoDTO struct {
	Memo
	gorm.Model
}

func ToMemoDTO(m Memo) MemoDTO {
	return MemoDTO{
		Memo:  m,
		Model: gorm.Model{},
	}
}

func FromMemoDTO(m MemoDTO) Memo {
	m.DisplayId = m.Model.ID
	return m.Memo
}
