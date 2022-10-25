go mod的使用：

    初始化一个 go.mod文件
        go mod init github.com/luwang-epic/study_go

    使用tag，进行版本控制
        git tag v1.0.0
        git push --tags

        推荐在这个状态下，再切出一个分支，用于后续v1.0.0的修复推送,不要直接在master分支修复

    如果是大版本更新，可以修改go.mod文件为（如改为v2版本）：
        module github.com/luwang-epic/study_go/v2
        这样就不会影响之前的了，使用时需要import mv2 "github.com/jacksonyoudi/gomodone/v2"

    如果有人需要使用，就可以使用
        // 下载最新版本
        go get github.com/luwang-epic/study_go
        // 指定包，指定版本
        go get github.com/luwang-epic/study_go@v1.0.1