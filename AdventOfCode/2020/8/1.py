def get_instructions():
  inputFile = open("input.txt", "r")
  instructions = inputFile.readlines()
  inputFile.close()
  return instructions

def main():
  instructions = get_instructions()
  acc = 0
  completed = {}
  found = False
  i = 0
  while not found:
    command = instructions[i][:3]
    sign = instructions[i][4]
    value = int(instructions[i][5:])

    if sign == "-":
      value = 0 - value

    completed.update({ i: True })
    if command == "jmp":
      i += value
    else:
      i += 1
      if command == "acc":
        acc += value
    try:
      if completed[i]:
        found = True
    except KeyError:
      pass 
  return acc

if __name__ == "__main__":
  print main()
