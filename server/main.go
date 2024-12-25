package main

import (
	"gin-vben-admin/crontab"
	"gin-vben-admin/dao"
	"gin-vben-admin/dao/migration"
	"gin-vben-admin/dao/query"
	"gin-vben-admin/dao/repo"
	"gin-vben-admin/global"
	"gin-vben-admin/pkg/avatar"
	"gin-vben-admin/pkg/cache"
	"gin-vben-admin/pkg/conf"
	"gin-vben-admin/pkg/constant"
	"gin-vben-admin/pkg/lock"
	"gin-vben-admin/pkg/shutdown"
	"gin-vben-admin/pkg/utils"
	"gin-vben-admin/pkg/validators"
	"gin-vben-admin/routers"
	"gin-vben-admin/tasks"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"time"
)

func init() {
	global.Viper, global.Conf = conf.Parse()

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	path := utils.RelativePath("logs/gin-vben-admin.log")
	writer, _ := rotatelogs.New(
		path+"%Y%m%d",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(24*120)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
	)

	if global.Conf.System.Env == constant.EnvProd {
		logrus.SetOutput(writer)
		gin.SetMode(gin.ReleaseMode)
		logrus.SetLevel(logrus.InfoLevel)
	}
	conf.InitLocalPath()
	cache.Init()
	dao.Init()
	query.SetDefault(global.DB)
	fontByte, _ := f.ReadFile("font/PingFangBold.ttf")
	global.Avatar = avatar.Init(fontByte)
	//执行迁移
	migration.Run()
	lock.Init()
	tasks.Server()
	tasks.Client()
	crontab.Init()
	validators.Init()
	repo.CasbinSrv.Casbin()
}

//	@title		ApiDoc 1.0.0
//	@version	1.0.0
//	@description.markdown
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swag.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@securityDefinitions.apiKey	AuthToken
//	@in							header
//	@name						x-token

//	@BasePath	/api/

// @Accept		x-www-form-urlencoded
// @Produce	json
func main() {
	api := routers.InitRouter()
	go func() {
		logrus.Infof("开始监听 %s", global.Conf.System.Listen)
		logrus.Error(api.Run(global.Conf.System.Listen))
	}()
	shutdown.NewHook().Close(func() {
		tasks.ServerClose()
		tasks.ClientClose()
		crontab.Close()
		logrus.Info("close all service")
	})

}
