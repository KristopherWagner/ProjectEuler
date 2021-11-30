def main():
  inputFile = open("input.txt", "r")
  lines = inputFile.readlines()
  inputFile.close()

  passport = []
  i = 0
  numValid = 0
  while i <= len(lines):
    if i < len(lines):
      items = clean_data(lines[i])
    if items[0] == '' or i == len(lines):
      if is_passport_valid(passport):
        numValid += 1
      passport = []
    else:
      passport = passport + items
    i += 1
  print numValid

def clean_data(line):
    data = line.split(" ")
    i = 0
    while i < len(data):
      data[i] = data[i].strip().split(":")[0]
      i += 1
    return data
 
def is_passport_valid(passport):
  fields = ['byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid']
  if len(passport) < len(fields):
    return False
  try:
    for field in fields:
      passport.index(field)
  except:
      return False
  return True

if __name__ == "__main__":
  main()
