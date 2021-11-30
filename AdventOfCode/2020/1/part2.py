def main():
  inputFile = open("input.txt", "r")
  expenses = inputFile.readlines()
  inputFile.close()
  expenses = sorted(expenses)

  i = 0
  answer = 0

  while answer == 0 and i < len(expenses):
    first = int(expenses[i])
    j = i + 1

    while answer == 0 and j < len(expenses):
      second = int(expenses[j])
      k = j + 1

      while answer == 0 and k < len(expenses) and first + second + int(expenses[k]) <= 2020:
        third = int(expenses[k])
        if first + second + third == 2020:
          answer = first * second * third
          print answer
        k += 1
      j += 1
    i += 1

if __name__ == '__main__': main()
