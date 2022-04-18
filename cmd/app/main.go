package main

import (
	stdhttp "net/http"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/mars-terminal/proxy-cache/internal/server/http"
	"github.com/mars-terminal/proxy-cache/internal/server/http/handler"
	service "github.com/mars-terminal/proxy-cache/internal/service/proxy"
	"github.com/mars-terminal/proxy-cache/internal/storage/mongo"
	"github.com/mars-terminal/proxy-cache/internal/storage/mongo/proxy"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatal("cannot read config")
		return
	}

	var options = struct {
		URL        string
		DBName     string
		Collection string

		HttpPort string
	}{
		URL:        viper.GetString("DB.URL"),
		DBName:     viper.GetString("DB.Name"),
		Collection: viper.GetString("DB.Collection"),

		HttpPort: viper.GetString("Port"),
	}

	mongoStorage, err := mongo.NewMongo(options.URL)
	if err != nil {
		logrus.Fatal("cannot connect to mongo")
		return
	}
	logrus.Info("DB connected")

	mongoDB := mongoStorage.Database(options.DBName)

	ProxyStorage := proxy.NewStorage(mongoDB.Collection(options.Collection))

	ProxyService := service.NewStore(ProxyStorage)

	r := stdhttp.NewServeMux()

	handler.NewHandler(ProxyService).SetupRouter(r)

	srv := new(http.Server)
	if err := srv.Run(options.HttpPort, r); err != nil {
		logrus.WithError(err).Fatal("error occured while running server server")
		return
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
