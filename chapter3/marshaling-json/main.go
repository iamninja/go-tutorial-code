package main

import (
	"fmt"
	"encoding/json"
)

// Data -> JSON: Marshaling (serializing)
// JSON -> Data: Unmarshaling

type Article struct {
	Name		string
	AuthorName	string
	draft		bool
}

type Product struct{
	// Field appears in JSON with key "name"
	Name string `json:"name"`

	// Field appears in JSON with key "author_name"
	// Does not appear if its value is empty
	AuthorName string `json:"author_name",omitempty`

	// Exported field but omitted in JSON
	CommissionPrice float64 `json:"-"`
}

type ArticleCollection struct {
	Articles 	[]Article 	`json:"articles"`
	Total		int 		`json:"total"`
}

func main() {
	article := Article{
		Name:		"JSON in Go",
		AuthorName:	"Mal Curtis",
		draft:		true,
	}
	data, err := json.Marshal(article)		// Convert only exported fields in a struct. draft will be omitted
	if err != nil {
		fmt.Println("Couldn't marshal article:", err)
	}else{
		fmt.Println(string(data))
	}

	//
	data, err = json.MarshalIndent(article, "", "  ")	//(data, prefix, indantation)
	if err != nil {
		fmt.Println("Couldn't marshal article:", err)
	}else{
		fmt.Println(string(data))
	}

	//
	product := Product{
		Name:				"A Study in Go",
		AuthorName:			"Arthur Conan Doyle, Sir",
		CommissionPrice:	23.4,
	}
	data, err = json.MarshalIndent(product, "", "  ")	//(data, prefix, indantation)
	if err != nil {
		fmt.Println("Couldn't marshal product:", err)
	}else{
		fmt.Println(string(data))
	}

	//
	article2 := Article{
		Name:				"A Study in Go",
		AuthorName:			"Arthur Conan Doyle, Sir",
	}
	articles := []Article{article, article2}
	collection := ArticleCollection{
		Articles:	articles,
		Total:		len(articles),
	}
	data, err = json.MarshalIndent(collection, "", "  ")
	if err != nil { panic(err) }
	fmt.Println(string(data))
}