def calc_score(family):
  score = 0
  for key in family:
    if family[key]:
      score += 1
  return score

def calc_family(family, line):
  for char in line:
    family.update({ char: True })
  return family

def main():
  inputFile = open('input.txt', 'r')
  form = inputFile.readlines()
  inputFile.close()

  family = {}
  answer = 0
  for i in xrange(0, len(form)):
    isLastLine = (i + 1) == len(form)
    line = form[i].strip()
    if len(line) > 0:
      family = calc_family(family, line)
    if len(line) == 0 or isLastLine:
      answer += calc_score(family)
      family.clear()
  print answer
      

if __name__ == '__main__':
  main()
