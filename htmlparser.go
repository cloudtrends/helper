package helper

//package dhtmlparser

import (
//"exp/html"
//"fmt"
//"log"
//"strings"
)

type HtmlDocument struct {
}
type LinkTag struct {
	Text string
	Href string
}

type LinkTags []LinkTag

func DGetAllLinks(html_code string) LinkTags {
	var link_tags LinkTags
	/**
		doc, err := html.Parse(strings.NewReader(html_code))
		if err != nil {
			log.Fatal(err)
		}
		var f func(*html.Node)
		f = func(n *html.Node) {
			if n.Type == html.ElementNode {
				if n.Data == "a" {
					for _, a := range n.Attr {
						if a.Key == "href" {
							log.Println(" - - - - - - ")
							log.Println(a.Val)
							//log.Println(a.Namespace)
							var lt LinkTag
							lt.Href = ""
							lt.Text = ""
							link_tags = append(link_tags, lt)
							log.Println(" - - - - - - end")
							break
						}
					}
				} else if n.Data == "div" {
					for x, a := range n.Attr {
						log.Println(x)
						log.Println(a)
						/*
							if a.Key == "href" {
								log.Println(" - - - - - - ")
								log.Println(a.Val)
								//log.Println(a.Namespace)
								var lt LinkTag
								lt.Href = ""
								lt.Text = ""
								link_tags = append(link_tags, lt)
								log.Println(" - - - - - - end")
								break
							}
						* /
	///above is commented
					}
					log.Println(*n)
					log.Println("all in div:", n.Data)
				} else {
					log.Println("element node:", n.Data)
				}
			} else if n.Type == html.TextNode {
				log.Println("textnode:", n.Data)
			} else if n.Type == html.CommentNode {
				log.Println(n.Data)
			} else if n.Type == html.DoctypeNode {
				log.Println(n.Data)
			} else if n.Type == html.ErrorNode {
				log.Println(n.Data)
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				log.Println("begin next sibling ... ")
				f(c)
			}
			log.Println("return from next sibling ")
		}
		f(doc)
	*/
	return link_tags
}

func Test2() {
	/**
	s := `` //`<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul><div>div text content </div>`

	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					log.Println(a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)
	*/
}
func test() {
	/*
		z := html.NewTokenizer(r)
		for {
			tt := z.Next()
			if tt == html.ErrorToken {
				// ...
				return ""
			}
			// Process the current token.
		}
	*/
}

func main() {
	s := `<p>Links:</p><ul><li><a href="foo">FooInnerText</a><li><a href="/bar/baz">BarBazInnerText</a></ul><div id ="id_a" class="cls_a"><a href="aa">aaa in div </a> div text content </div>`
	DGetAllLinks(s)
}
