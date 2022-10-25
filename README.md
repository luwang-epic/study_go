go mod的使用：

    初始化一个 go.mod文件
        go mod init github.com/luwang-epic/study_go

    使用tag，进行版本控制
        git tag v1.0.0
        git push --tags

        推荐在这个状态下，再切出一个分支，用于后续v1.0.0的修复推送,不要直接在master分支修复

    

    如果有人需要使用，就可以使用
        go get github.com/luwang-epic/study_go