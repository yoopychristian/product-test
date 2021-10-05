package functions

import (
	"crypto/rsa"
	"crypto/x509"
	"database/sql"
	"encoding/pem"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func EnvInt(envName string) int {
	val, _ := strconv.Atoi(os.Getenv(envName))
	return val
}
func EnvBool(envName string) bool {
	val := os.Getenv(envName)
	return strings.ToUpper(val) == "TRUE"
}
func EnvString(envName string) string {
	return os.Getenv(envName)
}

func GetPrivateKeyFromPem(keyPem string) (*rsa.PrivateKey, error) {
	handleErr := func(err error) (*rsa.PrivateKey, error) {
		return nil, fmt.Errorf("pem private key function : %s ", err)
	}

	pem, rest := pem.Decode([]byte(keyPem))
	if len(rest) != 0 {
		return handleErr(fmt.Errorf("failed to decode key pem , value :%s", keyPem))
	}

	pKey, err := x509.ParsePKCS1PrivateKey(pem.Bytes)
	if err != nil {
		return handleErr(fmt.Errorf("parse key : %s", err))
	}

	return pKey, nil
}

func GetPubKeyFromPem(keyPem string) (*rsa.PublicKey, error) {
	handleErr := func(err error) (*rsa.PublicKey, error) {
		return nil, fmt.Errorf("pem public key function : %s ", err)
	}

	appKeyBlock, rest := pem.Decode([]byte(keyPem))
	if len(rest) != 0 {
		return handleErr(fmt.Errorf("failed to decode key pem , value :%s", keyPem))
	}
	iAppKey, err := x509.ParsePKIXPublicKey(appKeyBlock.Bytes)
	if err != nil {
		return handleErr(fmt.Errorf("parse key : %s", err))
	}
	key, ok := iAppKey.(*rsa.PublicKey)
	if !ok {
		return handleErr(fmt.Errorf("final key : %s", err))
	}
	return key, nil
}

//LogInit initiate log file
func LogInit(d bool, f *os.File) *zap.Logger {
	pe := zap.NewDevelopmentEncoderConfig()

	fileEncoder := zapcore.NewJSONEncoder(pe)
	pe.EncodeTime = zapcore.ISO8601TimeEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(pe)

	level := zap.InfoLevel
	if d {
		level = zap.DebugLevel
	}

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(f), level),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level),
	)

	l := zap.New(core)

	return l
}

type DBParam struct {
	Host     string
	Port     string
	Name     string
	Schema   string
	User     string
	Password string
	AppName  string
	Timeout  int
	MaxOpen  int
	MaxIdle  int
	Logging  bool
}

func DBInit(p DBParam, l *zap.Logger, path string, debugMode bool) (*gorm.DB, error) {

	file, err := os.Create(path + "_" + time.Now().Format("2006_01_02__15_04") + ".log")
	if err != nil {
		return nil, err
	}
	level := logger.Silent
	if debugMode {
		level = logger.Info
	}
	newLogger := logger.New(
		log.New(file, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  level,       // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	//sql connection
	sqlConn, err := sql.Open("postgres", makePostgresString(p))
	if err != nil {
		return nil, errors.Wrap(err, "can't establish db connection")
	}
	sqlConn.SetMaxIdleConns(p.MaxIdle)
	sqlConn.SetMaxOpenConns(p.MaxOpen)
	sqlConn.SetConnMaxLifetime(time.Hour)

	db, err := gorm.Open(postgres.New(
		postgres.Config{Conn: sqlConn}),
		&gorm.Config{Logger: newLogger})
	if err != nil {
		return nil, errors.Wrap(err, "can't open db connection")
	}

	return db, err
}

func makePostgresString(p DBParam) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s connect_timeout=%d application_name=%s",
		p.Host, p.Port, p.User, p.Name, p.Password, p.Timeout, p.AppName)
}

//random number int
func RandomNumber() int {
	min := 100000
	max := 999999
	randomNumberPin := rand.Intn(max-min) + min
	return randomNumberPin
}

func StructToUrlValue(data interface{}) (url.Values, error) {
	return query.Values(data)
}
