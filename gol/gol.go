package main

func getWithWrapAround(a int, size int) int{
	if (a<0){
		a = size-1
	}
	if(a>=size){
		a = 0
	}
	return a
}

func calculateNextState(p golParams, world [][]byte) [][]byte {
	var newWorld [][]byte
	aliveCoords := calculateAliveCells(p,world)
	var matrixOfNeighbours [][]int

	// creating matrix of neighbour numbers
	for _,item := range aliveCoords{
		x := item.x
		y := item.y

		//because you are not your own neighbour
		matrixOfNeighbours[x][y] -=1

		// add neighbours
		for i :=-1;i<2;i++{
			for j :=-1;j<2;j++{
				Xindex := getWithWrapAround(x+i,p.imageWidth)
				Yindex := getWithWrapAround(y+j,p.imageHeight)
				matrixOfNeighbours[Xindex][Yindex] +=1
			}
		}
	}

	for i,_ := range world{
		for j,_ := range world[i]{
			neighbours := matrixOfNeighbours[i][j]
			if (neighbours == 3){
				newWorld[i][j] = 1;
			}
			if (neighbours == 2 && world[i][j]==1){
				newWorld[i][j] = 1
			}
		}
	}

	return newWorld
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	var alive []cell

	for i,_ := range world{
		for j,_ := range world[i]{
			if(world[i][j] == 1){
				var newCell cell
				newCell.x = i;
				newCell.y = j;
				alive = append(alive, newCell)
			}
		}
	}

	return []cell{}
}
