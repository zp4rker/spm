//go:build install

//go:generate mv -f ./spm $HOME/go/bin/
//go:generate mv -f ./spm-controller $HOME/go/bin/
//go:generate mv -f ./spm-wrapper $HOME/go/bin/

package build