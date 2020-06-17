package models

type Faculty struct{
    Name string
    Code int
    Department []*Department
}

type Department struct{
    Name string
    Abbreviation string
}