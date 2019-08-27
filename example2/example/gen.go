package main

import "context"

var TaskName = "{{YOUR_TASK_NAME}}"
var PrefixDateTime = "20190801_190000_"

func main() {
	ctx := context.Background()
	if err := Handler(ctx); err != nil {
		panic(err)
	}
}

func Handler(ctx context.Context) error {
	return nil
}
