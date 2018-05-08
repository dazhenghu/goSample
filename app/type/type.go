package _type

type errorA struct {
    s string
}

func (e *errorA) Error() string  {
    return e.s
}

type errorB struct {
    s string
}

func (e *errorB) Error() string  {
    return e.s
}

