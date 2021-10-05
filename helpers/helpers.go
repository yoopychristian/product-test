package helpers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"time"

	cfg "product-test/config"
	fx "product-test/functions"
	adt "product-test/repo-adaptor"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

const (
	WARN = iota
	ERROR
	DEBUG
)

type HTTPResponse struct {
	Status      bool   `json:"status"`
	ErrorCode   string `json:"error_code"`
	Description string `json:"description"`
	Data        string `json:"data"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func NewRepositoryContext(rr io.Reader, rt http.RoundTripper) (cfg.RepositoryContext, error) {
	//error handle
	handleErr := func(err error) (cfg.RepositoryContext, error) {
		return cfg.RepositoryContext{}, err
	}

	//get value from env vars
	config, err := cfg.GetRepositoryConfiguration()
	if err != nil {
		return handleErr(err)
	}

	//init log file
	file, err := os.Create(config.App.LogPath + config.App.Name + time.Now().Format("2006_01_02__15_04") + ".log")
	if err != nil {
		return handleErr(err)
	}

	var l *zap.Logger
	if config.App.Debug {
		l = fx.LogInit(true, file)
	} else {
		l = fx.LogInit(false, file)
	}
	defer l.Sync()

	//setup http client
	httpclient := adt.HttpClient{
		Client: &http.Client{
			Transport: rt,
			Timeout:   config.App.NetTimeOut,
		},
	}

	//init adaptor
	adaptor := adt.RepositoryAdaptor{
		Client: httpclient,
		URL: adt.RepositoryURL{
			Read:  fx.EnvString("REPOSITORY_URL_READ"),
			Write: fx.EnvString("REPOSITORY_URL_WRITE"),
		},
	}

	//init db
	dbCfg := fx.DBParam{
		Host:     config.DB.Host,
		Port:     config.DB.Port,
		Name:     config.DB.DBName,
		Schema:   config.DB.Schema,
		User:     config.DB.Username,
		Password: config.DB.Password,
		AppName:  config.App.Name,
		Timeout:  config.DB.ConnectTimeOut,
		MaxOpen:  config.DB.MaxOpenConn,
		MaxIdle:  config.DB.MaxIdleConn,
		Logging:  config.DB.Logging,
	}
	db, err := fx.DBInit(dbCfg, l, config.App.LogPath+"BODB_", dbCfg.Logging)
	if err != nil {
		return handleErr(err)
	}

	//return service context
	return cfg.RepositoryContext{
		Config:  config,
		Log:     l,
		Adaptor: adaptor,
		DB:      db,
	}, nil
}

type RespParams struct {
	Log      *zap.Logger
	Context  *gin.Context
	Severity int
	URL      string
	Section  string
	Reason   string
	Error    error
	Input    interface{}
}

func BadResponseExist(c *gin.Context, reason string) {
	response := HTTPResponse{
		Status:      false,
		Description: reason,
	}
	c.JSON(http.StatusBadRequest, response)
}

func BadResponse(rp RespParams) {
	switch rp.Severity {
	case DEBUG:
		rp.Log.Debug(rp.Section,
			zap.String("connection", rp.URL),
			zap.Any("parameters", rp.Input),
			zap.String("description", rp.Reason),
			zap.Error(rp.Error))
	case WARN:
		rp.Log.Warn(rp.Section,
			zap.String("connection", rp.URL),
			zap.Any("parameters", rp.Input),
			zap.String("description", rp.Reason),
			zap.Error(rp.Error))
	case ERROR:
		rp.Log.Error(rp.Section,
			zap.String("connection", rp.URL),
			zap.Any("parameters", rp.Input),
			zap.String("description", rp.Reason),
			zap.Error(rp.Error))
	}
	response := HTTPResponse{
		Status:      false,
		Description: rp.Reason,
	}

	rp.Context.JSON(http.StatusBadRequest, response)
}

func RepoBadResponse(rp RespParams) {
	response := HTTPResponse{
		Status:      false,
		Description: rp.Reason,
	}

	rp.Context.JSON(http.StatusBadRequest, response)
}

func BadLogging(rp RespParams) {
	switch rp.Severity {
	case DEBUG:
		rp.Log.Debug(rp.Section,
			zap.String("connection", rp.URL),
			zap.Any("parameters", rp.Input),
			zap.Error(rp.Error))
	case WARN:
		rp.Log.Warn(rp.Section,
			zap.String("connection", rp.URL),
			zap.Any("parameters", rp.Input),
			zap.Error(rp.Error))
	case ERROR:
		rp.Log.Error(rp.Section,
			zap.String("connection", rp.URL),
			zap.Any("parameters", rp.Input),
			zap.Error(rp.Error))
	}
}

func GoodResponse(c *gin.Context, data interface{}) {
	returnData, _ := json.Marshal(data)
	response := HTTPResponse{
		Status: true,
		Data:   string(returnData),
	}
	c.JSON(http.StatusOK, response)
}

type Mailer struct {
	Host     string
	Port     int
	Username string
	Password string
}

func NewMailer(host, username, password string, port int) Mailer {
	return Mailer{
		Host:     host,
		Username: username,
		Password: password,
		Port:     port,
	}
}
