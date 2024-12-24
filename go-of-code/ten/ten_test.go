package ten

import "testing"

func TestBuildTree(t *testing.T) {
	tree := &Tree{position: Point{x: 1, y: 1, height: 1}}
	topograficMap := &TopograficMap{
		topography: [][]uint8{
			{9, 2, 3},
			{8, 1, 4},
			{7, 6, 5},
		},
		width:  3,
		height: 3,
	}

	resultTree := topograficMap.buildTree(tree)

	if len(resultTree.branches) != 1 {
		t.Errorf("Expected 1 branch, got %d", len(resultTree.branches))
	}

	expectedBranchPosition := Point{x: 0, y: 1, height: 2}
	if resultTree.branches[0].position != expectedBranchPosition {
		t.Errorf("Expected branch position %+v, got %+v", expectedBranchPosition, resultTree.branches[0].position)
	}

}
