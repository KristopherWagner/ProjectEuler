def get_data():
  inputFile = open("input.txt", "r")
  data = inputFile.readlines()
  inputFile.close()

  for i in xrange(0, len(data)):
    data[i] = int(data[i].strip())
  return data

def is_valid(preamble, number):
  i = 0
  found = False
  while not found and i < len(preamble):
    j = i + 1
    while not found and j < len(preamble):
      found = (preamble[i] + preamble[j]) == number
      j += 1
    i += 1
  return found


def get_target(data, preambleLength):
  i = preambleLength
  answer = 0
  while answer == 0 and i < len(data):
    if not is_valid(data[i - preambleLength: i], data[i]):
      answer = data[i]
    i += 1
  return answer

def get_answer(data, target):
  i = 0
  answer = 0
  while answer == 0 and i < len(data):
    j = 0
    count = 0
    while answer == 0 and (i + j) < len(data) and count < target:
      count += data[i + j]
      if count == target:
        values = data[i: i + j + 1]
        answer = min(values) + max(values)
      j += 1
    i += 1
  return answer

if __name__ == "__main__":
  data = get_data()
  target = get_target(data, 25)
  print get_answer(data, target)
