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


def main(preambleLength):
  data = get_data()
  i = preambleLength
  answer = 0
  while answer == 0 and i < len(data):
    if not is_valid(data[i - preambleLength: i], data[i]):
      answer = data[i]
    i += 1
  return answer

if __name__ == "__main__":
  print main(25)
