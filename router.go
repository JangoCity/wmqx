package main

import (
	"github.com/buaazp/fasthttprouter"
	"rmqc/app/controllers"
)

type Router struct {

}

// return router
func NewRouter() *Router {
	return &Router{}
}

// api manager server router
func (r *Router) Api() *fasthttprouter.Router {
	router := fasthttprouter.New()

	// message router
	messageController := controllers.NewMessageController()
	router.GET("/message/add", messageController.Add)
	router.GET("/message/update", messageController.Update)
	router.GET("/message/delete", messageController.Delete)
	router.GET("/message/status", messageController.Status)
	router.GET("/message/list", messageController.List)
	router.GET("/message/getMessageByName", messageController.GetMessageByName)
	router.GET("/message/getConsumerByName", messageController.GetConsumerByName)

	// consumer router
	consumerController := controllers.NewConsumerController()
	router.GET("/consumer/add", consumerController.Add)
	router.GET("/consumer/update", consumerController.Update)
	router.GET("/consumer/delete", consumerController.Delete)
	router.GET("/consumer/status", consumerController.Status)
	router.GET("/consumer/getConsumerById", consumerController.GetConsumerById)

	// system router
	systemController := controllers.NewSystemController()
	router.GET("/system/reload", systemController.Reload)
	router.GET("/system/restart", systemController.Restart)

	// log router
	logController := controllers.NewLogController()
	router.GET("/log", logController.Index)
	router.GET("/log/file", logController.File)
	router.GET("/log/list", logController.List)

	return router
}

// publish server router
func (r *Router) Publish() *fasthttprouter.Router {
	router := fasthttprouter.New()

	// publish router
	publishController := controllers.NewPublishController()
	router.GET("/publish/:name", publishController.Publish)
	router.POST("/publish/:name", publishController.Publish)

	return router
}