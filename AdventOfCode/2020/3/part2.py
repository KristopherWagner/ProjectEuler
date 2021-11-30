def main():
  inputFile = open("input.txt")
  trees = inputFile.readlines()
  inputFile.close()
  i = 0
  while i < len(trees):
    trees[i] = trees[i].strip()
    i += 1

  numTrees = countTrees(trees, (1, 1))
  numTrees *= countTrees(trees, (3, 1))
  numTrees *= countTrees(trees, (5, 1))
  numTrees *= countTrees(trees, (7, 1))
  numTrees *= countTrees(trees, (1, 2))
  print numTrees


def countTrees(trees, slope):
  x = 0
  y = 0
  numTrees = 0

  while x != -1 and y != -1:
    (x, y, isTree) = move((x, y), trees, slope)
    if isTree:
      numTrees += 1
  return numTrees

def move(curPos, trees, slope):
  (x, y) = curPos
  (dx, dy) = slope
  x += dx
  if x >= len(trees[y]):
    x -= len(trees[y])

  y += dy
  if y >= len(trees):
    return (-1, -1, False)

  if trees[y][x] == '#':
    return (x, y, True)
  else:
    return (x, y, False)


if __name__ == "__main__":
  main()
