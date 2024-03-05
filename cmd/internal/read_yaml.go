package internal

import (
	"fmt"

	"github.com/go-openapi/loads"
)

func ReadYaml(path string) {
	// OpenAPIペックのファイルを読み込む
	doc, err := loads.Spec(path)
	if err != nil {
		fmt.Println("Could not load this spec")
		return
	}

	fmt.Println("Spec : ", doc.Spec().Info.Title)
	fmt.Println("Spec : ", doc.Spec().Info.Description)
	fmt.Println("Spec : ", doc.Spec().Info.Version)

	fmt.Println()

	fmt.Printf("API,Path,Method\n")
	for k, v := range doc.Spec().Paths.Paths {
		if v.Get != nil {
			fmt.Printf("%s,%s,GET\n", k, v.Get.Summary)
		}
		if v.Post != nil {
			fmt.Printf("%s,%s,POST\n", k, v.Post.Summary)
		}
		if v.Put != nil {
			fmt.Printf("%s,%s,PUT\n", k, v.Put.Summary)
		}
		if v.Delete != nil {
			fmt.Printf("%s,%s,DELETE\n", k, v.Delete.Summary)
		}
	}
}
