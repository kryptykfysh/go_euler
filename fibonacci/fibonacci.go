package fibonacci

func Generator() chan int {
  g := make(chan int)

  go func() {
    for i, j:= 1, 1; ; i, j = i+j, i {
      g <- i
    }
  }()

  return g
}
