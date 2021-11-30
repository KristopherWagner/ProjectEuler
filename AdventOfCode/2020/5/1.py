def find_seat(bPass, low, high):
  for char in bPass:
    middle = (high + low) / 2
    if char == 'F' or char == 'L':
      high = middle
    elif char == 'B' or char == 'R':
      low = middle
  return high

def main():
  #passes = ['BFFFBBFRRR', 'FFFBBBFRRR', 'BBFFBBFRLL']
  inputFile = open('input.txt', 'r')
  passes = inputFile.readlines()
  inputFile.close()

  highest = 0

  for bPass in passes:
    row = find_seat(bPass[0:7], 0, 127)
    column = find_seat(bPass[7:], 0, 7)
    seatID = row * 8 +  column
    if seatID > highest:
      highest = seatID
  print highest 

if __name__ == '__main__':
  main()
