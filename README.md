# somethingforsomethingother

The tool can be used to get MD5 of response for list of URL

On three examples I show you how to use it

# Example 1 
Let's run it throught next command - `make example`

You will see something like:
```
go build -o app main.go
./app http://www.adjust.com http://google.com
http://www.adjust.com 53811605a6fe178692408cbbbb82bb26
http://google.com e0381ef26a5c1694a2992dd955f322e7
```

# Example 2 (You can control maximum value of parallel requests, by default it equal to 10)
Let's run it throught next command - `make example2` and in output you should see flag to set another value of parallel requests

You will see something like:
```
go build -o app main.go
./app -parallel 3 adjust.com google.com facebook.com yahoo.com yandex.com twitter.com reddit.com/r/funny reddit.com/r/notfunny baroquemusiclibrary.com
http://baroquemusiclibrary.com 697056a33db2d236044fee127cd0af74
http://google.com cb6744676156b65e8b22c26112d3d252
http://facebook.com 0d7a626fa607ff7ff0fa0839dfeabf5e
http://twitter.com 4ab3b1210026da0f6eb40e2f33c8bc4f
http://reddit.com/r/funny 33ec625a7a4dee96106b7db184e1a71a
http://reddit.com/r/notfunny b1642486ddaff6aec4a65b76f2558122
http://adjust.com 53811605a6fe178692408cbbbb82bb26
http://yandex.com 25cc2f187cbbda8de26ef373e59da6f6
http://yahoo.com 89d190259ae04d32d662091120bace0c
```