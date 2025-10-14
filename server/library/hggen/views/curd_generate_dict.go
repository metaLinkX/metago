// Package views

package views

import (
	"bytes"
	"context"
	"fmt"
	"hash/fnv"
	"strings"

	"server/internal/dao"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mozillazg/go-pinyin"
)

// DictItem 字典项
type DictItem struct {
	Value     int64
	Label     string
	Name      string
	ValueType string
	Type      string
	ListClass string
	Extra     interface{}
}

// genHashListClass 根据label以hash算法生成表格回显样式
func genHashListClass(label string) string {
	strings := []string{"default", "primary", "info", "success", "warning", "error"}
	hash := fnv.New32()

	tag := "default"
	if _, err := hash.Write(gconv.Bytes(label)); err == nil {
		index := int(hash.Sum32()) % len(strings)
		if index < len(strings) {
			tag = strings[index]
		}
	}
	return tag
}

// GenerateDictConstFile 生成字典常量文件
func (l *gCurd) GenerateDictConstFile(ctx context.Context, dictType string) (content string, err error) {
	// 获取字典数据
	var dictDataList []*entity.SysDictData
	err = dao.SysDictData.Ctx(ctx).
		Where(dao.SysDictData.Columns().Type, dictType).
		Where(dao.SysDictData.Columns().Status, 1). // 只获取启用的字典项
		Order(dao.SysDictData.Columns().Sort + " ASC, " + dao.SysDictData.Columns().Id + " ASC").
		Scan(&dictDataList)
	if err != nil {
		return
	}

	if len(dictDataList) == 0 {
		return "", fmt.Errorf("字典类型 '%s' 未找到数据", dictType)
	}

	// 转换为字典项列表
	dictItems := make([]DictItem, 0, len(dictDataList))
	for _, item := range dictDataList {
		// 只处理整数类型的字典值
		if item.ValueType == "int" || item.ValueType == "" {
			value := gconv.Int64(item.Value)
			name := convertLabelToConstName(item.Label)
			dictItems = append(dictItems, DictItem{
				Value:     value,
				Label:     item.Label,
				Name:      name,
				ValueType: item.ValueType,
				Type:      item.Type,
				ListClass: "", // 将在生成时计算
				Extra:     nil,
			})
		}
	}

	// 生成文件内容
	return l.generateDictConstContent(dictType, dictItems)
}

// convertLabelToConstName 将标签转换为常量名称
func convertLabelToConstName(label string) string {
	// 移除空格和特殊字符，只保留字母、数字和中文
	name := gstr.Replace(label, " ", "")
	name = gstr.Replace(name, "-", "")
	name = gstr.Replace(name, "_", "")

	// 简单判断是否为英文
	isEnglish := true
	for _, r := range name {
		if r < 'A' || (r > 'Z' && r < 'a') || r > 'z' {
			if r < '0' || r > '9' {
				isEnglish = false
				break
			}
		}
	}

	if isEnglish {
		return gstr.CaseCamel(name)
	}

	// 对于中文，使用拼音库转换
	return convertChineseToPinyin(label)
}

// convertChineseToPinyin 将中文转换为拼音首字母大写
func convertChineseToPinyin(chinese string) string {
	// 设置拼音风格为首字母
	args := pinyin.NewArgs()
	args.Style = pinyin.FirstLetter

	// 获取拼音
	pys := pinyin.Pinyin(chinese, args)

	// 拼接拼音
	var result strings.Builder
	for _, py := range pys {
		if len(py) > 0 {
			// 首字母大写
			first := strings.ToUpper(py[0][:1])
			rest := py[0][1:]
			result.WriteString(first)
			result.WriteString(rest)
		}
	}

	pinyinStr := result.String()
	if pinyinStr == "" {
		return "DictItem"
	}

	return gstr.CaseCamel(pinyinStr)
}

// generateDictConstContent 生成字典常量文件内容
func (l *gCurd) generateDictConstContent(dictType string, items []DictItem) (content string, err error) {
	buffer := bytes.NewBuffer(nil)

	// 生成驼峰格式的字典类型名
	dictTypeName := gstr.CaseCamel(dictType)
	dictTypeName = strings.Replace(dictTypeName, " ", "", -1)

	// 文件头部
	buffer.WriteString(fmt.Sprintf("package consts\n\n"))
	buffer.WriteString(fmt.Sprintf("import (\n"))
	buffer.WriteString(fmt.Sprintf("\t\"hash/fnv\"\n"))
	buffer.WriteString(fmt.Sprintf("\t\"server/internal/model\"\n"))
	buffer.WriteString(fmt.Sprintf("\t\"github.com/gogf/gf/v2/util/gconv\"\n"))
	buffer.WriteString(fmt.Sprintf(")\n\n"))

	// 生成常量定义
	buffer.WriteString(fmt.Sprintf("const (\n"))
	for i, item := range items {
		constName := fmt.Sprintf("Dict%s%s", dictTypeName, item.Name)
		if item.Name == "DictItem" {
			// 如果是中文标签，使用序号
			constName = fmt.Sprintf("Dict%s%s%d", dictTypeName, item.Name, i+1)
		}

		buffer.WriteString(fmt.Sprintf("\t%s int = %d // %s\n", constName, item.Value, item.Label))
	}
	buffer.WriteString(fmt.Sprintf(")\n\n"))

	// 生成变量定义
	buffer.WriteString(fmt.Sprintf("var (\n"))

	// Map变量
	buffer.WriteString(fmt.Sprintf("\tDict%sMap = map[int]string{\n", dictTypeName))
	for i, item := range items {
		constName := fmt.Sprintf("Dict%s%s", dictTypeName, item.Name)
		if item.Name == "DictItem" {
			constName = fmt.Sprintf("Dict%s%s%d", dictTypeName, item.Name, i+1)
		}
		buffer.WriteString(fmt.Sprintf("\t\t%s: \"%s\",\n", constName, item.Label))
	}
	buffer.WriteString(fmt.Sprintf("\t}\n"))

	// Option变量
	buffer.WriteString(fmt.Sprintf("\tDict%sOption = []model.Option{\n", dictTypeName))
	for i, item := range items {
		constName := fmt.Sprintf("Dict%s%s", dictTypeName, item.Name)
		if item.Name == "DictItem" {
			constName = fmt.Sprintf("Dict%s%s%d", dictTypeName, item.Name, i+1)
		}

		// 使用hash算法生成ListClass
		listClass := genHashListClass(item.Label)

		buffer.WriteString(fmt.Sprintf("\t\t{\n"))
		buffer.WriteString(fmt.Sprintf("\t\t\tKey:       %s,\n", constName))
		buffer.WriteString(fmt.Sprintf("\t\t\tLabel:     \"%s\",\n", item.Label))
		buffer.WriteString(fmt.Sprintf("\t\t\tValue:     %s,\n", constName))
		buffer.WriteString(fmt.Sprintf("\t\t\tValueType: \"%s\",\n", item.ValueType))
		buffer.WriteString(fmt.Sprintf("\t\t\tType:      \"%s\",\n", item.Type))
		buffer.WriteString(fmt.Sprintf("\t\t\tListClass: \"%s\",\n", listClass))
		buffer.WriteString(fmt.Sprintf("\t\t\tExtra:     nil,\n"))
		buffer.WriteString(fmt.Sprintf("\t\t},\n"))
	}
	buffer.WriteString(fmt.Sprintf("\t}\n"))

	buffer.WriteString(fmt.Sprintf(")\n"))

	return buffer.String(), nil
}
