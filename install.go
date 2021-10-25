//go:build install

//go:generate mv -f ./jpm $HOME/go/bin/
//go:generate mv -f ./jpm-controller $HOME/go/bin/
//go:generate mv -f ./jpm-wrapper $HOME/go/bin/

package build