package main

import "fmt"

func S(x string) bool {
	if len(x) >= 1 && x[0] == 'k' {
		if len(x) >= 2 && x[1] == 'i' {
			if len(x) >= 3 && x[2] == 't' {
				if len(x) >= 4 && x[3] == 'a' {
					return true
				}
			}
		} else if len(x) >= 2 && x[1] == 'a' {
			if  len(x) >= 3 && x[2] == 'k' {
				if len(x) >= 4 && x[3] == 'u' {
				return true
			}
		}
		}
	} else if len(x) >= 1 && x[0] == 'd' {
		if len(x) >= 2 && x[1] == 'i' {
			if len(x) >= 3 && x[2] == 'a' {
				return true
			}
		}
	} else if len(x) >= 1 && x[0] == 'u' {
		if len(x) >= 2 && x[1] == 'd' {
			if len(x) >= 3 && x[2] == 'i' {
				if len(x) >= 4 && x[3] == 'n' {
					return true
				}
			}
		}
	} else if len(x) >= 1 && x[0] == 'a' {
		if len(x) >= 2 && x[1] == 'n' {
			if len(x) >= 3 && x[2] == 'i' {
				return true
			}
		}
	}
	return false
}

func P(x string) bool {
	if len(x) >= 3 && x[0] == 'm' && x[1] == 'e' && x[2] == 'm' {
		if len(x) >= 4 && x[3] == 'a' {
			if len(x) >= 5 && x[4] == 's' {
				if len(x) >= 6 && x[5] == 'a' {
					if len(x) >= 7 && x[6] == 'k' {
						return true
					}
				}
			} else if len(x) >= 5 && x[4] == 'k' {
				if len(x) >= 6 && x[5] == 'a' {
					if len(x) >= 7 && x[6] == 'n' {
						return true
					}
				}
			}
		} else if len(x) >= 4 && x[3] == 'b' {
			if len(x) >= 5 && x[4] == 'e' {
				if len(x) >= 6 && (x[5] == 'l' || x[5] == 'r') {
					if len(x) >= 7 && x[6] == 'i' {
						return true
					}
				}
			} else if len(x) >= 5 && x[4] == 'u' {
				if len(x) >= 6 && x[5] == 'k' {
					if len(x) >= 7 && x[6] == 'a' {
						return true
					}
				}
			}
		}
	}
	return false
}

func O(x string) bool {
	if len(x) >= 1 && x[0] == 'p' {
		if len(x) >= 2 && x[1] == 'i' {
			if len(x) >= 3 && x[2] == 'e' {
				return true
			}
		}
	} else if len(x) >= 1 && x[0] == 'a' {
		if len(x) >= 2 && x[1] == 'y' {
			if len(x) >= 3 && x[2] == 'a' {
				if len(x) >= 4 && x[3] == 'm' {
					return true
				}
			}
		}
	} else if len(x) >= 1 && x[0] == 'k' {
		if len(x) >= 2 && x[1] == 'u' {
			if len(x) >= 3 && x[2] == 'e' {
				return true
			}
		}
	} else if len(x) >= 1 && x[0] == 'm' {
		if len(x) >= 2 && x[1] == 'i' {
			if len(x) >= 3 && x[2] == 'e' {
				return true
			}
		}
	} else if len(x) >= 1 && x[0] == 's' {
		if len(x) >= 2 && x[1] == 'n' {
			if len(x) >= 3 && x[2] == 'a' {
				if len(x) >= 4 && x[3] == 'c' {
					if len(x) >= 5 && x[4] == 'k' {
						return true
					}
				}
			}
		}
	}
	return false
}

func K(x string) bool {
	if len(x) >= 2 && x[0] == 'd' && x[1] == 'i' {
		if len(x) >= 3 && x[2] == 't' {
			if len(x) >= 4 && x[3] == 'e' {
				if len(x) >= 5 && x[4] == 'k' {
					if len(x) >= 6 && x[5] == 'o' {
						return true
					}
				}
			}
		} else if len(x) >= 3 && x[2] == 'k' {
			if len(x) >= 4 && x[3] == 'o' {
				if len(x) >= 5 && x[4] == 't' {
					if len(x) >= 6 && x[5] == 'a' {
						return true
					}
				}
			} else if len(x) >= 4 && x[3] == 'o' {
				if len(x) >= 5 && x[4] == 's' {
					return true
				}
			}
		} else if len(x) >= 3 && x[2] == 'r' {
			if len(x) >= 4 && x[3] == 'u' {
				if len(x) >= 5 && x[4] == 'm' {
					if len(x) >= 6 && x[5] == 'a' {
						if len(x) >= 7 && x[6] == 'h'{
						return true
					}
				}
			}
		}
		} else if len(x) >= 3 && x[2] == 'b' {
			if len(x) >= 4 && x[3] == 'u' {
				if len(x) >= 5 && x[4] == 's' {
					return true
				}
			}
		}
	}
	return false
}


type stack [100]string

func push(s *stack, kata string, x *int) {
	s[*x] = kata
	*x = *x + 1
}

func pop(s *stack, n *int) string {
	status := s[*n-1]
	*n = *n - 1
	return status
}

func top(s stack, x int) string {
	return s[x-1]
}

func pSPOK(x [100]byte, y int) bool {
	var s stack
	var n int
	var i int = 0
	push(&s, "#", &n)
	push(&s, "x", &n)
	for top(s, n) != "#" && i < y {
		if top(s, n) == "x" {
			pop(&s, &n)
			push(&s, "K", &n)
			push(&s, "O", &n)
			push(&s, "P", &n)
			push(&s, "S", &n)
		} else if top(s, n) == "S" {
			if x[i] != 'S' {
				return false
			}
			i++
			pop(&s, &n)
		} else if top(s, n) == "P" {
			if x[i] != 'P' {
				return false
			}
			i++
			pop(&s, &n)
		} else if top(s, n) == "O" {
			if x[i] != 'O' {
				return false
			}
			i++
			pop(&s, &n)
		} else if top(s, n) == "K" {
			if x[i] != 'K' {
				return false
			}
			i++
			pop(&s, &n)
		}
	}
	if pop(&s, &n) == "#" {
		return true
	}
	return false
}

func pSPO(x [100]byte, y int) bool {
	var s stack
	var n int
	var i int = 0
	push(&s, "#", &n)
	push(&s, "x", &n)
	for top(s, n) != "#" && i < y {
		if top(s, n) == "x" {
			pop(&s, &n)
			push(&s, "O", &n)
			push(&s, "P", &n)
			push(&s, "S", &n)
		} else if top(s, n) == "S" {
			if x[i] != 'S' {
				return false
			}
			i++
			pop(&s, &n)
		} else if top(s, n) == "P" {
			if x[i] != 'P' {
				return false
			}
			i++
			pop(&s, &n)
		} else if top(s, n) == "O" {
			if x[i] != 'O' {
				return false
			}
			i++
			pop(&s, &n)
		}
	}
	if pop(&s, &n) == "#" {
		return true
	}
	return false
}

func pSPK(x [100]byte, y int) bool {
	var s stack
	var n int
	var i int = 0
	push(&s, "#", &n)
	push(&s, "x", &n)
	for top(s, n) != "#" && i < y {
		if top(s, n) == "x" {
			pop(&s, &n)
			push(&s, "K", &n)
			push(&s, "P", &n)
			push(&s, "S", &n)
		} else if top(s, n) == "S" {
			if x[i] != 'S' {
				return false
			}
			i++
			pop(&s, &n)
		} else if top(s, n) == "P" {
			if x[i] != 'P' {
				return false
			}
			i++
			pop(&s, &n)
		} else if top(s, n) == "K" {
			if x[i] != 'K' {
				return false
			}
			i++
			pop(&s, &n)
		}
	}
	if pop(&s, &n) == "#" {
		return true
	}
	return false
}

func pSP(x [100]byte, y int) bool {
	var s stack
	var n int
	var i int = 0
	push(&s, "#", &n)
	push(&s, "x", &n)
	for top(s, n) != "#" && i < y {
		if top(s, n) == "x" {
			pop(&s, &n)
			push(&s, "P", &n)
			push(&s, "S", &n)
		} else if top(s, n) == "S" {
			if x[i] != 'S' {
				return false
			}
			i++
			pop(&s, &n)
		} else if top(s, n) == "P" {
			if x[i] != 'P' {
				return false
			}
			i++
			pop(&s, &n)
		}
	}
	if pop(&s, &n) == "#" {
		return true
	}
	return false
}

func main() {
	var kata string
	var arr [100]string
	var parse [100]byte
	var yp int = 0
	var n int = 0
	var i int
	fmt.Scan(&kata)
	for kata != "." {
		arr[n] = kata
		n++
		fmt.Scan(&kata)
	}
	for i = 0; i < n; i++ {
		if S(arr[i]) {
			parse[yp] = 'S'
			yp++
		} else if P(arr[i]) {
			parse[yp] = 'P'
			yp++
		} else if O(arr[i]) {
			parse[yp] = 'O'
			yp++
		} else if K(arr[i]) {
			parse[yp] = 'K'
			yp++
		}
	}

	if yp != 0 && pSPOK(parse, yp) {
		fmt.Print("S-P-O-K")
		fmt.Println(" Valid Input")
	} else if yp != 0 && pSPO(parse, yp) {
		fmt.Print("S-P-O")
		fmt.Println(" Valid Input")
	} else if yp != 0 && pSPK(parse, yp) {
		fmt.Print("S-P-K")
		fmt.Println(" Valid Input")
	} else if yp != 0 && pSP(parse, yp) {
		fmt.Print("S-P")
		fmt.Println(" Valid Input")
	} else {
		fmt.Println("Invalid Input")
	}
}