package soln

import (
    "fmt"
    
    "aoc2018/util"
)

func solve3() {
    data := util.ReadLines("data/day03.txt", "[#@,:x]")
    claims := getClaims(data)
    f, o := getFabricAndOverlaps(claims)

    overlapCount := countOverlapsMoreThanK(f, 1)
    fmt.Printf("3A: %d\n", overlapCount)

    claimIDWithNoOverlap := getClaimIDWithNoOverlap(o)
    fmt.Printf("3B: %d\n", claimIDWithNoOverlap)
}

type fabricMap map[coordinateType]*util.IntSet

type overlapsMap map[int]*util.IntSet

type coordinateType struct {
    x int
    y int
}

type claimType struct {
    id int
    coord coordinateType
    width int
    height int
}

func getClaims(data [][]string) []claimType {
    claims := []claimType{}

    for _, datum := range data {
        id := util.StringToInt(datum[1])
        x := util.StringToInt(datum[2])
        y := util.StringToInt(datum[3])
        width := util.StringToInt(datum[4])
        height := util.StringToInt(datum[5])

        claims = append(claims, claimType{id, coordinateType{x, y}, width, height})
    }

    return claims
}

func getFabricAndOverlaps(claims []claimType) (fabricMap, overlapsMap) {
    fabric := make(fabricMap)
    overlaps := make(overlapsMap)

    for _, claim := range claims {
        updateFabricAndOverlaps(claim, fabric, overlaps)
    }

    return fabric, overlaps
}

func updateFabricAndOverlaps(claim claimType, fabric fabricMap, overlaps overlapsMap) {
    overlaps[claim.id] = util.NewIntSet()

    for i := claim.coord.x; i < claim.coord.x + claim.width; i++ {
        for j := claim.coord.y; j < claim.coord.y + claim.height; j++ {
            if ids, ok := fabric[coordinateType{i, j}]; ok {
                for _, id := range ids.Values() {
                    overlaps[claim.id].Add(id)
                    overlaps[id].Add(claim.id)
                }
            } else {
                fabric[coordinateType{i, j}] = util.NewIntSet()
            }

            fabric[coordinateType{i, j}].Add(claim.id)
        }
    }
}

func countOverlapsMoreThanK(fabric fabricMap, k int) int {
    count := 0

    for _, ids := range fabric {
        if ids.Size() > k {
            count++
        }
    }

    return count
}

func getClaimIDWithNoOverlap(overlaps overlapsMap) int {	
    for id, overlap := range overlaps {
        if overlap.Size() == 0 {
            return id
        }
    }

    return -1
}