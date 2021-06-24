package bytedance

// type ByteDance struct {
// 	Config  Config
// 	AccessToken auth.AccessToken
// 	Logger *log.Logger
// 	Cache *gcache.Cache
// }

// func New(config Config) *ByteDance {
// 	app := NewWithOutToken(config)
// 	app.AccessToken = app.getDefaultAccessToken()
// 	return app
// }

// func NewWithOutToken(config Config) *MiniProgram {
// 	if config.Cache == nil {
// 		config.Cache = cache.New("ewa.wechat.miniprogram")
// 	}
// 	if config.Logger == nil {
// 		config.Logger = log.New()
// 		if config.Logger.LogPath != "" {
// 			if err := config.Logger.SetPath(config.Logger.LogPath); err != nil {
// 				panic(fmt.Sprintf("[miniprogram] set log path '%s' error: %v", config.Logger.LogPath, err))
// 			}
// 		}

// 		// default set close debug / close stdout print
// 		config.Logger.LogStdout = false
// 	}
// 	var app = &MiniProgram{
// 		Config: config,
// 		Logger: config.Logger,
// 		Cache:  config.Cache,
// 	}
// 	return app
// }
