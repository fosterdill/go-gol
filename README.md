# Game of life in go

In order to set up a workspace add this to your `~/.bash_profile`

    export GOPATH=$HOME/[go folder path]
    export PATH=$HOME/[go folder path]/bin:$PATH
    
where `[go folder path]` is where you have the root of your go workspace

Once you have setup your go folder, you need to run

    cd ~/[go folder path]
    go get github.com/fosterdill/gol
    cd github.com/fosterdill/gol
    go install
    gol
    
The binary will be put in ~/[go folder path]/bin/gol
