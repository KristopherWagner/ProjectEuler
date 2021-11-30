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
  [times, letter, pwd] = password.split(' ')
  [minimum, maximum] = times.split('-')
  minimum = int(minimum)
  maximum = int(maximum)

  c = 0
  i = 0
  while i < len(pwd) and c <= maximum:
    if pwd[i] == letter[0]:
      c += 1
    i += 1
  return c <= maximum and c >= minimum

if __name__ == '__main__': main()
