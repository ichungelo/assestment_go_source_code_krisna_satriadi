package gormadapter

import (
	"github.com/ichungelo/assestment_go_source_code_krisna_satriadi/core/model"
)

func (g *gormAdapter) DeleteQuantityById(itemId *int, invoiceId *int) error {
	err := g.Model(&model.Quantity{}).Where("item_id = ?", itemId).Where("invoice_id = ?", invoiceId).Delete(&model.Quantity{}).Error
	if err != nil {
		return err
	}

	return nil

}
