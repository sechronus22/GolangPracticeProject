package models

type Faculty struct{
    Name string
    Code int
    Information string
    Department []*Department
}

type Department struct{
    Name string
    Abbreviation string
}