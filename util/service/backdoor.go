// Copyright 2014 The roc Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rocserv

import (
	"github.com/julienschmidt/httprouter"
	"os"

	"github.com/shawnfeng/sutil/slog"
	"github.com/shawnfeng/sutil/snetutil"
)

type backDoorHttp struct{}

func (m *backDoorHttp) Init() error {
	return nil
}

func (m *backDoorHttp) Driver() (string, interface{}) {
	//fun := "backDoorHttp.Driver -->"

	router := httprouter.New()
	// 重启
	//router.POST("/restart", snetutil.HttpRequestWrapper(FactoryRestart))
	router.GET("/backdoor/health/check", snetutil.HttpRequestWrapper(FactoryHealthCheck))

	return "0.0.0.0:60000", router
}

// ==============================
type Restart struct {
}

func FactoryRestart() snetutil.HandleRequest {
	return new(Restart)
}

func (m *Restart) Handle(r *snetutil.HttpRequest) snetutil.HttpResponse {

	slog.Infof("RECEIVE RESTART COMMAND")
	os.Exit(0)
	// 这里的代码执行不到了，因为之前已经退出了
	return snetutil.NewHttpRespString(200, "{}")
}

// ==============================
type HealthCheck struct {
}

func FactoryHealthCheck() snetutil.HandleRequest {
	return new(HealthCheck)
}

func (m *HealthCheck) Handle(r *snetutil.HttpRequest) snetutil.HttpResponse {
	fun := "HealthCheck -->"
	slog.Infof("%s in", fun)

	return snetutil.NewHttpRespString(200, "{}")
}
