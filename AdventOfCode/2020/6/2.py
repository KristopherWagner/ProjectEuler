def calc_score(family):
  score = 0
  for i in range(97, 123):
    yes = True
    j = 0
    while j < len(family) and yes:
      yes = family[j].find(chr(i)) != -1
      j += 1
    if yes:
      score += 1
  return score

def main():
  inputFile = open('input.txt', 'r')
  form = inputFile.readlines()
  inputFile.close()

  family = []
  answer = 0
  for i in xrange(0, len(form)):
    isLastLine = (i + 1) == len(form)
    line = form[i].strip()
    if len(line) > 0:
      family.append(line)
    if len(line) == 0 or isLastLine:
      answer += calc_score(family)
      family = []
  print answer
      

if __name__ == '__main__':
  main()
