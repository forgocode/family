package article

import (
	"log"
	"os/exec"
	"strconv"

	"github.com/forgocode/family/internal/webservice/controller/article"
	"github.com/forgocode/family/internal/webservice/middleware"
	"github.com/forgocode/family/internal/webservice/router/base"
	"github.com/gin-gonic/gin"
)

type ArticlePlugin struct {
	PluginName  string `json:"name" gorm:"column:name"`
	Md5         string `json:"md5" gorm:"column:md5"`
	Version     string `json:"version" gorm:"column:version"`
	Author      string `json:"author" gorm:"author"`
	Description string `json:"description" gorm:"description"`
	Status      string `json:"status" gorm:"column:status"`
	ExecPath    string `json:"execPath"`
	ListenPort  int32
}

func init() {
	p := &ArticlePlugin{
		PluginName:  "文章服务",
		Version:     "0.0.1_base",
		Author:      "forgocode",
		Description: "用于发布文章",
		ExecPath:    "",
		ListenPort:  10002,
	}
	base.RegisterPlugin(p)
}

func (p *ArticlePlugin) Name() string {
	return p.PluginName
}

func (p *ArticlePlugin) Run() (*exec.Cmd, error) {
	if p.ExecPath == "" {
		return nil, nil
	}
	cmd := exec.Command(p.ExecPath, "-port", strconv.Itoa(int(p.ListenPort)))
	err := cmd.Start()
	if err != nil {
		return nil, err
	}
	go func() {
		err := cmd.Wait()
		if err != nil {
			log.Printf("failed to wait plugin: %+v, err: %+v\n", p, err)
		}
	}()
	return cmd, nil
}

func (p *ArticlePlugin) Router() []base.RouterInfo {
	return []base.RouterInfo{
		{Group: "admin", Path: "/article/publish", Method: "PUT", Handles: []gin.HandlerFunc{article.AdminPublishArticle}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
		{Group: "admin", Path: "/article/ban", Method: "PUT", Handles: []gin.HandlerFunc{article.AdminBanArticle}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
		{Group: "admin", Path: "/article/sendback", Method: "PUT", Handles: []gin.HandlerFunc{article.AdminSendBackArticle}, Middleware: []gin.HandlerFunc{middleware.AuthAdmin()}},
		{Group: "", Path: "/article", Method: "GET", Handles: []gin.HandlerFunc{article.NormalGetArticle}},
		{Group: "", Path: "/article/:id", Method: "GET", Handles: []gin.HandlerFunc{article.NormalGetArticleInfo}},
		{Group: "admin", Path: "/article", Method: "GET", Handles: []gin.HandlerFunc{article.AdminGetArticle}},
	}
}

func (p *ArticlePlugin) Uninstall() {

}

func (p *ArticlePlugin) Upgrade() {

}
