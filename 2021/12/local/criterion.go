package local

type Criterion func(route Track) (startFlag, endFlag, customFlag bool)
