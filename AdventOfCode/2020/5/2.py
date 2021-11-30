def find_seat(bPass, low, high):
  for char in bPass:
    middle = (high + low) / 2
    if char == 'F' or char == 'L':
      high = middle
    elif char == 'B' or char == 'R':
      low = middle
  return high

def valid_seats():
  seats = {}
  i = 0
  while i <= (127 * 8 + 7):
    seats.update({i: False})
    i += 1
  return seats

def main():
  inputFile = open('input.txt', 'r')
  passes = inputFile.readlines()
  inputFile.close()

  seats = valid_seats()

  for bPass in passes:
    row = find_seat(bPass[0:7], 0, 127)
    column = find_seat(bPass[7:], 0, 7)
    seatID = row * 8 + column
    seats.update({ seatID: True })

  keys = seats.keys()
  mySeat = 0
  i = 1
  while i < (127 * 8 + 7 - -1):
    if seats[i] == False:
      if seats[i - 1] == True and seats[i + 1] == True:
        mySeat = i  
    i += 1
  print mySeat
    
if __name__ == '__main__':
  main()
