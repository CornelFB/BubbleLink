package api

import (
	"net/http"
)

func (rt *_router) Handler() http.Handler {

	rt.router.POST("/session", rt.wrap(rt.doLogin))

	rt.router.PUT("/users/:id/name", rt.wrap(rt.setMyUserName))
	rt.router.GET("/users/:id/name", rt.wrap(rt.getMyUserName))
	rt.router.PUT("/users/:id/photo", rt.wrap(rt.setMyPhoto))
	rt.router.GET("/users/:id/photo", rt.wrap(rt.getMyPhoto))

	//rt.router.GET("/placeExterior", rt.wrap(rt.getAllPlaces))

	//rt.router.POST("/placeExterior/add", rt.wrap(rt.addPlace))
	rt.router.GET("/placeExterior/:placeId", rt.wrap(rt.getPlace))

	//rt.router.GET("/users/:id/existence", rt.wrap(rt.getUserExistence))

	rt.router.GET("/users/:id/existence", rt.wrap(rt.getUserExistence))

	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
