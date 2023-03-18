package models

import "fmt"

type Saiyan struct {
	Name     string
	Powerlvl int
	SSJ      bool
}

func (s *Saiyan) DescribeSaiyan() {
	fmt.Printf("Name : %v, Powerlvl : %v, SSJ : %v \n", s.Name, s.Powerlvl, s.SSJ)
}
