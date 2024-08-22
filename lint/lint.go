package lint

type Output struct {
	Rows [][]any
}

type Foo struct {
	Name any
}

func castRawRowsToDataPoints2(output *Output, aggregationGroupingKey *string) ([]Foo, error) {
	result := make([]Foo, len(output.Rows))

	for i := range output.Rows {
		if aggregationGroupingKey == nil {
			result[i].Name = ""
			continue
		}

	}
	return result, nil
}
