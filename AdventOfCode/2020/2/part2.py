def main():
  inputFile = open("input.txt", "r")
  passwords = inputFile.readlines()
  inputFile.close()
  valid = 0

  for pwd in passwords:
    if isValid(pwd):
      valid += 1
  print valid

def isValid(password):
  [positions, letter, pwd] = password.split(' ')
  [first, second] = positions.split('-')
  first = int(first) - 1
  second = int(second) - 1
  letter = letter[0]

  return (pwd[first] == letter) ^ (pwd[second] == letter)

if __name__ == '__main__': main()
