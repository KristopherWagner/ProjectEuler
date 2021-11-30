def get_instructions():
  inputFile = open("input.txt", "r")
  instructions = inputFile.readlines()
  inputFile.close()
  return instructions

def change_value(lastChange, instructions):
  i = 0
  j = 0
  found = False
  while not found and i < len(instructions):
    command = instructions[i][:3]
    if command == "nop" or command == "jmp":
      j  += 1
      if j > lastChange:
        found = True
        if command == "nop":
          instructions[i] = "jmp" + instructions[i][3:]
        else:
          instructions[i] = "nop" + instructions[i][3:]
    i += 1
  return instructions

def follow_loop(instructions):
  acc = 0
  completed = {}
  found = False
  i = 0
  while not found:
    try:
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
    except IndexError:
      return (acc, True)
  return (acc, False)

def main():
  i = 0
  acc = 0
  foundAnswer = False
  while not foundAnswer:
    instructions = get_instructions()
    newInstructions = change_value(i, instructions)
    (acc, foundAnswer) = follow_loop(newInstructions)
    i += 1
  return acc

if __name__ == "__main__":
  print main()
