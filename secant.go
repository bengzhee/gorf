//package rootfind
package main
import ( 
    "fmt" 
    "math" 
    "time"
)
type evalfunc func(float64) float64

func secant(fn evalfunc, x0 float64, tol float64) float64 {
    next_f := 1.0 
    x1 := x0 * 1.1  
    for next_f >= tol {  
        fx0 := fn(x0)  
        fx1 := fn(x1)  
        next_x := (x0 * fx1 - x1 * fx0) / (fx1 - fx0)  
        x0 = x1  
        x1 = next_x  
        next_f = fn(next_x)  
        fmt.Println(next_x, next_f) }
    return x1
}

type bond struct { 
    startDate time.Time 
    endDate time.Time 
    coupon float64 
    hazardRate float64 
    frequency int64 
    discount float64
}

func NewBond(options ...func(*bond)) *bond {
    bond := bond{}  
        for _, option := range options {  
            option(&bond) 
        }  
    return &bond
}

func bondCreate() *bond { 
  var (  
      start time.Time = time.Date(2016, time.October, 18, 0, 0, 0, 0, time.UTC)  
      end time.Time = time.Date(2017, time.February, 22, 23, 59, 0, 0, time.UTC)  
      coupon, discount float64 = 0.05, 0.98  
      frequency int64 = 2 
  )  

  startDate := func(bond *bond){bond.startDate = start} 
  endDate := func(bond *bond){bond.endDate = end} 
  couponrate := func(bond *bond){bond.coupon = coupon} 
  freq := func(bond *bond){bond.frequency = frequency} 
  disc := func(bond *bond){bond.discount = discount}  
  bond := NewBond(startDate, endDate, couponrate, freq, disc)
  return bond}

func (b *bond) bondPrice(hazrat float64) float64 {
    var cashFlows []float64 
    var cashFlowDates []time.Time 
    b.hazardRate = hazrat  
    
    totalCashFlows := 0 
    delMonths := int(12 / b.frequency) 
    cashFlowDates = append(cashFlowDates, b.endDate) 
    cashFlows = append(cashFlows, b.coupon / float64(b.frequency) * 100 + 100) 
    
    for cashFlowDates[len(cashFlowDates)-1].After(b.startDate) {  
        totalCashFlows++  cashFlows = append(cashFlows, b.coupon / float64(b.frequency) * 100 * b.hazardRate)  
        thisCFDate := b.endDate.AddDate(0, -delMonths, 0)  
        cashFlowDates = append(cashFlowDates, thisCFDate) 
    }
    
    price := 0.0 
    
    for _, cf := range cashFlows {  
        price += cf 
    }  
    
    return price
}

func newcubic(x float64) float64 { 
    return 8.213 * math.Pow(x, 3) + 12.47 * x + 4.512
}

func main() { 
    var t float64 = math.Pow(10, -15) 
    bond1 := bondCreate() 
    
    z := secant(bond1.bondPrice, -0.34, t) 
    fmt.Println(z)
}
