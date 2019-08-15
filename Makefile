data := $(shell date "+%d %b %Y")

all: doc-todo.html


doc-todo.html: doc-todo.asc main
	asciidoc -b html5 -a icons -a toc2 -a theme=flask -a revdate="$(data)"  $<

main: main.go decoration.go style.go system.go
	go build $^
