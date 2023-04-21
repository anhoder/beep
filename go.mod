module github.com/faiface/beep

go 1.14

require (
	github.com/cocoonlife/goflac v0.0.0-20170210142907-50ea06ed5a9d
	github.com/cocoonlife/testify v0.0.0-20160218172820-792cc1faeb64 // indirect
	github.com/gdamore/tcell v1.3.0
	github.com/hajimehoshi/go-mp3 v0.3.1
	github.com/hajimehoshi/oto v0.7.1
	github.com/jfreymuth/oggvorbis v1.0.1
	github.com/mewkiz/flac v1.0.7
	github.com/pkg/errors v0.9.1
	github.com/tosone/minimp3 v1.0.1
)

replace (
	github.com/cocoonlife/goflac v0.0.0-20170210142907-50ea06ed5a9d => github.com/go-musicfox/goflac v0.1.5
	github.com/tosone/minimp3 v1.0.1 => github.com/go-musicfox/minimp3 v1.0.6
)
