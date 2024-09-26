function numIslands(grid) {
  if (!grid || grid.length === 0) return 0;

  let count = 0;
  const rows = grid.length;
  const cols = grid[0].length;

  // Helper function for DFS
  function dfs(r, c) {
    // Check bounds and if the cell is water ('0') or already visited ('2')
    if (r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] === '0') return;

    // Mark the cell as visited
    grid[r][c] = '0';

    // Visit all four adjacent cells (up, down, left, right)
    dfs(r + 1, c);
    dfs(r - 1, c);
    dfs(r, c + 1);
    dfs(r, c - 1);
  }

  // Iterate through each cell in the grid
  for (let r = 0; r < rows; r++) {
    for (let c = 0; c < cols; c++) {
      // If the cell is land ('1'), start a DFS to mark the entire island
      if (grid[r][c] === '1') {
        count++;
        dfs(r, c);
      }
    }
  }

  return count;
}

const grid1 = [
  ['1', '1', '1', '1', '0'],
  ['1', '1', '0', '1', '0'],
  ['1', '1', '0', '0', '0'],
  ['0', '0', '0', '0', '0'],
];
console.log(numIslands(grid1)); // Expected output: 1

const grid2 = [
  ['1', '1', '0', '0', '0'],
  ['1', '1', '0', '0', '0'],
  ['0', '0', '1', '0', '0'],
  ['0', '0', '0', '1', '1'],
];
console.log(numIslands(grid2)); // Expected output: 3