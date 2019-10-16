
## Install
`go get github.com/alabianca/spin`

## Use
To initialize a spinner first create the instance using `spinners.NewSpinner(spinners.Dots, os.Stdout)`

Call `Start()` on the spinner inside a separate go routine.
The spinner will spin until `Stop()` is called.