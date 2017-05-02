package templating

type TemplatesList []Template

type Template struct {
    //
    //
    filePath               string
    fileExtension          string
    //
    sectionTemplate        string
    lineTemplate           string
    lineTemplateDataSize   int
    //
    bindToExtension        bool
    bindToFilePath         bool
    //
    //
}


func(tl *TemplatesList)Append(t *Template)(err error) {
    //
    //
    //
    return

}

func(tl *TemplatesList)TemplateForThisExtensionExists(ext string)(err error) {
    //
    for i := range (*tl) {
        template := (*tl)[i]
        _ = template
    }
    //
    return
}

func(tl *TemplatesList)TemplateForThisFilePathExists(path string)(err error) {
    //
    for i := range (*tl) {
        template := (*tl)[i]
        _ = template
    }
    //
    return
}
