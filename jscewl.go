package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)


var stripOff = []string{"align-content", "align-items", "align-self", "all", "animation", "animation-delay", "animation-direction", "animation-duration", "animation-fill-mode", "animation-iteration-count", "animation-name", "animation-play-state", "animation-timing-function", "backface-visibility", "background", "background-attachment", "background-blend-mode", "background-clip", "background-color", "background-image", "background-origin", "background-position", "background-repeat", "background-size", "border", "border-bottom", "border-bottom-color", "border-bottom-left-radius", "border-bottom-right-radius", "border-bottom-style", "border-bottom-width", "border-collapse", "border-color", "border-image", "border-image-outset", "border-image-repeat", "border-image-slice", "border-image-source", "border-image-width", "border-left", "border-left-color", "border-left-style", "border-left-width", "border-radius", "border-right", "border-right-color", "border-right-style", "border-right-width", "border-spacing", "border-style", "border-top", "border-top-color", "border-top-left-radius", "border-top-right-radius", "border-top-style", "border-top-width", "border-width", "bottom", "box-decoration-break", "box-shadow", "box-sizing", "break-after", "break-before", "break-inside", "caption-side", "caret-color", "@charset", "clear", "clip", "color", "column-count", "column-fill", "column-gap", "column-rule", "column-rule-color", "column-rule-style", "column-rule-width", "column-span", "column-width", "columns", "content", "counter-increment", "counter-reset", "cursor", "direction", "display", "empty-cells", "filter", "flex", "flex-basis", "flex-direction", "flex-flow", "flex-grow", "flex-shrink", "flex-wrap", "float", "font", "@font-face", "font-family", "font-feature-settings", "@font-feature-values", "font-kerning", "font-language-override", "font-size", "font-size-adjust", "font-stretch", "font-style", "font-synthesis", "font-variant", "font-variant-alternates", "font-variant-caps", "font-variant-east-asian", "font-variant-ligatures", "font-variant-numeric", "font-variant-position", "font-weight", "grid", "grid-area", "grid-auto-columns", "grid-auto-flow", "grid-auto-rows", "grid-column", "grid-column-end", "grid-column-gap", "grid-column-start", "grid-gap", "grid-row", "grid-row-end", "grid-row-gap", "grid-row-start", "grid-template", "grid-template-areas", "grid-template-columns", "grid-template-rows", "hanging-punctuation", "height", "hyphens", "image-rendering", "@import", "isolation", "justify-content", "@keyframes", "left", "letter-spacing", "line-break", "line-height", "list-style", "list-style-image", "list-style-position", "list-style-type", "margin", "margin-bottom", "margin-left", "margin-right", "margin-top", "max-height", "max-width", "@media", "min-height", "min-width", "mix-blend-mode", "object-fit", "object-position", "opacity", "order", "orphans", "outline", "outline-color", "outline-offset", "outline-style", "outline-width", "overflow-wrap", "overflow-x", "overflow-y", "padding", "padding-bottom", "padding-left", "padding-right", "padding-top", "page-break-after", "page-break-before", "page-break-inside", "perspective", "perspective-origin", "pointer-events", "position", "quotes", "resize", "right", "scroll-behavior", "tab-size", "table-layout", "text-align", "text-align-last", "text-combine-upright", "text-decoration", "text-decoration-color", "text-decoration-line", "text-decoration-style", "text-indent", "text-justify", "text-orientation", "text-overflow", "text-shadow", "text-transform", "text-underline-position", "top", "transform", "transform-origin", "transform-style", "transition", "transition-delay", "transition-duration", "transition-property", "transition-timing-function", "unicode-bidi", "user-select", "vertical-align", "visibility", "white-space", "widows", "width", "word-break", "word-spacing", "word-wrap", "writing-mode", "z-index", "await", "break", "case", "catch", "class", "const", "continue", "debugger", "default", "delete", "do", "else", "enum", "export", "extends", "false", "finally", "for", "function", "if", "implements", "import", "in", "instanceof", "interface", "let", "new", "null", "package", "private", "protected", "public", "return", "super", "switch", "static", "this", "throw", "try", "true", "typeof", "var", "void", "while", "with", "abstract", "else", "instanceof", "super", "boolean", "enum", "int", "switch", "break", "export", "interface", "synchronized", "byte", "extends", "let", "this", "case", "false", "long", "throw", "catch", "final", "native", "throws", "char", "finally", "new", "transient", "class", "float", "null", "true", "const", "for", "package", "try", "continue", "function", "private", "typeof", "debugger", "goto", "protected", "var", "default", "if", "public", "void", "delete", "implements", "return", "volatile", "do", "import", "short", "while", "double", "in", "static", "with", "alert", "frames", "outerheight", "all", "framerate", "outerwidth", "anchor", "function", "packages", "anchors", "getclass", "pagexoffset", "area", "hasownproperty", "pageyoffset", "array", "hidden", "parent", "assign", "history", "parsefloat", "blur", "image", "parseint", "button", "images", "password", "checkbox", "infinity", "pkcs11", "clearinterval", "isfinite", "plugin", "cleartimeout", "isnan", "prompt", "clientinformation", "isprototypeof", "propertyisenum", "close", "java", "prototype", "closed", "javaarray", "radio", "confirm", "javaclass", "reset", "constructor", "javaobject", "screenx", "crypto", "javapackage", "screeny", "date", "innerheight", "scroll", "decodeuri", "innerwidth", "secure", "decodeuricomponent", "layer", "select", "defaultstatus", "layers", "self", "document", "length", "setinterval", "element", "link", "settimeout", "elements", "location", "status", "embed", "math", "string", "embeds", "mimetypes", "submit", "encodeuri", "name", "taint", "encodeuricomponent", "nan", "text", "escape", "navigate", "textarea", "eval", "navigator", "top", "event", "number", "tostring", "fileupload", "object", "undefined", "focus", "offscreenbuffering", "unescape", "form", "open", "untaint", "forms", "opener", "valueof", "frame", "option", "window", "yield", "a", "abbr", "acronym", "address", "applet", "area", "article", "aside", "audio", "b", "base", "basefont", "bdi", "bdo", "big", "body", "br", "button", "canvas", "caption", "center", "cite", "code", "col", "colgroup", "data", "datalist", "dd", "del", "details", "dfn", "dialog", "dir", "div", "dl", "dt", "em", "embed", "fieldset", "figure", "font", "footer", "form", "frame", "frameset", "h1", "head", "header", "hr", "html", "i", "iframe", "img", "input", "ins", "kbd", "label", "legend", "li", "link", "main", "map", "mark", "meta", "meter", "nav", "noframes", "noscript", "object", "ol", "optgroup", "option", "output", "p", "param", "picture", "pre", "progress", "q", "rp", "rt", "ruby", "s", "samp", "script", "section", "select", "small", "source", "span", "strike", "strong", "style", "sub", "summary", "sup", "svg", "table", "tbody", "td", "template", "textarea", "tfoot", "th", "thead", "time", "title", "tr", "track", "tt", "u", "ul", "var", "video", "wbr", "input", "form", "area", "img", "script", "meta", "input", "textarea", "td", "th", "meta", "area", "object", "del", "track", "script", "input", "a", "form", "label", "button", "td", "th", "canvas", "meter", "a", "a", "meta", "img", "track", "track", "input", "audio", "meter", "input", "a", "form", "input", "video", "button", "form", "audio", "audio", "audio", "body", "body", "body", "body", "body", "audio", "audio", "body", "form", "body", "body", "form", "body", "details", "meter", "input", "video", "audio", "a", "area", "ol", "textarea", "td", "th", "iframe", "th", "area", "input", "img", "col", "audio", "iframe", "track", "img", "ol", "input", "a", "a", "img", "button", "canvas", "textarea"}


func main() {

	stdio := bufio.NewScanner(os.Stdin)
	for stdio.Scan() {
		fetchJsFile(stdio.Text())
	}
}


func fetchJsFile(JsFile string) {
	
	if strings.Contains(fmt.Sprintf("%v", JsFile), "://") == false || strings.Contains(fmt.Sprintf("%v", JsFile), ".js") == false {
		fmt.Printf("Bad url: %v \nMake sure the url points to a Js file", JsFile)
	} else {
		uri := JsFile
		resp, err := http.Get(uri)
		if err != nil {
			print(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			print(err)
		}
		wordWorker(string(body))
	}
}


func wordWorker(content string) {
	
	var bWords []string

	re, _ := regexp.Compile(`[a-zA-Z0-9_\-.]+`)
	var regexContent = re.FindAllString(content, -1)

	for _, word := range regexContent {
		
		if strings.Contains(word, ".") {
			filteredWord := strings.Split(word, ".")
			
			for _, separated := range filteredWord {
				if stringInSlice(separated, bWords) == false {
					bWords = append(bWords, separated)
				}
			}
		} else {
			bWords = append(bWords, word)
		}
	}
	
	var anew []string
	for _, _word := range bWords {
		effWord := strings.ToLower(_word)
		if len(effWord) > 1 {
			if stringInSlice(effWord, stripOff) == false {
				if stringInSlice(effWord, anew) == false {
					anew = append(anew, effWord)
				}
			}
		}
	}
	for _, w := range anew {
		fmt.Println(w)
	}
}


func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}












