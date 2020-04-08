package darknet

// #cgo CFLAGS: -I/usr/local/include -I/usr/local/cuda/include -I/usr/local/cuda-10.0/targets/x86_64-linux/include
// #cgo LDFLAGS: -L/usr/lib -ldarknet -lm -L/usr/local/cuda-10.0/targets/x86_64-linux/lib -L/usr/local/cuda-10.0/compat/
import "C"
