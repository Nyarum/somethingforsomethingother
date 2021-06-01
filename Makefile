
build:
	go build -o app main.go

run: build
	./app

example: build
	./app http://www.adjust.com http://google.com

example2: build
	./app -parallel 3 adjust.com google.com facebook.com yahoo.com yandex.com twitter.com reddit.com/r/funny reddit.com/r/notfunny baroquemusiclibrary.com

example3: build
	./app adjust.com