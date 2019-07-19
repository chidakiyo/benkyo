module github.com/chidakiyo/benkyo/appengine-relative-mod/appli_a

go 1.12

require (
	github.com/chidakiyo/benkyo/appengine-relative-mod/lib_a v0.0.0
	github.com/chidakiyo/benkyo/appengine-relative-mod/lib_b v0.0.0 // indirect
	google.golang.org/appengine v1.6.1
)

replace (
	github.com/chidakiyo/benkyo/appengine-relative-mod/lib_a => ../lib_a
	github.com/chidakiyo/benkyo/appengine-relative-mod/lib_b => ../lib_b
)
