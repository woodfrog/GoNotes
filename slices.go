package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) (pics [][]uint8) {
        pics = make([][]uint8, dy)
        for i := range(pics) {
                    pics[i] = make([]uint8, dx)
                    for j := range(pics[i]){
                                    pics[i][j] = uint8((i+j) / 2)
                                            
                    }
                        
        }
            return

}

func main() {
        pic.Show(Pic)

}
