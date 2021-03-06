package html

import (
	"bytes"
	golang_html "html"
	"io"
)

type Node struct {
	nodeType         string
	Attributes       map[string]string
	children         []Node
	text             string
	headTagMetaMagic string // This is for html doctype
}

func (node Node) attributesAsString() string {
	var result bytes.Buffer
	for key, value := range node.Attributes {
		//Note this is done for performance
		result.WriteString(" ")
		result.WriteString(key)
		result.WriteString("=\"")
		result.WriteString(value)
		result.WriteString("\"")
	}
	return result.String()
}

func (node Node) WriteTo(writer io.Writer) {
	io.WriteString(writer, node.String())
}

func (node Node) String() string {
	buffer := &bytes.Buffer{}
	node.writeToBuffer(buffer)
	return buffer.String()
}

func (node Node) writeToBuffer(buffer *bytes.Buffer) {

	//TODO peformance bench this 'if' vs 'if len' vs no if
	if node.headTagMetaMagic != "" {
		buffer.WriteString(node.headTagMetaMagic)
	}
	buffer.WriteString("<")
	buffer.WriteString(node.nodeType)
	buffer.WriteString(node.attributesAsString())
	buffer.WriteString(">")

	if len(node.children) > 0 {
		for i := range node.children {
			node.children[i].writeToBuffer(buffer)
		}
	} else {
		buffer.WriteString(node.text)
	}

	buffer.WriteString("</")
	buffer.WriteString(node.nodeType)
	buffer.WriteString(">")

}

// Text will excape any input as html
func (node Node) Text(text string) Node {
	node.text = golang_html.EscapeString(text)
	return node
}

// Same as Text() but doesn't escape html
func (node Node) TextUnsafe(text string) Node {
	node.text = text
	return node
}

func (node *Node) Append(children ...Node) {
	node.children = append(node.children, children...)
}

func (node Node) Children(children ...Node) Node {
	node.children = append(node.children, children...)
	return node
}

func (node Node) Attribute(attributeName, value string) Node {
	if node.Attributes == nil {
		node.Attributes = make(map[string]string)
	}

	node.Attributes[attributeName] = value
	return node
}
