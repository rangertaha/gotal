package indicators


var (
	err           error
	Example       Series
	ExampleStream Stream
)

func init() {
	Example, ExampleStream, err = Get("example")

	if err != nil {
		panic(err)
	}
}
