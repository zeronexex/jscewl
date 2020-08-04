package main

import (
	"fmt"
	"regexp"
	"strings"
)

var blacklist = []string{ "await", "break", "case", "catch", "class", "const", "continue", "debugger", "default", "delete", "do", "else", "enum", "export", "extends", "false", "finally", "for", "function", "if", "implements", "import", "in", "instanceof", "interface", "let", "new", "null", "package", "private", "protected", "public", "return", "super", "switch", "static", "this", "throw", "try", "true", "typeof", "var", "void", "while", "with", "abstract", "else", "instanceof", "super", "boolean", "enum", "int", "switch", "break", "export", "interface", "synchronized", "byte", "extends", "let", "this", "case", "false", "long", "throw", "catch", "final", "native", "throws", "char", "finally", "new", "transient", "class", "float", "null", "true", "const", "for", "package", "try", "continue", "function", "private", "typeof", "debugger", "goto", "protected", "var", "default", "if", "public", "void", "delete", "implements", "return", "volatile", "do", "import", "short", "while", "double", "in", "static", "with", "alert", "frames", "outerheight", "all", "framerate", "outerwidth", "anchor", "function", "packages", "anchors", "getclass", "pagexoffset", "area", "hasownproperty", "pageyoffset", "array", "hidden", "parent", "assign", "history", "parsefloat", "blur", "image", "parseint", "button", "images", "password", "checkbox", "infinity", "pkcs11", "clearinterval", "isfinite", "plugin", "cleartimeout", "isnan", "prompt", "clientinformation", "isprototypeof", "propertyisenum", "close", "java", "prototype", "closed", "javaarray", "radio", "confirm", "javaclass", "reset", "constructor", "javaobject", "screenx", "crypto", "javapackage", "screeny", "date", "innerheight", "scroll", "decodeuri", "innerwidth", "secure", "decodeuricomponent", "layer", "select", "defaultstatus", "layers", "self", "document", "length", "setinterval", "element", "link", "settimeout", "elements", "location", "status", "embed", "math", "string", "embeds", "mimetypes", "submit", "encodeuri", "name", "taint", "encodeuricomponent", "nan", "text", "escape", "navigate", "textarea", "eval", "navigator", "top", "event", "number", "tostring", "fileupload", "object", "undefined", "focus", "offscreenbuffering", "unescape", "form", "open", "untaint", "forms", "opener", "valueof", "frame", "option", "window", "yield"}



func main() {
	fmt.Println(getJsWords(jsF))
}


func getJsWords(content string) []string {
	// This function will filter the words accordingly
	var allWords []string

	/* getting js words from file with regex */
	re := regexp.MustCompile(`[a-zA-Z0-9_\-.]+`)
	var regexContent = re.FindAllString(content, -1)


	for _, word := range regexContent {

		/* blacklisting and filtering */
		if len(word) > 1 {

			/* taking action if words of filtered regex contains "." */
			if strings.Contains(word, ".") {
				filteredWord := strings.Split(word, ".")

				/* split words turn into one array thus iterating those */
				for _, separated := range filteredWord {
					if stringInSlice(separated, allWords) == false {
						allWords = append(allWords, separated)
					}
				}
			}
		}
	}
	var FallWords []string
	for _, word_ := range allWords {
		if stringInSlice(word_, blacklist) == false {
			FallWords = append(FallWords, word_)
		}
	}
	return FallWords
}


func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
