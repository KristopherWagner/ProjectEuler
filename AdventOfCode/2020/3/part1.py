def main():
  inputFile = open("input.txt")
  trees = inputFile.readlines()
  inputFile.close()
  i = 0
  while i < len(trees):
    trees[i] = trees[i].strip()
    i += 1

  x = 0
  y = 0
  numTrees = 0

  while x != -1 and y != -1:
    (x, y, isTree) = move((x, y), trees)
    if isTree:
      numTrees += 1

  print numTrees

def move(curPos, trees):
  (x, y) = curPos
  x += 3
  if x >= len(trees[y]):
    x -= len(trees[y])

  y += 1
  if y >= len(trees):
    return (-1, -1, False)

  if trees[y][x] == '#':
    return (x, y, True)
  else:
    return (x, y, False)


if __name__ == "__main__":
  main()
