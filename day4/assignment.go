package main

type Section struct {
	Start int
	End   int
}

func (s Section) Contains(is Section) bool {
	return s.Start <= is.Start && s.End >= is.End
}

func (s Section) HasCommon(is Section) bool {
	return (s.Start <= is.Start && s.End >= is.Start && s.End <= is.End) ||
		(is.Start <= s.Start && is.End >= s.Start && is.End <= s.End) ||
		s.Contains(is) ||
		is.Contains(s)
}

type Assignment struct {
	Section1 Section
	Section2 Section
}

func (a Assignment) FullOverlap() bool {
	return a.Section1.Contains(a.Section2) || a.Section2.Contains(a.Section1)
}

func (a Assignment) Overlap() bool {
	return a.Section1.HasCommon(a.Section2)
}
