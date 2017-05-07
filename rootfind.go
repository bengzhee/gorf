package rootfind

import "errors"

type Evalfunc func(float64) float64

func Secant(fn Evalfunc, x0 float64, tol float64) float64 {
    var next_f, x1 float64

	if x0 >= 0 {
		x1 = x0 * (1 + 1e-4) + 1e-4
	} else {
		x1 = x0 * (1 + 1e-4) - 1e-4
	}

	next_f = tol 
    for next_f >= tol {  
        fx0 := fn(x0)  
        fx1 := fn(x1)  
        next_x := (x0 * fx1 - x1 * fx0) / (fx1 - fx0)  
        x0 = x1  
        x1 = next_x  
        next_f = fn(next_x)  
	}

    return x1
}

func Bisect(fn Evalfunc, left float64, right float64, tol float64) (float64, error) {
	var centre, fcentre float64

	fleft := fn(left)
	fright := fn(right)
	if fleft == 0 {

		return left, nil

	} else if fright == 0 {

		return right, nil

	} else if fleft * fright > 0 {
		
		return -1, errors.New("Signs are the same for given bracketing interval")

	} else {

		fcentre = tol
		for fcentre >= tol {

			centre = (left + right) / 2
			fcentre = fn(centre)
			
			if fcentre * fleft > 0 {

				left = centre
				
			} else {

				right = centre
				
			}
		}
	}
	
	return centre, nil
}
