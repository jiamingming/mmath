module github.com/jiamingming/mmath

go 1.12

replace (
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2 => ./github.com/golang/crypto
	golang.org/x/exp v0.0.0-20190125153040-c74c464bbbf2 => ./github.com/golang/exp
	golang.org/x/exp v0.0.0-20190306152737-a1d7652674e8 => ./github.com/golang/exp
	golang.org/x/image v0.0.0-20190227222117-0694c2d4d067 => ./github.com/golang/image

	golang.org/x/mobile v0.0.0-20190312151609-d3739f865fa6 => ./github.com/golang/mobile
	golang.org/x/net v0.0.0-20190311183353-d8887717615a => ./github.com/golang/net
	golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3 => ./github.com/golang/net

	golang.org/x/sync v0.0.0-20190423024810-112230192c58 => ./github.com/golang/sync
	golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a => ./github.com/golang/sys
	golang.org/x/sys v0.0.0-20190312061237-fead79001313 => ./github.com/golang/sys
	golang.org/x/sys v0.0.0-20190412213103-97732733099d => ./github.com/golang/sys
	golang.org/x/text v0.3.0 => ./github.com/golang/text
	golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e => ./github.com/golang/tools
	golang.org/x/tools v0.0.0-20190206041539-40960b6deb8e => ./github.com/golang/tools

	golang.org/x/tools v0.0.0-20190312151545-0bb0c0a6e846 => ./github.com/golang/tools

)

require (
	github.com/sjwhitworth/golearn v0.0.0-20190323185925-82e59c89f502 // indirect
	gonum.org/v1/gonum v0.0.0-20190608115022-c5f01565d866 // indirect
)
