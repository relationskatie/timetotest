package controller

import "go.uber.org/fx"

func StartFx(lc fx.Lifecycle, ctrl Controller) {
	lc.Append(fx.Hook{
		OnStart: ctrl.Run,
		OnStop:  ctrl.Shutdown,
	})
}
