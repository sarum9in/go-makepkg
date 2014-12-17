package main

import "text/template"

var pkgbuildTemplate = template.Must(
	template.New("pkgbuild").Parse(`pkgname={{.PkgName}}
pkgver="autogenerated"
pkgrel={{.PkgRel}}
pkgdesc="{{.PkgDesc}}"
arch=('i686' 'x86_64')
license=('{{.License}}')
makedepends=('go' 'git')

source=(
	"{{.RepoUrl}}"{{range .Files}}
	"{{.Name}}"{{end}}
)

md5sums=(
	'SKIP'{{range .Files}}
	'{{.Hash}}'{{end}}
)

backup=({{range .Backup}}
	"{{.}}"{{end}}
)

pkgver() {
	cd "$srcdir/$pkgname"
	git log -1 --format="%cd" --date=short | sed s/-//g
}

build() {
	cd "$srcdir/$pkgname"

	rm -rf "$srcdir/.go/src"

	mkdir -p "$srcdir/.go/src"

	export GOPATH="$srcdir/.go"

	mv "$srcdir/$pkgname" "$srcdir/.go/src/"

	cd "$srcdir/.go/src/$pkgname/"
	ln -sf "$srcdir/.go/src/$pkgname/" "$srcdir/$pkgname"

	echo "Running 'go get'..."
	go get
}

package() {
	install -DT "$srcdir/.go/bin/$pkgname" "$pkgdir/usr/bin/$pkgname"{{range .Files}}
	install -DT -m0755 "$srcdir/{{.Name}}" "$pkgdir/{{.Path}}"{{end}}
}
`))
