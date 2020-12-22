module github.com/tiiuae/dronsole-containers/gzserver

go 1.14

require (
	github.com/julienschmidt/httprouter v1.3.0
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	k8s.io/api v0.19.3
	k8s.io/apimachinery v0.19.3
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/utils v0.0.0-20201104234853-8146046b121e // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.0.0-20200214081623-ecbd4af0fc33
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20200214081019-7490b3ed6e92
	k8s.io/client-go => k8s.io/client-go v0.0.0-20200214082307-e38a84523341
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20200214080538-dc8f3adce97c
)
