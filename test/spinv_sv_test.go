// Lute - 一款对中文语境优化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package test

import (
	"testing"

	"github.com/88250/lute"
)

var spinVditorSVDOMTests = []*parseTest{

	{"51", "> 1. foo\n> 2. ```\n>    bar\n>    ‸\n>    ```", "<span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"li-marker\" class=\"vditor-sv__marker\">1. </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"li-marker\" class=\"vditor-sv__marker\">2. </span><span data-type=\"code-block-open-marker\" class=\"vditor-sv__marker\">```</span><span class=\"vditor-sv__marker--info\" data-type=\"code-block-info\"></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"padding\">   </span><span data-type=\"text\">bar<span data-type=\"padding\"></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"padding\">   </span><wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"padding\">   </span><span data-type=\"code-block-close-marker\" class=\"vditor-sv__marker\">```</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"50", "[foo]\n\n[foo]: bar‸", "<span class=\"vditor-sv__marker--bracket\">[</span><span>foo</span><span class=\"vditor-sv__marker--bracket\">]</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span class=\"vditor-sv__marker--bracket\">[</span><span class=\"vditor-sv__marker--link\" data-type=\"footnotes-link\">foo</span><span class=\"vditor-sv__marker--bracket\">]</span><span>: </span>bar<wbr><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"49", "![foo](bar)‸", "<span class=\"vditor-sv__marker\">!</span><span class=\"vditor-sv__marker--bracket\">[</span><span class=\"vditor-sv__marker--bracket\">foo</span><span class=\"vditor-sv__marker--bracket\">]</span><span class=\"vditor-sv__marker--paren\">(</span><span class=\"vditor-sv__marker--link\">bar</span><span class=\"vditor-sv__marker--paren\">)</span><span data-type=\"text\"><wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"48", "![foo](bar)‸", "<span class=\"vditor-sv__marker\">!</span><span class=\"vditor-sv__marker--bracket\">[</span><span class=\"vditor-sv__marker--bracket\">foo</span><span class=\"vditor-sv__marker--bracket\">]</span><span class=\"vditor-sv__marker--paren\">(</span><span class=\"vditor-sv__marker--link\">bar</span><span class=\"vditor-sv__marker--paren\">)</span><span data-type=\"text\"><wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"47", "> \n‸", "<span data-type=\"text\">&gt;</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"text\"><wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"46", "foo\n\n## bar‸", "<span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span class=\"vditor-sv__marker--heading h2\" data-type=\"heading-marker\">## </span><span data-type=\"text\" class=\"h2\">bar<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"45", "**foo**\n\n*bar‸*", "<span class=\"vditor-sv__marker--bi strong\">**</span><span data-type=\"text\" class=\"strong\">foo</span><span class=\"vditor-sv__marker--bi strong\">**</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span class=\"vditor-sv__marker--bi em\">*</span><span data-type=\"text\" class=\"em\">bar<wbr></span><span class=\"vditor-sv__marker--bi em\">*</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"44", "> foo\n> \n> ```\n> bar\n> ‸\n> ```", "<span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"code-block-open-marker\" class=\"vditor-sv__marker\">```</span><span class=\"vditor-sv__marker--info\" data-type=\"code-block-info\"></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"text\">bar<span data-type=\"padding\"></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"code-block-close-marker\" class=\"vditor-sv__marker\">```</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"43", "* foo\n  * bar\n* b‸", "<span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"padding\">  </span><span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"text\">bar</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"text\">b<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"42", "* foo\n  1. ba‸", "<span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"padding\">  </span><span data-type=\"li-marker\" class=\"vditor-sv__marker\">1. </span><span data-type=\"text\">ba<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"41", "> * foo\n> \n>   bar‸\n", "<span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"padding\">  </span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"padding\">  </span><span data-type=\"text\">bar<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"40", "> * foo\n> * bar‸", "<span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"text\">bar<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"39", "    f‸", "<span data-type=\"code-block-open-marker\" class=\"vditor-sv__marker\">```</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"text\">f<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span class=\"vditor-sv__marker--info\" data-type=\"code-block-info\">```</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"38", "```  `code`  ```‸", "<span class=\"vditor-sv__marker\">```</span><span> `code` </span><span class=\"vditor-sv__marker\">```</span><span data-type=\"text\"><wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"37", "`` `code` ``‸", "<span class=\"vditor-sv__marker\">`` </span><span>`code`</span><span class=\"vditor-sv__marker\"> ``</span><span data-type=\"text\"><wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"36", "# foo {id}b‸", "<span class=\"vditor-sv__marker--heading h1\" data-type=\"heading-marker\"># </span><span data-type=\"text\" class=\"h1\">foo {id}b<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"35", "* foo\n  ```‸", "<span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"padding\">  </span><span data-type=\"code-block-open-marker\" class=\"vditor-sv__marker\">```</span><span class=\"vditor-sv__marker--info\" data-type=\"code-block-info\"><wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"padding\">  </span><span data-type=\"text\"></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"padding\">  </span><span data-type=\"code-block-close-marker\" class=\"vditor-sv__marker\">```</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"34", "> foo\n> \n> > bar\n> \n> baz‸", "<span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"text\">bar</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"text\">baz<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"33", "foo\n> bar‸", "<span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"text\">bar<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"32", "# foo {id‸}", "<span class=\"vditor-sv__marker--heading h1\" data-type=\"heading-marker\"># </span><span data-type=\"text\" class=\"h1\">foo</span><span class=\"vditor-sv__marker h1\"> {id<wbr>}</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"31", "# foo {id}‸", "<span class=\"vditor-sv__marker--heading h1\" data-type=\"heading-marker\"># </span><span data-type=\"text\" class=\"h1\">foo</span><span class=\"vditor-sv__marker h1\"> {id}</span><span data-type=\"text\" class=\"h1\"><wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"30", "```\nfoo\n```‸", "<span data-type=\"code-block-open-marker\" class=\"vditor-sv__marker\">```</span><span class=\"vditor-sv__marker--info\" data-type=\"code-block-info\"></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"text\">foo<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"code-block-close-marker\" class=\"vditor-sv__marker\">```</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"29", "> foo\n> >‸\n", "<span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"text\">&gt;<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"28", "> >‸", "<span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"text\">&gt;<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"27", "这里是一个脚注引用[^1]，这里是另一个脚注引用[^bignote]。\n\n[^1]: 第一个脚注定义。\n[^bignote]: 脚注定义可使用多段内容。\n\n    缩进对齐的段落包含在这个脚注定义内。‸\n", "<span data-type=\"text\">这里是一个脚注引用</span><span class=\"sup\"><span class=\"vditor-sv__marker--bracket\">[</span><span class=\"vditor-sv__marker--link\">^1</span><span class=\"vditor-sv__marker--bracket\">]</span></span><span data-type=\"text\">，这里是另一个脚注引用</span><span class=\"sup\"><span class=\"vditor-sv__marker--bracket\">[</span><span class=\"vditor-sv__marker--link\">^bignote</span><span class=\"vditor-sv__marker--bracket\">]</span></span><span data-type=\"text\">。</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span class=\"vditor-sv__marker--bracket\">[</span><span class=\"vditor-sv__marker--link\" data-type=\"footnotes-link\">^1</span><span class=\"vditor-sv__marker--bracket\">]</span><span>: </span><span data-type=\"text\">第一个脚注定义。</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span class=\"vditor-sv__marker--bracket\">[</span><span class=\"vditor-sv__marker--link\" data-type=\"footnotes-link\">^bignote</span><span class=\"vditor-sv__marker--bracket\">]</span><span>: </span><span data-type=\"text\">脚注定义可使用多段内容。</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"padding\">    </span><span data-type=\"text\">缩进对齐的段落包含在这个脚注定义内。<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"26", "|foo|bar|\n| ---| ---|\n|‸", "<span data-type=\"table\">|foo|bar|\n| ---| ---|\n|<wbr><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span></span>"},
	{"25", "‸", "<span data-type=\"text\"><wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"24", "> * f‸\n> \n", "<span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"text\">f<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"23", "> ## foo\n>  \n> * b‸", "<span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span class=\"vditor-sv__marker--heading h2\" data-type=\"heading-marker\">## </span><span data-type=\"text\" class=\"h2\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"text\">b<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"22", "> foo\n> \n> b‸", "<span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"text\">b<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"21", "> * foo\n>‸", "<span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"padding\">  </span><span data-type=\"text\"><wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"20", "* foo\n\n  bar\n\n* baz\n\n  b‸", "<span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"padding\">  </span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"padding\">  </span><span data-type=\"text\">bar</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"text\">baz</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"padding\">  </span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"padding\">  </span><span data-type=\"text\">b<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"19", "* foo\n\n  bar‸", "<span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"padding\">  </span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"padding\">  </span><span data-type=\"text\">bar<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"18", "* foo\n\n‸\n", "<span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"text\"><wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"17", "*‸", "<span data-type=\"text\">*<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"16", "> * foo\n> \n>‸", "<span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"text\"><wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"15", "> foo\n> # b‸", "<span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span class=\"vditor-sv__marker--heading h1\" data-type=\"heading-marker\"># </span><span data-type=\"text\" class=\"h1\">b<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"14", "* foo\n\n‸bar", "<span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"text\"><wbr>bar</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"13", "* foo\n  ```\n  bar\n  baz‸\n  ```\n", "<span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"padding\">  </span><span data-type=\"code-block-open-marker\" class=\"vditor-sv__marker\">```</span><span class=\"vditor-sv__marker--info\" data-type=\"code-block-info\"></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"padding\">  </span><span data-type=\"text\">bar<span data-type=\"padding\"></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"padding\">  </span>baz<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"padding\">  </span><span data-type=\"code-block-close-marker\" class=\"vditor-sv__marker\">```</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"12", "1. foo\n   1. bar\n2. ‸", "<span data-type=\"li-marker\" class=\"vditor-sv__marker\">1. </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"padding\">   </span><span data-type=\"li-marker\" class=\"vditor-sv__marker\">1. </span><span data-type=\"text\">bar</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"li-marker\" class=\"vditor-sv__marker\">2. </span><span data-type=\"text\"><wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"11", "* foo\n\n  bar\nbaz‸", "<span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"padding\">  </span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"padding\">  </span><span data-type=\"text\">bar</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"padding\">  </span><span data-type=\"text\">baz<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"10", "* foo\n  * ba‸", "<span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"padding\">  </span><span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"text\">ba<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"9", "* foo\n‸", "<span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"padding\">  </span><span data-type=\"text\"><wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"8", "> foo\n> \n> bar‸", "<span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"text\">bar<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"7", "* foo\n  >‸", "<span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"padding\">  </span><span data-type=\"text\">&gt;<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"6", "* foo‸", "<span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"text\">foo<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"5", "> * foo\n>   > bar\n>   > b‸  ", "<span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"padding\">  </span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"text\">bar</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"padding\">  </span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"text\">b<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"4", "> * foo\n>   > b‸", "<span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"padding\">  </span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"text\">b<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"3", "> * foo\n> * bar‸", "<span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"text\">foo</span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"text\">bar<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"2", "> * foo‸", "<span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"li-marker\" class=\"vditor-sv__marker\">* </span><span data-type=\"text\">foo<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"1", "> foo‸", "<span data-type=\"blockquote-marker\" class=\"vditor-sv__marker\">&gt; </span><span data-type=\"text\">foo<wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
	{"0", "*foo*‸", "<span class=\"vditor-sv__marker--bi em\">*</span><span data-type=\"text\" class=\"em\">foo</span><span class=\"vditor-sv__marker--bi em\">*</span><span data-type=\"text\"><wbr></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span><span data-type=\"newline\"><br /><span style=\"display: none\">\n</span></span>"},
}

func TestSpinVditorSVDOM(t *testing.T) {
	luteEngine := lute.New()
	luteEngine.ToC = true

	for _, test := range spinVditorSVDOMTests {
		html := luteEngine.SpinVditorSVDOM(test.from)
		if test.to != html {
			t.Fatalf("test case [%s] failed\nexpected\n\t%q\ngot\n\t%q\noriginal html\n\t%q", test.name, test.to, html, test.from)
		}
	}
}
