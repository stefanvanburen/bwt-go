name: Go
on: [push]
jobs:

  build:
    name: Build
    runs-on: ${{ matrix.os }}
    strategy:
      matrix:
        os: [ubuntu-latest, macOS-latest]
    
    steps:

    - name: Set up Go 1.12
      uses: actions/setup-go@v1
      with:
        go-version: 1.12
      id: go

    - name: Check out code into the Go module directory
      uses: actions/checkout@v1

    - name: Get dependencies
      run: |
        go mod download
    - name: Test
      run: go test -v -race ./...
