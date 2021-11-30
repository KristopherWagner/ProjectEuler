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
    second = 0
    while answer == 0 and first + second < 2020:
      second = int(expenses[j])
      if first + second == 2020:
        answer = first * second
      j += 1
    i += 1
  print answer

if __name__ == '__main__': main()
