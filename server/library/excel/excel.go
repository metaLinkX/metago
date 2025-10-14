// Package excel

package excel

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/xuri/excelize/v2"
	"net/url"
	"reflect"
	"server/internal/model"
	"server/library/contexts"
	"time"
	"unicode"
)

var (
	// 单元格表头
	char = []string{"", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	// 默认行样式
	defaultRowStyle = &excelize.Style{
		Font: &excelize.Font{
			Color:  "#666666",
			Size:   13,
			Family: "Arial",
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	}
	defaultTopRowStyle = &excelize.Style{
		Font: &excelize.Font{
			Color:  "#666666",
			Size:   13,
			Family: "Arial",
		},
		Alignment: &excelize.Alignment{
			Horizontal: "left",
			Vertical:   "top",
			WrapText:   true,
		},
	}
)

// ExportByStructs 导出切片结构体到excel表格
func ExportByStructs(ctx context.Context, tags []string, list interface{}, fileName string, sheetName string) (err error) {
	f := excelize.NewFile()
	f.SetSheetName("Sheet1", sheetName)
	_ = f.SetRowHeight("Sheet1", 1, 30)

	rowStyleID, _ := f.NewStyle(defaultRowStyle)
	if err != nil {
		return
	}
	_ = f.SetSheetRow(sheetName, "A1", &tags)

	var (
		length    = len(tags)
		headStyle = letter(length)
		lastRow   string
		widthRow  string
	)

	for k, v := range headStyle {
		if k == length-1 {
			lastRow = fmt.Sprintf("%s1", v)
			widthRow = v
		}
	}

	if err = f.SetColWidth(sheetName, "A", widthRow, 30); err != nil {
		return err
	}

	var rowNum = 1
	for _, v := range gconv.Interfaces(list) {
		t := reflect.TypeOf(v)
		value := reflect.ValueOf(v)
		row := make([]interface{}, 0)
		for l := 0; l < t.NumField(); l++ {
			val := value.Field(l).Interface()
			row = append(row, val)
		}
		rowNum++
		if err = f.SetSheetRow(sheetName, "A"+gconv.String(rowNum), &row); err != nil {
			return
		}
		if err = f.SetCellStyle(sheetName, fmt.Sprintf("A%d", rowNum), lastRow, rowStyleID); err != nil {
			return
		}
	}

	r := ghttp.RequestFromCtx(ctx)
	if r == nil {
		err = gerror.New("ctx not http request")
		return
	}

	writer := r.Response.Writer
	writer.Header().Set("Content-Type", "application/octet-stream")
	writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.xlsx", url.QueryEscape(fileName)))
	writer.Header().Set("Content-Transfer-Encoding", "binary")
	writer.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")

	if err = f.Write(writer); err != nil {
		return
	}

	// 加入到上下文
	contexts.SetResponse(ctx, &model.Response{
		Code:      gcode.CodeOK.Code(),
		Message:   "export successfully!",
		Timestamp: time.Now().Unix(),
		TraceID:   gctx.CtxId(ctx),
	})
	return
}

func ExportExplainByStructs(ctx context.Context, tags []string, list interface{}, fileName string, sheetName string, topDesc string, topLineHeight float64) (err error) {
	f := excelize.NewFile()
	f.SetSheetName("Sheet1", sheetName)
	_ = f.SetRowHeight("Sheet1", 1, topLineHeight)

	rowStyleID, _ := f.NewStyle(defaultRowStyle)
	if err != nil {
		return
	}
	rowTopStyleID, _ := f.NewStyle(defaultTopRowStyle)
	if err != nil {
		return
	}

	// Merge first row cells and set content to "京津冀"
	length := len(tags)
	headStyle := letter(length)
	lastCol := headStyle[length-1]
	if err = f.MergeCell(sheetName, "A1", fmt.Sprintf("%s1", lastCol)); err != nil {
		return
	}
	if err = f.SetCellValue(sheetName, "A1", topDesc); err != nil {
		return
	}
	if err = f.SetCellStyle(sheetName, "A1", fmt.Sprintf("%s1", lastCol), rowTopStyleID); err != nil {
		return
	}

	// Set header from second row (A2)
	_ = f.SetSheetRow(sheetName, "A2", &tags)
	_ = f.SetCellStyle(sheetName, "A2", fmt.Sprintf("%s1", lastCol), rowTopStyleID)

	var (
		lastRow  string
		widthRow string
	)

	for k, v := range headStyle {
		if k == length-1 {
			lastRow = fmt.Sprintf("%s2", v) // Updated to row 2
			widthRow = v
		}
	}

	if err = f.SetColWidth(sheetName, "A", widthRow, 30); err != nil {
		return err
	}

	var rowNum = 2 // Start from row 2
	for _, v := range gconv.Interfaces(list) {
		t := reflect.TypeOf(v)
		value := reflect.ValueOf(v)
		row := make([]interface{}, 0)
		for l := 0; l < t.NumField(); l++ {
			val := value.Field(l).Interface()
			row = append(row, val)
		}
		rowNum++
		// Write data from row 3 onwards
		if err = f.SetSheetRow(sheetName, "A"+gconv.String(rowNum), &row); err != nil {
			return
		}
		if err = f.SetCellStyle(sheetName, fmt.Sprintf("A%d", rowNum), lastRow, rowStyleID); err != nil {
			return
		}
	}

	r := ghttp.RequestFromCtx(ctx)
	if r == nil {
		err = gerror.New("ctx not http request")
		return
	}

	writer := r.Response.Writer
	writer.Header().Set("Content-Type", "application/octet-stream")
	writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.xlsx", url.QueryEscape(fileName)))
	writer.Header().Set("Content-Transfer-Encoding", "binary")
	writer.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")

	if err = f.Write(writer); err != nil {
		return
	}

	// 加入到上下文
	contexts.SetResponse(ctx, &model.Response{
		Code:      gcode.CodeOK.Code(),
		Message:   "export successfully!",
		Timestamp: time.Now().Unix(),
		TraceID:   gctx.CtxId(ctx),
	})
	return
}

func GetImportList(ctx context.Context, titles []string, sheetName string) (data []map[string]string, err error) {
	r := g.RequestFromCtx(ctx)
	file := r.GetUploadFile("file")

	f, err := file.Open()
	if err != nil {
		g.Log().Error(ctx, "读取文件失败，请重新上传", err)
		err = gerror.NewCode(gcode.CodeBusinessValidationFailed, "读取文件失败，请重新上传")
		return
	}
	defer f.Close()

	ff, err := excelize.OpenReader(f)
	defer ff.Close()
	if err != nil {
		g.Log().Error(ctx, "读取excel文件失败", err)
		err = gerror.NewCode(gcode.CodeBusinessValidationFailed, "读取excel文件失败")
		return
	}
	rows, err := ff.GetRows(sheetName)
	if err != nil {
		return
	}
	data = make([]map[string]string, 0)
	for rowIndex, row := range rows {
		if rowIndex == 0 {
			continue
		}
		if rowIndex == 1 {
			// 校验标题头
			for titleIndex, title := range titles {
				if (len(row) - 1) < titleIndex {
					err = gerror.NewCode(gcode.CodeBusinessValidationFailed, "标题头长度不对，请严格按照模板标题头来")
					return data, err
				}
				rowTitle := CleanNamespace(row[titleIndex])
				if rowTitle != title {
					err = gerror.NewCodef(gcode.CodeBusinessValidationFailed, "标题头不正确，模板标题头为【%s】，导入的标题头为【%s】,请严格按照模板标题头来", title, rowTitle)
					return data, err
				}
			}
			continue
		}
		item := make(map[string]string)
		for titleIndex, title := range titles {
			content := TrimNamespace(row[titleIndex])
			if gstr.Contains(title, "(必填)") && content == "" {
				err = gerror.NewCodef(gcode.CodeBusinessValidationFailed, "第%d行第%d列【%s】必填，当前为空", rowIndex+1, titleIndex+1, title)
				return data, err
			}
			item[title] = content
		}
		data = append(data, item)
	}
	return data, err
}

// letter 生成完整的表头
func letter(length int) []string {
	var str []string
	for i := 1; i <= length; i++ {
		str = append(str, numToChars(i))
	}
	return str
}

// numToChars 将数字转换为具体的表格表头名称
func numToChars(num int) string {
	var cols string
	v := num
	for v > 0 {
		k := v % 26
		if k == 0 {
			k = 26
		}
		v = (v - k) / 26
		cols = char[k] + cols
	}
	return cols
}

// NextLetter 传入一个字母，获取下一个字母
func NextLetter(input string) string {
	if len(input) == 0 {
		return ""
	}
	upperInput := unicode.ToUpper(rune(input[0]))
	if upperInput >= 'A' && upperInput < 'Z' {
		return string(upperInput + 1)
	}
	return "A"
}

type ColMateModel struct {
	Title    string
	Index    int
	Required bool
}

func CleanNamespace(str string) string {
	return gstr.ReplaceByArray(str, []string{"\n", " ", "\t"})
}
func TrimNamespace(str string) string {
	return gstr.TrimAll(str, []string{"\n", " ", "\t"}...)
}
