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

	fmt.Printf("Tag,API,Path,Method\n")
	for k, v := range doc.Spec().Paths.Paths {
		tag := "no Tag"

		if v.Get != nil {
			if v.Get.Tags != nil {
				tag = v.Get.Tags[0]
			}
			fmt.Printf("%s,%s,%s,GET\n", tag, k, v.Get.Summary)
		}
		if v.Post != nil {
			if v.Post.Tags != nil {
				tag = v.Post.Tags[0]
			}
			fmt.Printf("%s,%s,%s,POST\n", tag, k, v.Post.Summary)
		}
		if v.Put != nil {
			if v.Put.Tags != nil {
				tag = v.Put.Tags[0]
			}
			fmt.Printf("%s,%s,%s,PUT\n", tag, k, v.Put.Summary)
		}
		if v.Delete != nil {
			if v.Delete.Tags != nil {
				tag = v.Delete.Tags[0]
			}
			fmt.Printf("%s,%s,%s,DELETE\n", tag, k, v.Delete.Summary)
		}
	}
}
