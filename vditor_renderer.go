// Lute - A structured markdown engine.
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under the Mulan PSL v1.
// You can use this software according to the terms and conditions of the Mulan PSL v1.
// You may obtain a copy of Mulan PSL v1 at:
//     http://license.coscl.org.cn/MulanPSL
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR
// PURPOSE.
// See the Mulan PSL v1 for more details.

package lute

// Vditor DOM Renderer

import (
	"strconv"
	"strings"
)

// VditorRenderer 描述了 Vditor DOM 渲染器。
type VditorRenderer struct {
	*BaseRenderer
}

// newVditorRenderer 创建一个 Vditor DOM 渲染器。
func (lute *Lute) newVditorRenderer(treeRoot *Node) *VditorRenderer {
	ret := &VditorRenderer{&BaseRenderer{rendererFuncs: map[nodeType]RendererFunc{}, option: lute.options, treeRoot: treeRoot}}
	ret.rendererFuncs[NodeDocument] = ret.renderDocument
	ret.rendererFuncs[NodeParagraph] = ret.renderParagraph
	ret.rendererFuncs[NodeText] = ret.renderText
	ret.rendererFuncs[NodeCodeSpan] = ret.renderCodeSpan
	ret.rendererFuncs[NodeCodeBlock] = ret.renderCodeBlock
	ret.rendererFuncs[NodeMathBlock] = ret.renderMathBlock
	ret.rendererFuncs[NodeInlineMath] = ret.renderInlineMath
	ret.rendererFuncs[NodeEmphasis] = ret.renderEmphasis
	ret.rendererFuncs[NodeEmA6kOpenMarker] = ret.renderEmAsteriskOpenMarker
	ret.rendererFuncs[NodeEmA6kCloseMarker] = ret.renderEmAsteriskCloseMarker
	ret.rendererFuncs[NodeEmU8eOpenMarker] = ret.renderEmUnderscoreOpenMarker
	ret.rendererFuncs[NodeEmU8eCloseMarker] = ret.renderEmUnderscoreCloseMarker
	ret.rendererFuncs[NodeStrong] = ret.renderStrong
	ret.rendererFuncs[NodeStrongA6kOpenMarker] = ret.renderStrongA6kOpenMarker
	ret.rendererFuncs[NodeStrongA6kCloseMarker] = ret.renderStrongA6kCloseMarker
	ret.rendererFuncs[NodeStrongU8eOpenMarker] = ret.renderStrongU8eOpenMarker
	ret.rendererFuncs[NodeStrongU8eCloseMarker] = ret.renderStrongU8eCloseMarker
	ret.rendererFuncs[NodeBlockquote] = ret.renderBlockquote
	ret.rendererFuncs[NodeBlockquoteMarker] = ret.renderBlockquoteMarker
	ret.rendererFuncs[NodeHeading] = ret.renderHeading
	ret.rendererFuncs[NodeList] = ret.renderList
	ret.rendererFuncs[NodeListItem] = ret.renderListItem
	ret.rendererFuncs[NodeThematicBreak] = ret.renderThematicBreak
	ret.rendererFuncs[NodeHardBreak] = ret.renderHardBreak
	ret.rendererFuncs[NodeSoftBreak] = ret.renderSoftBreak
	ret.rendererFuncs[NodeHTMLBlock] = ret.renderHTML
	ret.rendererFuncs[NodeInlineHTML] = ret.renderInlineHTML
	ret.rendererFuncs[NodeLink] = ret.renderLink
	ret.rendererFuncs[NodeImage] = ret.renderImage
	ret.rendererFuncs[NodeStrikethrough] = ret.renderStrikethrough
	ret.rendererFuncs[NodeStrikethrough1OpenMarker] = ret.renderStrikethrough1OpenMarker
	ret.rendererFuncs[NodeStrikethrough1CloseMarker] = ret.renderStrikethrough1CloseMarker
	ret.rendererFuncs[NodeStrikethrough2OpenMarker] = ret.renderStrikethrough2OpenMarker
	ret.rendererFuncs[NodeStrikethrough2CloseMarker] = ret.renderStrikethrough2CloseMarker
	ret.rendererFuncs[NodeTaskListItemMarker] = ret.renderTaskListItemMarker
	ret.rendererFuncs[NodeTable] = ret.renderTable
	ret.rendererFuncs[NodeTableHead] = ret.renderTableHead
	ret.rendererFuncs[NodeTableRow] = ret.renderTableRow
	ret.rendererFuncs[NodeTableCell] = ret.renderTableCell
	ret.rendererFuncs[NodeEmojiUnicode] = ret.renderEmojiUnicode
	ret.rendererFuncs[NodeEmojiImg] = ret.renderEmojiImg
	return ret
}

func (r *VditorRenderer) renderEmojiImg(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.write(node.tokens)
	}
	return WalkStop, nil
}

func (r *VditorRenderer) renderEmojiUnicode(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.write(node.tokens)
	}
	return WalkStop, nil
}

func (r *VditorRenderer) renderTableCell(node *Node, entering bool) (WalkStatus, error) {
	tag := "td"
	if NodeTableHead == node.parent.typ {
		tag = "th"
	}
	if entering {
		var attrs [][]string
		switch node.tableCellAlign {
		case 1:
			attrs = append(attrs, []string{"align", "left"})
		case 2:
			attrs = append(attrs, []string{"align", "center"})
		case 3:
			attrs = append(attrs, []string{"align", "right"})
		}
		r.tag(tag, node, attrs, false)
	} else {
		r.tag("/"+tag, node, nil, false)
	}
	return WalkContinue, nil
}

func (r *VditorRenderer) renderTableRow(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.tag("tr", node, nil, false)
	} else {
		r.tag("/tr", node, nil, false)
		if node == node.parent.lastChild {
			r.tag("/tbody", node, nil, false)
		}
	}
	return WalkContinue, nil
}

func (r *VditorRenderer) renderTableHead(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.tag("thead", node, nil, false)
		r.tag("tr", node, nil, false)
	} else {
		r.tag("/tr", node, nil, false)
		r.tag("/thead", node, nil, false)
		if nil != node.next {
			r.tag("tbody", node, nil, false)
		}
	}
	return WalkContinue, nil
}

func (r *VditorRenderer) renderTable(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.tag("table", node, nil, false)
	} else {
		r.tag("/table", node, nil, false)
	}
	return WalkContinue, nil
}

func (r *VditorRenderer) renderStrikethrough(node *Node, entering bool) (WalkStatus, error) {
	return WalkContinue, nil
}

func (r *VditorRenderer) renderStrikethrough1OpenMarker(node *Node, entering bool) (WalkStatus, error) {
	r.tag("del", node, nil, false)
	return WalkStop, nil
}

func (r *VditorRenderer) renderStrikethrough1CloseMarker(node *Node, entering bool) (WalkStatus, error) {
	r.tag("/del", node, nil, false)
	return WalkStop, nil
}

func (r *VditorRenderer) renderStrikethrough2OpenMarker(node *Node, entering bool) (WalkStatus, error) {
	r.tag("del", node, nil, false)
	return WalkStop, nil
}

func (r *VditorRenderer) renderStrikethrough2CloseMarker(node *Node, entering bool) (WalkStatus, error) {
	r.tag("/del", node, nil, false)
	return WalkStop, nil
}

func (r *VditorRenderer) renderImage(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		if 0 == r.disableTags {
			r.writeString("<img src=\"")
			r.write(escapeHTML(node.destination))
			r.writeString("\" alt=\"")
		}
		r.disableTags++
		return WalkContinue, nil
	}

	r.disableTags--
	if 0 == r.disableTags {
		r.writeString("\"")
		if nil != node.title {
			r.writeString(" title=\"")
			r.write(escapeHTML(node.title))
			r.writeString("\"")
		}
		r.writeString(" />")
	}
	return WalkContinue, nil
}

func (r *VditorRenderer) renderLink(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.tag("span", nil, nil, false)
		attrs := [][]string{{"class", "marker"}}
		r.tag("span", nil, attrs, false)
		r.writeByte(itemOpenBracket)
		r.tag("/span", nil, nil, false)

		attrs = [][]string{{"href", itemsToStr(escapeHTML(node.destination))}}
		if nil != node.title {
			attrs = append(attrs, []string{"title", itemsToStr(escapeHTML(node.title))})
		}
		r.tag("a", nil, attrs, false)
	} else {
		r.tag("/a", node, nil, false)
		attrs := [][]string{{"class", "marker"}}
		r.tag("span", nil, attrs, false)
		r.writeByte(itemCloseBracket)
		r.tag("/span", nil, nil, false)
		attrs = [][]string{{"class", "marker"}}
		r.tag("span", nil, attrs, false)
		r.writeByte(itemOpenParen)
		r.tag("/span", nil, nil, false)
		r.tag("span", nil, nil, false)
		r.write(node.destination)
		r.tag("/span", nil, nil, false)
		// TODO: title
		attrs = [][]string{{"class", "marker"}}
		r.tag("span", nil, attrs, false)
		r.writeByte(itemCloseParen)
		r.tag("/span", nil, nil, false)
		r.tag("/span", nil, nil, false)
	}

	return WalkContinue, nil
}

func (r *VditorRenderer) renderHTML(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.write(node.tokens)
	}
	return WalkContinue, nil
}

func (r *VditorRenderer) renderInlineHTML(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.write(node.tokens)
	}
	return WalkContinue, nil
}

func (r *VditorRenderer) renderDocument(node *Node, entering bool) (WalkStatus, error) {
	return WalkContinue, nil
}

func (r *VditorRenderer) renderParagraph(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.tag("p", node, nil, false)
	} else {
		r.writeString("</p><span><br><span class=\"newline\">\n</span><span class=\"newline\">\n</span></span>")
	}
	return WalkContinue, nil
}

func (r *VditorRenderer) renderText(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.tag("span", node, nil, false)
		r.write(escapeHTML(node.tokens))
	} else {
		r.tag("/span", node, nil, false)
	}
	return WalkContinue, nil
}

func (r *VditorRenderer) renderCodeSpan(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.tag("span", nil, nil, false)
		attrs := [][]string{{"class", "marker"}}
		r.tag("span", nil, attrs, false)
		r.writeByte(itemBacktick)
		if 1 < node.codeMarkerLen {
			r.writeByte(itemBacktick)
		}
		r.tag("/span", nil, nil, false)
		r.tag("code", node, nil, false)
		r.write(escapeHTML(node.tokens))
	} else {
		r.tag("/code", node, nil, false)
		attrs := [][]string{{"class", "marker"}}
		r.tag("span", nil, attrs, false)
		r.writeByte(itemBacktick)
		if 1 < node.codeMarkerLen {
			r.writeByte(itemBacktick)
		}
		r.tag("/span", nil, nil, false)
	}
	return WalkContinue, nil
}

func (r *VditorRenderer) renderInlineMath(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		tokens := node.tokens
		tokens = escapeHTML(tokens)
		r.write(tokens)
	}
	return WalkStop, nil
}

func (r *VditorRenderer) renderMathBlock(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.newline()
		tokens := node.tokens
		tokens = escapeHTML(tokens)
		r.write(tokens)
		r.newline()
	}
	return WalkStop, nil
}

func (r *VditorRenderer) renderCodeBlock(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		tokens := node.tokens
		if 0 < len(node.codeBlockInfo) {
			infoWords := split(node.codeBlockInfo, itemSpace)
			language := infoWords[0]
			r.writeString("<pre><code class=\"language-")
			r.write(language)
			r.writeString("\">")
			tokens = escapeHTML(tokens)
			r.write(tokens)
		} else {
			r.writeString("<pre><code>")
			tokens = escapeHTML(tokens)
			r.write(tokens)
		}
		return WalkSkipChildren, nil
	}
	r.writeString("</code></pre>")
	return WalkContinue, nil
}

func (r *VditorRenderer) renderEmphasis(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		attrs := [][]string{{"class", "node"}}
		r.tag("span", node, attrs, false)
	} else {
		r.tag("/span", nil, nil, false)
	}
	return WalkContinue, nil
}

func (r *VditorRenderer) renderEmAsteriskOpenMarker(node *Node, entering bool) (WalkStatus, error) {
	attrs := [][]string{{"class", "marker"}}
	r.tag("span", node, attrs, false)
	r.writeByte(itemAsterisk)
	r.tag("/span", nil, nil, false)
	r.tag("em", node.parent, nil, false)
	return WalkStop, nil
}

func (r *VditorRenderer) renderEmAsteriskCloseMarker(node *Node, entering bool) (WalkStatus, error) {
	r.tag("/em", node, nil, false)
	attrs := [][]string{{"class", "marker"}}
	r.tag("span", node, attrs, false)
	r.writeByte(itemAsterisk)
	r.tag("/span", nil, nil, false)
	return WalkStop, nil
}

func (r *VditorRenderer) renderEmUnderscoreOpenMarker(node *Node, entering bool) (WalkStatus, error) {
	attrs := [][]string{{"class", "marker"}}
	r.tag("span", nil, attrs, false)
	r.writeByte(itemUnderscore)
	r.tag("/span", nil, nil, false)
	r.tag("em", node, nil, false)
	return WalkStop, nil
}

func (r *VditorRenderer) renderEmUnderscoreCloseMarker(node *Node, entering bool) (WalkStatus, error) {
	attrs := [][]string{{"class", "marker"}}
	r.tag("span", nil, attrs, false)
	r.writeByte(itemUnderscore)
	r.tag("/span", nil, nil, false)
	r.tag("/em", node, nil, false)
	return WalkStop, nil
}

func (r *VditorRenderer) renderStrong(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		attrs := [][]string{{"class", "node"}}
		r.tag("span", nil, attrs, false)
	} else {
		r.tag("/span", nil, nil, false)
	}
	return WalkContinue, nil
}

func (r *VditorRenderer) renderStrongA6kOpenMarker(node *Node, entering bool) (WalkStatus, error) {
	attrs := [][]string{{"class", "marker"}}
	r.tag("span", nil, attrs, false)
	r.writeString("**")
	r.tag("/span", nil, nil, false)
	r.tag("strong", node, nil, false)
	return WalkStop, nil
}

func (r *VditorRenderer) renderStrongA6kCloseMarker(node *Node, entering bool) (WalkStatus, error) {
	attrs := [][]string{{"class", "marker"}}
	r.tag("span", nil, attrs, false)
	r.writeString("**")
	r.tag("/span", nil, nil, false)
	r.tag("/strong", node, nil, false)
	return WalkStop, nil
}

func (r *VditorRenderer) renderStrongU8eOpenMarker(node *Node, entering bool) (WalkStatus, error) {
	attrs := [][]string{{"class", "marker"}}
	r.tag("span", nil, attrs, false)
	r.writeString("__")
	r.tag("/span", nil, nil, false)
	r.tag("strong", node, nil, false)
	return WalkStop, nil
}

func (r *VditorRenderer) renderStrongU8eCloseMarker(node *Node, entering bool) (WalkStatus, error) {
	attrs := [][]string{{"class", "marker"}}
	r.tag("span", nil, attrs, false)
	r.writeString("__")
	r.tag("/span", nil, nil, false)
	r.tag("/strong", node, nil, false)
	return WalkStop, nil
}

func (r *VditorRenderer) renderBlockquote(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.tag("blockquote", node, nil, false)
	} else {
		r.tag("/blockquote", node, nil, false)
	}
	return WalkContinue, nil
}

func (r *VditorRenderer) renderBlockquoteMarker(node *Node, entering bool) (WalkStatus, error) {
	attrs := [][]string{{"class", "node"}}
	r.tag("span", nil, attrs, false)
	attrs = [][]string{{"class", "marker"}}
	r.tag("span", node, attrs, false)
	r.writeString("&gt;")
	r.tag("/span", node, nil, false)
	r.tag("/span", nil, nil, false)
	return WalkStop, nil
}

func (r *VditorRenderer) renderHeading(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.tag("h"+" 123456"[node.headingLevel:node.headingLevel+1], node, nil, false)
	} else {
		r.writeString("</h" + " 123456"[node.headingLevel:node.headingLevel+1] + ">")
	}
	return WalkContinue, nil
}

func (r *VditorRenderer) renderList(node *Node, entering bool) (WalkStatus, error) {
	tag := "ul"
	if 1 == node.listData.typ {
		tag = "ol"
	}
	if entering {
		attrs := [][]string{{"start", strconv.Itoa(node.start)}}
		if nil == node.bulletChar && 1 != node.start {
			r.tag(tag, node, attrs, false)
		} else {
			r.tag(tag, node, nil, false)
		}
	} else {
		r.tag("/"+tag, node, nil, false)
	}
	return WalkContinue, nil
}

func (r *VditorRenderer) renderListItem(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		if 3 == node.listData.typ && "" != r.option.GFMTaskListItemClass {
			r.tag("li", node, [][]string{{"class", r.option.GFMTaskListItemClass}}, false)
		} else {
			r.tag("li", node, nil, false)
		}
		attrs := [][]string{{"class", "node"}}
		r.tag("span", nil, attrs, false)
		attrs = [][]string{{"class", "marker"}}
		r.tag("span", nil, attrs, false)

		marker := node.listData.marker
		if !isNilItem(node.listData.delimiter) {
			marker = append(marker, node.listData.delimiter)
		}
		r.writeString(itemsToStr(marker) + " ")
		r.tag("/span", nil, nil, false)
		r.tag("/span", nil, nil, false)
		r.tag("p", nil, nil, false)
	} else {
		r.tag("/p", nil, nil, false)
		r.tag("/li", node, nil, false)
	}
	return WalkContinue, nil
}

func (r *VditorRenderer) renderTaskListItemMarker(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		var attrs [][]string
		if node.taskListItemChecked {
			attrs = append(attrs, []string{"checked", ""})
		}
		attrs = append(attrs, []string{"disabled", ""}, []string{"type", "checkbox"})
		r.tag("input", node, attrs, true)
	}
	return WalkContinue, nil
}

func (r *VditorRenderer) renderThematicBreak(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.tag("hr", node, nil, true)
	}
	return WalkSkipChildren, nil
}

func (r *VditorRenderer) renderHardBreak(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.tag("span", node, nil, false)
		r.tag("br", nil, nil, false)
		attrs := [][]string{{"class", "newline"}}
		r.tag("span", node, attrs, true)
		r.writeByte(itemNewline)
		r.tag("/span", node, nil, false)
	}
	return WalkStop, nil
}

func (r *VditorRenderer) renderSoftBreak(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.tag("span", node, nil, false)
		r.tag("br", nil, nil, false)
		attrs := [][]string{{"class", "newline"}}
		r.tag("span", node, attrs, true)
		r.writeByte(itemNewline)
		r.tag("/span", node, nil, false)
	}
	return WalkStop, nil
}

func (r *VditorRenderer) tag(name string, node *Node, attrs [][]string, selfclosing bool) {
	if r.disableTags > 0 {
		return
	}

	r.writeString("<")
	r.writeString(name)

	isClosing := itemSlash == name[0]
	if !isClosing {
		if nil == attrs {
			attrs = [][]string{}
		}
		if nil != node {
			attrs = append(attrs, []string{"data-ntype", node.typ.String()})
			attrs = append(attrs, []string{"data-mtype", r.mtype(node.typ)})
			// TODO: 源码位置
			//attrs = append(attrs, []string{"data-pos-start", strconv.Itoa(node.ranges[0].startLn) + ":" + strconv.Itoa(node.ranges[0].startCol)})
			//attrs = append(attrs, []string{"data-pos-end", strconv.Itoa(node.ranges[0].endLn) + ":" + strconv.Itoa(node.ranges[0].endCol)})
			if "" != node.caret {
				attrs = append(attrs, []string{"data-caret", node.caret}, []string{"data-caretoffset", strconv.Itoa(node.caretOffset)})
			}
			if node.expand {
				r.appendClass(&attrs, "node--expand")
			}
		}
		for _, attr := range attrs {
			r.writeString(" " + attr[0] + "=\"" + attr[1] + "\"")
		}
	}
	if selfclosing {
		r.writeString(" /")
	}
	r.writeString(">")
}

func (r *VditorRenderer) appendClass(attrs *[][]string, class string) {
	for _, attr := range *attrs {
		if "class" == attr[0] && !strings.Contains(attr[1], class) {
			attr[1] += " " + class
			return
		}
	}
	*attrs = append(*attrs, []string{"class", class})
}

// mtype 返回节点类型 nodeType 对应的 Markdown 元素类型。
//   0：叶子块元素
//   1：容器块元素
//   2：行级元素
func (r *VditorRenderer) mtype(nodeType nodeType) string {
	switch nodeType {
	case NodeThematicBreak, NodeHeading, NodeCodeBlock, NodeMathBlock, NodeHTMLBlock, NodeParagraph:
		return "0"
	case NodeBlockquote, NodeList, NodeListItem:
		return "1"
	default:
		return "2"
	}
}

// mapSelection 用于映射文本选段。
// 根据 Markdown 原文中的 startOffset 和 endOffset 位置从根节点 root 上遍历查找该范围内的节点，找到对应节点后进行标记以支持后续渲染。
func (r *VditorRenderer) mapSelection(root *Node, startOffset, endOffset int) {
	var nodes []*Node
	for c := root.firstChild; nil != c; c = c.next {
		r.findSelection(root, startOffset, endOffset, &nodes)
	}
	for _, node := range nodes {
		node.caret = "start"
		node.caretOffset = 1024
	}
	if 0 < len(nodes) {
		r.expand(nodes[0])
	}
}

// expand 用于在 node 上或者 node 的祖先节点上标记展开。
func (r *VditorRenderer) expand(node *Node) {
	for p := node; nil != p; p = p.parent {
		if "0" == r.mtype(p.typ) {
			return
		}
		switch p.typ {
		case NodeEmphasis, NodeStrong, NodeBlockquoteMarker, NodeListItem:
			p.expand = true
			return
		}
	}
}

// findSelection 在 node 上递归查找 startOffset 和 endOffset 选段，选中节点累计到 selected 中。
func (r *VditorRenderer) findSelection(node *Node, startOffset, endOffset int, selected *[]*Node) {
	//for _, rng := range node.ranges {
	//	if rng.startLn > startLn || rng.endLn < endLn {
	//		return
	//	}
	//	if rng.startLn == startLn && (rng.startCol > startCol) {
	//		return
	//	}
	//	if rng.endLn == endLn && (rng.endCol < endCol) {
	//		return
	//	}
	//}

	if nil == node.firstChild {
		// 说明找到了选段内的叶子结点
		*selected = append(*selected, node)
		return
	}

	// 在子节点中递归查找
	for c := node.firstChild; nil != c; c = c.next {
		r.findSelection(c, startOffset, endOffset, selected)
	}
	return
}
