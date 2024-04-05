# echo the current working directory
cwd() {
  echo "$(cwd)"
}

go build -o test.1s \
&& mv test.1s ~/Library/Application\ Support/xbar/plugins/test.1s.cgo \
&& chmod +x ~/Library/Application\ Support/xbar/plugins/test.1s.cgo
