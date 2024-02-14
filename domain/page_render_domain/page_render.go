package page_render_domain

import (
	"bytes"
	"github.com/itzmeerkat/icmd/domain/page_render_domain/model"
	"github.com/itzmeerkat/icmd/infra/config"
	"github.com/jedib0t/go-pretty/v6/table"
	"path/filepath"
	"text/template"
)

func RenderPage(content model.PageContent) string {
	data := map[string]string{}

	t := table.NewWriter()
	for _, e := range content.Todo {
		t.AppendRow(table.Row{e.DisplayId, e.Content})
	}
	data["todo"] = t.RenderHTML()

	t = table.NewWriter()
	for _, e := range content.Wishlist {
		t.AppendRow(table.Row{e.DisplayId, e.Content})
	}
	data["wishlist"] = t.RenderHTML()

	t = table.NewWriter()
	for _, e := range content.PlainMemo {
		t.AppendRow(table.Row{e.DisplayId, e.Content})
	}
	data["plain"] = t.RenderHTML()

	filePath := filepath.Join(config.GetAppFolder(), "template.html")
	templ, err := template.New("template.html").ParseFiles(filePath)
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	err = templ.Execute(&buf, data)
	return buf.String()
}
