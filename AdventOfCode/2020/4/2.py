import re

def clean_data(line):
    data = line.split(" ")
    i = 0
    while i < len(data):
      data[i] = data[i].strip()
      i += 1
    return data
 
def is_passport_present(passport):
  fields = ['byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid']
  passportWithoutValues = []
  if len(passport) < len(fields):
    return False
  for item in passport:
    passportWithoutValues = passportWithoutValues + item.split(':')
  try:
    for field in fields:
      passportWithoutValues.index(field)
  except:
      return False
  return True

def get_present_passports(lines):
  presentPassports = []
  passport = []
  i = 0
  while i <= len(lines):
    if i < len(lines):
      items = clean_data(lines[i])
    if items[0] == '' or i == len(lines):
      if is_passport_present(passport):
        presentPassports.append(passport)
      passport = []
    else:
      passport = passport + items
    i += 1
  return presentPassports

def is_hgt_valid(hgt):
  height = 0
  unit = ''
  if len(hgt) == 4:
    height = hgt[0:2]
    unit = hgt[2:]
  elif len(hgt) ==5:
    height = hgt[0:3]
    unit = hgt[3:]
  else:
    return False

  if height.isdigit():
    if unit == 'cm':
      return height <= '193' and height >= '150'
    elif unit == 'in':
      return height <= '76' and height >= '59'
    else:
      return False
  else:
    return False

def is_hcl_valid(hcl):
 return re.match('^#[0-9a-f]{6}', hcl) != None


def is_passport_valid(passport):
  for item in passport:
    [field, value] = item.split(':')
    if field == 'byr':
      if not (value.isdigit() and value >= '1920' and value <= '2002'):
        return False
    elif field == 'iyr':
      if not (value.isdigit() and value >= '2010' and value <= '2020'):
        return False
    elif field == 'eyr':
      if not (value.isdigit() and value >= '2020' and value <= '2030'):
        return False
    elif field == 'hgt':
      if not is_hgt_valid(value):
        return False
    elif field == 'hcl':
      if not is_hcl_valid(value):
        return False
    elif field == 'ecl':
      try:
        ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth'].index(value)
      except:
        return False
    elif field == 'pid':
      if not (value.isdigit() and len(value) == 9):
        return False
    elif field == 'cid':
      continue
    else:
      return False
  return True 
      
 
def main():
  inputFile = open("input.txt", "r")
  lines = inputFile.readlines()
  inputFile.close()

  presentPassports = get_present_passports(lines)
  answer = 0
  for passport in presentPassports:
    if is_passport_valid(passport):
      answer += 1
  print answer

if __name__ == "__main__":
  main()
